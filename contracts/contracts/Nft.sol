// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "@openzeppelin/contracts/utils/math/Math.sol";
import "./NftPriceLib.sol";
import "./NftFeeLib.sol";

// ============================================================================
// 接口定义
// ============================================================================

/// @notice IToken0Token1AmmQuote 由 NftFeeLib 提供
/// @notice 质押分红合约接口：NFT 在 setRewardOnce 盈利时调用
interface IUmbrellaStakeDividend {
    function setReward(address user, uint256 profit) external;
}

/// @notice 二级市场月化分配：查 seller 的 B(parentOf)、C(topOf) 地址；B 0.3%/月、C 0.2%/月
interface IStakeReferral {
    function parentOf(address user) external view returns (address);
    function topOf(address user) external view returns (address);
    function isTopEligible(address top) external view returns (bool);
}

/// @notice 可选卖家收款接收器：若 seller 是 vault 等合约，可声明真实收款地址，实现点对点结算
interface ISalePayoutRecipient {
    function payoutRecipient() external view returns (address);
}

/// @title BlindBox NFT + Escrow Marketplace + HoldingValue
/// @author BlindBox Team
/// @notice 盲盒 NFT：铸造、开盒、二级市场托管交易、持有价值、槽位/开盒次数购买
/// @dev 核心逻辑：
///   - HoldingValue = 未开盒盲盒 usdtPaid 聚合 + Piece burn 追加价值
///   - 未开盒持仓不再受仓位限制；仅保留每日开盒次数：用户基础次数(未设置则 baseDailyOpens) + extraOpenCount
///   - 开盒：自造自开收 1% 等值 HRX；市场买回来的盒子开盒不再收 USDT/HRX，仅消耗链上 gas
///   - 二级市场：月化 4% 分配（卖家 3.5%、B 0.3%、C 0.2%）；1% USDT 或 5‰ HRX 手续费
contract BlindBoxNFT is ERC721, AccessControl, ReentrancyGuard {
    bytes32 public constant SUPER_ADMIN_ROLE = keccak256("SUPER_ADMIN_ROLE");
    error ErrZero();
    error ErrDenom();
    error ErrFee();
    error ErrTier();
    error ErrBase();
    error ErrPool();
    error ErrOnlyPiece();
    error ErrQuantity();
    error ErrOwner();
    error ErrOpened();
    error ErrSelfMint();
    error ErrReceiver();
    error ErrFeeB();
    error ErrLimit();
    error ErrNonexistent();
    error ErrRewardSet();
    error ErrCount();
    error ErrListed();
    error ErrUnopened();
    error ErrSeller();
    error ErrEscrow();
    error ErrSelfBuy();
    error ErrToken();
    error ErrNotImplemented();
    using SafeERC20 for IERC20;

    // -------------------- Roles --------------------
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    // -------------------- Constants --------------------
    uint256 public constant DECIMALS_18 = 1e18;
    uint256 public constant MONTH_SECONDS = 30 days;

    // open窗口：12:00~次日12:00
    uint256 public constant DAY_SHIFT_SECONDS = 12 hours;

    /// @notice B 分成资格：holdingValueOf(B) >= 300 USDT 才能获得分成
    uint256 public constant B_HOLDING_THRESHOLD_E18 = 300e18;

    // -------------------- Adjustable Params --------------------
    uint256 public bpsDenominator = 10_000;

    /// @notice mint 分账：给 AccountA 的 bps（默认 9500 = 95%）
    uint256 public accountAFeeBps = 9_500;

    /// @notice 二级市场线性增长：新铸造 token 使用的默认线性月增长率（18位）
    uint256 public monthlyRateDefaultE18 = 40e15; // 4%

    // -------------------- Daily base open (adjustable) --------------------
    /// @notice 每天基础可 open 次数（可调，默认 3）
    uint256 public baseDailyOpens;
    mapping(address => uint256) public userBaseDailyOpens;
    mapping(address => bool) public userBaseDailyOpensSet;

    // -------------------- Tokens --------------------
    IERC20 public immutable usdt;     // BSC USDT(18)
    IERC20 public aToken;             // A Token(18)，预留扩展用

    /// @notice 铸造 USDT 分账收款地址（accountAFeeBps 份额，默认 95%）
    address public accountA;

    /// @notice 部署时默认作为 slotPurchaseReceiver 的初始值；可通过 setSlotPurchaseReceiver 独立修改
    address public accountB;

    /// @notice 二级市场：1% USDT 手续费收款地址；月化分配中 B/C 不达标部分的项目方份额收款地址
    address public accountC;

    /// @notice 二级市场：5‰ HRX 手续费（buyPayFeeInB）收款地址
    address public accountD;

    /// @notice 质押分红合约：mint 剩余 USDT 直接打给这里；盈利也会回调这里
    IUmbrellaStakeDividend public stakingContract;

    /// @notice 二级市场月化分配：读取卖家的 parentOf(B) 与 topOf(C)，B 0.3%/月、C 0.2%/月；未设置则无分佣
    IStakeReferral public stakeReferralReader;

    /// @notice HRX 报价池：token1=USDT, token0=HRX(B)，用于 quoteFeeInB、开盒/购买 HRX 支付
    IToken0Token1AmmQuote public feeQuotePool;
    IERC20 public feeTokenB; // = feeQuotePool.token0()

    // -------------------- Piece integration --------------------
    /// @notice Piece 合约地址（只有它能调用 holdingAdd）
    address public pieceToken;

    /// @notice 1 Piece(1e18) 价值多少 USDT（18位）
    uint256 public pieceValueUsdtE18;

    /// @dev Piece burn 追加的持有价值累计（USDT 18位）
    mapping(address => uint256) private _pieceHoldingValue;

    // -------------------- Mint tier prices (fixed original mint price) --------------------
    mapping(uint8 => uint256) public tierPrice; // 1=>20,2=>50,3=>100 (18 decimals)

    // -------------------- NFT Data --------------------
    struct BoxInfo {
        uint8 tier;
        uint256 usdtPaid;     // 铸造原价（固定档位价）
        uint64 mintedAt;      // 仅用于二级市场线性增长价格计算
        uint64 openedAt;      // open时间（不可逆）
        int256 reward;
        uint64 rewardSetAt;   // 0=未设置；设置后不可再改

        uint256 tokenId;      // 在 BoxInfo 里带 tokenId
    }
    mapping(uint256 => BoxInfo) public boxInfo;

    /// @notice open 为 public，且不可逆
    mapping(uint256 => bool) public open;

    /// @notice 每个 token 在 mint 时快照自己的增长率（用于二级市场价格，不用于 holdingValue）
    mapping(uint256 => uint256) public tokenMonthlyRateE18;

    // -------------------- Arrays log (global) --------------------
    uint256[] public mintedTokenIds;       // 所有铸造tokenId（日志用，不做删除）
    uint256[] public openActionTokenIds;   // 全局开盲盒记录 tokenId（append-only）

    // -------------------- TokenId Counter --------------------
    uint256 private _nextTokenId = 1;

    // -------------------- Daily open usage --------------------
    struct Usage {
        uint64 dayIndex; // (ts + 12h)/1d
        uint64 used;     // 当前窗口已用open次数
    }
    mapping(address => Usage) private _usage;
    mapping(address => Usage) private _secondaryBuyUsage;

    // -------------------- 开盒次数（可配置首项与等差）；仓位购买接口仅保留兼容 --------------------
    /// @notice 未开盒持仓已无限制；仅开盒次数仍受用户基础次数 + extraOpenCount 控制
    uint256 public slotFirstCost = 5 * DECIMALS_18;  // 首项，默认 5 USDT
    uint256 public slotStepCost = 2 * DECIMALS_18;   // 等差，默认 2 USDT
    uint256 public baseStorageSlots = 1;             // 兼容保留；不再限制持仓
    mapping(address => uint256) public extraStorageSlots;   // 兼容保留；不再限制持仓
    mapping(address => uint256) public extraOpenCount;     // 购买的开盒次数（每天额外次数）
    address public slotPurchaseReceiver;             // 槽位/开盒次数购买收款地址

    // -------------------- 开盒手续费 --------------------
    /// @notice 自造自开：1% USDT价值的 HRX；市场买回来的盒子开盒不再收 USDT/HRX
    mapping(uint256 => address) public minterOf;   // tokenId => 铸造者
    uint256 public openFeeUsdtPurchasedBps = 0;   // 保留兼容；市场买回来的盒子开盒不收 USDT/HRX
    uint256 public openFeeHrxSelfMintBps = 100;   // 自造自开用 HRX 支付时 1% (100/10000)
    uint256 public openFeeHrxPurchasedBps = 0;    // 保留兼容；市场买回来的盒子开盒不收 USDT/HRX
    address public openFeeUsdtReceiver;           // 开盒 USDT 手续费收款地址
    address public openFeeHrxReceiver;            // 开盒 HRX 手续费收款地址
    uint256 public openCountFirstCostB;           // 开盒次数用 HRX 购买时首项成本（启用固定模式时生效）
    uint256 public openCountStepCostB;            // 开盒次数用 HRX 购买时等差步长（启用固定模式时生效）
    bool public openCountBCostFixedEnabled;       // true=使用后台配置的 HRX 数量，false=按池子报价换算

    // -------------------- Quant reserve interface --------------------
    /// @notice 量化资金接收地址（预留接口）；不改变当前 mint 的 95%/5% 主路径，仅用于转出合约内已持有资金
    address public quantReceiver;

    // -------------------- Escrow Marketplace (no manual pricing) --------------------
    struct Listing {
        address seller;
        uint64 listedAt;
        bool active;
    }
    mapping(uint256 => Listing) public listings;

    /// @notice 当前正在上架(托管中)的 tokenId 列表（动态数组）
    uint256[] public listedTokenIds;
    mapping(uint256 => uint256) private _listedPosPlusOne; // tokenId => index+1

    // -------------------- Trade Fee (buyer extra pay, 二级市场购买) --------------------
    /// @notice USDT 支付：手续费 1% = feeRate/feeBase
    uint256 public feeRate;     // 100 = 1%
    uint256 public feeBase;

    /// @notice HRX 支付：手续费 5‰(金本位) = feeRateTwo/feeBaseTwo，换算成 HRX 用 quoteFeeInB
    uint256 public feeRateTwo;  // 50 = 5‰
    uint256 public feeBaseTwo;

    // -------------------- Holding Value (simplified) --------------------
    /// @dev 聚合：Σ base（仅未开盒盲盒的 usdtPaid 之和）
    mapping(address => uint256) private _sumUnopenedBase;

    // -------------------- User Lists: My Unopened & My Opened --------------------
    /// @notice 我当前“未开盒盲盒列表”（托管中也仍在卖家列表里）
    mapping(address => uint256[]) private _myUnopenedBoxes;
    mapping(address => mapping(uint256 => uint256)) private _myUnopenedPosPlusOne; // user => tokenId => idx+1

    /// @notice 我的“已打开盲盒列表”
    mapping(address => uint256[]) private _myOpenedBoxes;
    mapping(address => mapping(uint256 => uint256)) private _myOpenedPosPlusOne; // user => tokenId => idx+1

    // -------------------- Seller sold list (after purchase) --------------------
    /// @notice 卖家的成交 tokenId 列表（仅记录成交，不删除）
    mapping(address => uint256[]) private _sellerSoldTokenIds;

    // -------------------- Events --------------------
    event ParamsChanged(uint256 bpsDenominator, uint256 accountAFeeBps, uint256 monthlyRateDefaultE18);

    event BaseDailyOpensChanged(uint256 oldBaseDailyOpens, uint256 newBaseDailyOpens);
    event UserBaseDailyOpensSet(address indexed user, uint256 count, bool enabled);
    event BaseStorageSlotsChanged(uint256 oldVal, uint256 newVal);
    event UserExtraStorageSlotsSet(address indexed user, uint256 count);
    event UserExtraOpenCountSet(address indexed user, uint256 count);

    event AccountAChanged(address indexed oldAccountA, address indexed newAccountA);
    event AccountBChanged(address indexed oldAccountB, address indexed newAccountB);
    event AccountCChanged(address indexed oldAccountC, address indexed newAccountC);
    event AccountDChanged(address indexed oldAccountD, address indexed newAccountD);

    event StakingContractChanged(address indexed oldStaking, address indexed newStaking);

    event ATokenChanged(address indexed oldToken, address indexed newToken);

    event FeeQuotePoolChanged(address indexed pool, address indexed token0B);

    event TierPriceChanged(uint8 indexed tier, uint256 oldPrice, uint256 newPrice);

    event Minted(address indexed minter, uint256 indexed tokenId, uint8 indexed tier, uint256 usdtPaid);
    event Opened(address indexed owner, uint256 indexed tokenId, uint256 timestamp, uint256 usdtFeePaid, uint256 feeBPaid);

    event RewardSetOnce(address indexed admin, uint256 indexed tokenId, int256 reward, uint256 timestamp, uint256 profitNotified);

    /// @param stepCostLocked 保留事件兼容，恒为 0（槽位/开盒次数已改为等差数列计价）
    /// @param newExtraSlots 购买槽位时为新 extraStorageSlots，购买开盒次数时为新 extraOpenCount
    event SlotsPurchased(address indexed buyer, uint256 count, uint256 stepCostLocked, uint256 totalCost, uint256 newExtraSlots);

    event SlotPriceConfigChanged(uint256 oldFirstCost, uint256 newFirstCost, uint256 oldStepCost, uint256 newStepCost);
    event OpenCountBCostConfigChanged(
        uint256 oldFirstCostB,
        uint256 newFirstCostB,
        uint256 oldStepCostB,
        uint256 newStepCostB,
        bool enabled
    );

    event Listed(address indexed seller, uint256 indexed tokenId, uint256 timestamp);
    event Unlisted(address indexed operator, uint256 indexed tokenId);

    event Purchased(
        address indexed buyer,
        address indexed seller,
        uint256 indexed tokenId,
        uint256 priceUSDT,
        bool feePaidInB,
        uint256 feeUSDT,
        uint256 feeB
    );

    event SecondarySaleSplit(
        address indexed buyer,
        address indexed seller,
        uint256 indexed tokenId,
        address b,
        address c,
        uint256 priceUSDT,
        uint256 toSeller,
        uint256 toB,
        uint256 toC,
        uint256 toProject,
        bool feePaidInB,
        uint256 feeUSDT,
        uint256 feeB
    );

    event FeeConfigUSDTChanged(uint256 rate, uint256 base);
    event FeeConfigBChanged(uint256 rateTwo, uint256 baseTwo);

    event ERC20Withdrawn(address indexed token, address indexed to, uint256 amount);
    event QuantReceiverChanged(address indexed oldReceiver, address indexed newReceiver);
    event QuantTransferExecuted(
        address indexed token,
        address indexed receiver,
        uint256 amount,
        address indexed operator,
        string bizNo
    );
    event InsurancePurchaseRequested(address indexed user, uint256 indexed tokenId, uint256 amountB);
    event PropPurchaseRequested(address indexed user, uint256 indexed propType, uint256 amountB);

    event PieceTokenChanged(address indexed oldPiece, address indexed newPiece);
    event PieceValueChanged(uint256 oldValueUsdtE18, uint256 newValueUsdtE18);

    event PieceHoldingAdded(address indexed user, uint256 burnAmount, uint256 valueAddedUsdt, uint256 newPieceHoldingValue);

    constructor(
        string memory name_,
        string memory symbol_,
        address usdt_,
        address aToken_,
        address accountA_,
        address accountB_,
        address accountC_,
        address accountD_,
        uint256 /* globalSlotStepCost_ */,
        address staking_
    ) ERC721(name_, symbol_) {
        if (usdt_ == address(0)) revert ErrZero();
        if (accountA_ == address(0)) revert ErrZero();

        usdt = IERC20(usdt_);
        aToken = IERC20(aToken_);
        accountA = accountA_;
        accountB = accountB_;
        accountC = accountC_;
        accountD = accountD_;
        stakingContract = IUmbrellaStakeDividend(staking_);
        slotPurchaseReceiver = accountB_;  // 槽位/开盒购买收款默认指向 accountB
        openFeeUsdtReceiver = accountA_;   // 开盒 USDT 手续费默认收给项目方（accountA）
        openFeeHrxReceiver = accountD_;    // 开盒 HRX 手续费默认收给 accountD
        quantReceiver = accountA_;         // 默认先指向项目热钱包，可再单独切到交易所充值/量化地址

        // 每天基础开盒次数，默认 3；槽位可购买，价格等差数列从 slotFirstCost 开始
        baseDailyOpens = 3;

        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(SUPER_ADMIN_ROLE, msg.sender);
        _grantRole(ADMIN_ROLE, msg.sender);

        // 固定原价（18位）
        tierPrice[1] = 20 * DECIMALS_18;
        tierPrice[2] = 50 * DECIMALS_18;
        tierPrice[3] = 100 * DECIMALS_18;

        // 需求：二级市场购买 1% USDT + 5‰ HRX
        // buyPayFeeInUSDT: 手续费 1% USDT -> accountC
        // buyPayFeeInB: 手续费 5‰ 金本位折算 HRX -> accountD（若需 5% 则 feeRateTwo=500）
        feeRate = 100;    // 1% = 100/10000
        feeBase = 10_000;

        feeRateTwo = 50;  // 5‰ = 50/10000（千分之五）
        feeBaseTwo = 10_000;
    }

    // ============================================================
    //                    Admin setters
    // ============================================================

    function setParams(uint256 denom, uint256 aFeeBps, uint256 monthlyRateDefault)
        external
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        if (denom == 0) revert ErrDenom();
        if (aFeeBps > denom) revert ErrFee();
        bpsDenominator = denom;
        accountAFeeBps = aFeeBps;
        monthlyRateDefaultE18 = monthlyRateDefault;
        emit ParamsChanged(denom, aFeeBps, monthlyRateDefault);
    }

    function setBaseDailyOpens(uint256 newBaseDailyOpens) external onlyRole(DEFAULT_ADMIN_ROLE) {
        uint256 old = baseDailyOpens;
        baseDailyOpens = newBaseDailyOpens;
        emit BaseDailyOpensChanged(old, newBaseDailyOpens);
    }

    function setUserBaseDailyOpens(address user, uint256 count, bool enabled) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (user == address(0)) revert ErrZero();
        userBaseDailyOpens[user] = count;
        userBaseDailyOpensSet[user] = enabled;
        emit UserBaseDailyOpensSet(user, count, enabled);
    }

    function setBaseStorageSlots(uint256 newBaseStorageSlots) external onlyRole(DEFAULT_ADMIN_ROLE) {
        uint256 old = baseStorageSlots;
        baseStorageSlots = newBaseStorageSlots;
        emit BaseStorageSlotsChanged(old, newBaseStorageSlots);
    }

    /// @notice 管理员设置某地址的额外槽位（仓库容量）
    function setUserExtraStorageSlots(address user, uint256 count) external onlyRole(DEFAULT_ADMIN_ROLE) {
        extraStorageSlots[user] = count;
        emit UserExtraStorageSlotsSet(user, count);
    }

    /// @notice 管理员设置某地址的额外开盒次数（每天）
    function setUserExtraOpenCount(address user, uint256 count) external onlyRole(DEFAULT_ADMIN_ROLE) {
        extraOpenCount[user] = count;
        emit UserExtraOpenCountSet(user, count);
    }

    function setAccountA(address newAccountA) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (newAccountA == address(0)) revert ErrZero();
        address old = accountA;
        accountA = newAccountA;
        emit AccountAChanged(old, newAccountA);
    }

    function setAccountB(address newAccountB) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (newAccountB == address(0)) revert ErrZero();
        address old = accountB;
        accountB = newAccountB;
        emit AccountBChanged(old, newAccountB);
    }

    function setAccountC(address newAccountC) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (newAccountC == address(0)) revert ErrZero();
        address old = accountC;
        accountC = newAccountC;
        emit AccountCChanged(old, newAccountC);
    }

    function setAccountD(address newAccountD) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (newAccountD == address(0)) revert ErrZero();
        address old = accountD;
        accountD = newAccountD;
        emit AccountDChanged(old, newAccountD);
    }

    function setStakingContract(address newStaking) external onlyRole(DEFAULT_ADMIN_ROLE) {
        address old = address(stakingContract);
        stakingContract = IUmbrellaStakeDividend(newStaking);
        emit StakingContractChanged(old, newStaking);
    }

    function setSlotPurchaseReceiver(address receiver) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (receiver == address(0)) revert ErrZero();
        slotPurchaseReceiver = receiver;
    }

    /// @notice 槽位价格等差数列：第 k 个成本 = slotFirstCost + (k-1) * slotStepCost；总成本用 sum 公式；管理员可调，便于后端配置
    function setSlotPriceConfig(uint256 firstCost, uint256 stepCost) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (firstCost == 0) revert ErrZero();
        uint256 oldFirst = slotFirstCost;
        uint256 oldStep = slotStepCost;
        slotFirstCost = firstCost;
        slotStepCost = stepCost;
        emit SlotPriceConfigChanged(oldFirst, firstCost, oldStep, stepCost);
    }

    /// @notice 开盒次数用 HRX 购买时，允许后台直接配置固定数量规则；关闭时仍按池子报价换算
    function setOpenCountCostConfigB(uint256 firstCostB, uint256 stepCostB, bool enabled)
        external
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        uint256 oldFirst = openCountFirstCostB;
        uint256 oldStep = openCountStepCostB;
        openCountFirstCostB = firstCostB;
        openCountStepCostB = stepCostB;
        openCountBCostFixedEnabled = enabled;
        emit OpenCountBCostConfigChanged(oldFirst, firstCostB, oldStep, stepCostB, enabled);
    }

    /// @dev 买 count 个槽位（当前已有 n 个）的总成本 = Σ(slotFirstCost + (n+i)*slotStepCost), i=0..count-1
    function _slotCost(uint256 n, uint256 count) internal view returns (uint256) {
        return count * slotFirstCost + slotStepCost * count * (2 * n + count - 1) / 2;
    }

    function setOpenFeeUsdtConfig(uint256 bps, address receiver) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (bps > bpsDenominator) revert ErrFee();
        openFeeUsdtPurchasedBps = bps;
        openFeeUsdtReceiver = receiver;
    }

    function setOpenFeeHrxConfig(uint256 selfMintBps, uint256 purchasedBps, address receiver) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (selfMintBps > bpsDenominator || purchasedBps > bpsDenominator) revert ErrFee();
        openFeeHrxSelfMintBps = selfMintBps;
        openFeeHrxPurchasedBps = purchasedBps;
        openFeeHrxReceiver = receiver;
    }

    /// @notice 设置量化接收地址；仅作为预留接口，不改变当前 mint 的直达热钱包路径
    function setQuantReceiver(address receiver) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (receiver == address(0)) revert ErrZero();
        address old = quantReceiver;
        quantReceiver = receiver;
        emit QuantReceiverChanged(old, receiver);
    }

    function setAToken(address newAToken) external onlyRole(DEFAULT_ADMIN_ROLE) {
        address old = address(aToken);
        aToken = IERC20(newAToken);
        emit ATokenChanged(old, newAToken);
    }

    function setTierPrice(uint8 tier, uint256 newPrice) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (tier < 1) revert ErrTier();
        if (newPrice == 0) revert ErrTier();
        uint256 old = tierPrice[tier];
        tierPrice[tier] = newPrice;
        emit TierPriceChanged(tier, old, newPrice);
    }

    function setFeeConfigUSDT(uint256 rate, uint256 base) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (base == 0) revert ErrBase();
        feeRate = rate;
        feeBase = base;
        emit FeeConfigUSDTChanged(rate, base);
    }

    function setFeeConfigB(uint256 rateTwo, uint256 baseTwo) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (baseTwo == 0) revert ErrBase();
        feeRateTwo = rateTwo;
        feeBaseTwo = baseTwo;
        emit FeeConfigBChanged(rateTwo, baseTwo);
    }

    function setFeeQuotePool(address pool) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (pool == address(0)) revert ErrZero();

        IToken0Token1AmmQuote p = IToken0Token1AmmQuote(pool);
        // liquidityInited 在 quoteFeeInB 调用时检查，此处仅校验 token 配对
        if (p.token1() != address(usdt)) revert ErrPool();

        address b = p.token0();
        if (b == address(0)) revert ErrPool();

        feeQuotePool = p;
        feeTokenB = IERC20(b);

        emit FeeQuotePoolChanged(pool, b);
    }

    /// @notice 设置二级市场出售分佣读取器（Stake 地址）；未设置时手续费全给 accountC/accountD
    function setStakeReferralReader(address reader) external onlyRole(DEFAULT_ADMIN_ROLE) {
        stakeReferralReader = IStakeReferral(reader);
    }

    // ---- Piece config ----

    function setPieceToken(address newPiece) external onlyRole(DEFAULT_ADMIN_ROLE) {
        address old = pieceToken;
        pieceToken = newPiece;
        emit PieceTokenChanged(old, newPiece);
    }

    function setPieceValueUsdtE18(uint256 newValueUsdtE18) external onlyRole(DEFAULT_ADMIN_ROLE) {
        uint256 old = pieceValueUsdtE18;
        pieceValueUsdtE18 = newValueUsdtE18;
        emit PieceValueChanged(old, newValueUsdtE18);
    }

    /// @notice Piece burn 回调：只有 pieceToken 能调用
    function holdingAdd(address user, uint256 burnAmount) external {
        if (msg.sender != pieceToken) revert ErrOnlyPiece();
        if (user == address(0)) revert ErrZero();
        if (burnAmount == 0) revert ErrZero();
        if (pieceValueUsdtE18 == 0) revert ErrZero();

        uint256 valueAdded = Math.mulDiv(burnAmount, pieceValueUsdtE18, DECIMALS_18);
        _pieceHoldingValue[user] += valueAdded;

        emit PieceHoldingAdded(user, burnAmount, valueAdded, _pieceHoldingValue[user]);
    }

    // ============================================================
    //                           Minting
    // ============================================================

    function mint(uint8 tier) external nonReentrant returns (uint256 tokenId) {
        uint256 price = _requireTierPrice(tier);
        _collectUsdtMint(msg.sender, price);
        tokenId = _mintInternal(msg.sender, tier, price);
    }

    function mintBatch(uint8 tier, uint256 quantity) external nonReentrant returns (uint256[] memory tokenIds) {
        if (quantity == 0) revert ErrQuantity();
        uint256 priceEach = _requireTierPrice(tier);
        uint256 total = priceEach * quantity;

        _collectUsdtMint(msg.sender, total);

        tokenIds = new uint256[](quantity);
        for (uint256 i = 0; i < quantity; i++) {
            tokenIds[i] = _mintInternal(msg.sender, tier, priceEach);
        }
    }

    function _requireTierPrice(uint8 tier) internal view returns (uint256 p) {
        if (tier < 1) revert ErrTier();
        p = tierPrice[tier];
        if (p == 0) revert ErrTier();
    }

    function _mintInternal(address to, uint8 tier, uint256 paid) internal returns (uint256 tokenId) {
        tokenId = _nextTokenId++;

        BoxInfo storage b = boxInfo[tokenId];
        b.tier = tier;
        b.usdtPaid = paid;
        b.mintedAt = uint64(block.timestamp);

        b.tokenId = tokenId;

        tokenMonthlyRateE18[tokenId] = monthlyRateDefaultE18;
        minterOf[tokenId] = to;

        _safeMint(to, tokenId);

        mintedTokenIds.push(tokenId);

        emit Minted(to, tokenId, tier, paid);
    }

    /// @dev mint收款：toA 给 AccountA；其余直接给 stakingContract
    function _collectUsdtMint(address payer, uint256 total) internal {
        if (accountA == address(0)) revert ErrZero();
        if (address(stakingContract) == address(0)) revert ErrZero();

        uint256 toA = Math.mulDiv(total, accountAFeeBps, bpsDenominator);
        uint256 toStake = total - toA;

        if (toA > 0) usdt.safeTransferFrom(payer, accountA, toA);
        if (toStake > 0) usdt.safeTransferFrom(payer, address(stakingContract), toStake);
    }

    // ============================================================
    //                    Open (irreversible)
    // 自造自开：1% USDT价值的 HRX；购买的开盒：1% USDT 或 5% USDT价值的 HRX
    // ============================================================

    /// @notice 开盒（USDT 支付）：仅适用于购买的开盒，收 1% USDT
    /// @dev Purchased boxes open without token fees; this legacy entrypoint is kept for ABI compatibility.
    function openBoxPayUsdt(uint256 tokenId) external nonReentrant {
        if (ownerOf(tokenId) != msg.sender) revert ErrOwner();
        if (open[tokenId]) revert ErrOpened();

        uint256 paid = boxInfo[tokenId].usdtPaid;
        bool isSelfMint = (minterOf[tokenId] == msg.sender);
        if (isSelfMint) revert ErrSelfMint();

        _doOpen(msg.sender, tokenId, paid, 0, 0);
    }

    /// @notice 开盒（HRX 支付）：自造自开收 1% USDT价值的 HRX，购买的开盒收 5% USDT价值的 HRX
    /// @dev Only self-minted boxes pay HRX here. Purchased boxes open without token fees.
    function openBoxPayHRX(uint256 tokenId, uint256 maxFeeB) external nonReentrant {
        if (ownerOf(tokenId) != msg.sender) revert ErrOwner();
        if (open[tokenId]) revert ErrOpened();

        uint256 paid = boxInfo[tokenId].usdtPaid;
        bool isSelfMint = (minterOf[tokenId] == msg.sender);
        if (!isSelfMint) {
            maxFeeB;
            _doOpen(msg.sender, tokenId, paid, 0, 0);
            return;
        }

        if (address(feeQuotePool) == address(0)) revert ErrPool();
        uint256 feeU = Math.mulDiv(paid, openFeeHrxSelfMintBps, bpsDenominator);

        uint256 feeB = 0;
        if (feeU > 0) {
            feeB = quoteFeeInB(feeU);
            if (maxFeeB > 0 && feeB > maxFeeB) revert ErrFeeB();
            if (openFeeHrxReceiver == address(0)) revert ErrReceiver();
            feeTokenB.safeTransferFrom(msg.sender, openFeeHrxReceiver, feeB);
        }

        _doOpen(msg.sender, tokenId, paid, 0, feeB);
    }

    function _doOpen(address owner_, uint256 tokenId, uint256 base, uint256 feeU, uint256 feeB) internal {
        _consumeOpenAllowance(owner_);

        _holdingRemoveBox(owner_, base);
        _myUnopenedRemove(owner_, tokenId);
        _myOpenedAdd(owner_, tokenId);

        open[tokenId] = true;
        boxInfo[tokenId].openedAt = uint64(block.timestamp);

        openActionTokenIds.push(tokenId);

        emit Opened(owner_, tokenId, block.timestamp, feeU, feeB);
    }

    function _consumeOpenAllowance(address user) internal {
        uint256 dayIdx = _dayIndex(block.timestamp);

        Usage storage u = _usage[user];
        if (u.dayIndex != dayIdx) {
            u.dayIndex = uint64(dayIdx);
            u.used = 0;
        }

        uint256 limit = _baseDailyOpensOf(user) + extraOpenCount[user];
        if (uint256(u.used) >= limit) revert ErrLimit();
        u.used += 1;
    }

    function baseDailyOpensOf(address user) public view returns (uint256) {
        return _baseDailyOpensOf(user);
    }

    function _baseDailyOpensOf(address user) internal view returns (uint256) {
        return userBaseDailyOpensSet[user] ? userBaseDailyOpens[user] : baseDailyOpens;
    }

    function _dayIndex(uint256 ts) internal pure returns (uint256) {
        return (ts + DAY_SHIFT_SECONDS) / 1 days;
    }

    // ============================================================
    //                        Reward (set once + transfer payout + notify staking)
    // ============================================================

    /// @notice 只设置一次；并根据 reward 发放本金/部分本金给当前 owner
    ///         - reward >= paid => 转 paid
    ///         - 0 < reward < paid => 转 reward
    ///         - reward > paid 的利润部分 => 通知 stakingContract.setReward(holder, reward - paid)
    /// @dev 合约必须有足够 USDT 余额，否则 safeTransfer 会 revert
    function setRewardOnce(uint256 tokenId, int256 reward_) external onlyRole(ADMIN_ROLE) nonReentrant {
        if (_ownerOf(tokenId) == address(0)) revert ErrNonexistent();
        if (!open[tokenId]) revert ErrOpened();
        if (boxInfo[tokenId].rewardSetAt != 0) revert ErrRewardSet();

        boxInfo[tokenId].reward = reward_;
        boxInfo[tokenId].rewardSetAt = uint64(block.timestamp);

        address holder = ownerOf(tokenId);
        uint256 paid = boxInfo[tokenId].usdtPaid;

        uint256 profitNotified = 0;

        // 先发本金/部分本金
        if (reward_ > 0) {
            uint256 r = uint256(reward_);
            uint256 payout = (r >= paid) ? paid : r; // min(r, paid)

            if (payout > 0) {
                usdt.safeTransfer(holder, payout);
            }

            // 盈利部分通知 staking
            if (r > paid) {
                profitNotified = r - paid;
                if (address(stakingContract) == address(0)) revert ErrZero();
                stakingContract.setReward(holder, profitNotified);
            }
        }

        emit RewardSetOnce(msg.sender, tokenId, reward_, block.timestamp, profitNotified);
    }

    // ============================================================
    //                    买槽位（仓库容量） / 买开盒次数
    // ============================================================
    /// @notice 仓位已无限；保留接口兼容，恒不收费
    function buyStorageSlots(uint256 count) external nonReentrant {
        if (count == 0) revert ErrCount();
        emit SlotsPurchased(msg.sender, count, 0, 0, extraStorageSlots[msg.sender]);
    }

    /// @notice 买开盒次数（增加每天可开盒次数），USDT 支付
    function buyOpenCount(uint256 count) external nonReentrant {
        if (count == 0) revert ErrCount();
        if (slotPurchaseReceiver == address(0)) revert ErrReceiver();

        uint256 n = extraOpenCount[msg.sender];
        uint256 cost = _slotCost(n, count);

        usdt.safeTransferFrom(msg.sender, slotPurchaseReceiver, cost);

        extraOpenCount[msg.sender] = n + count;

        emit SlotsPurchased(msg.sender, count, 0, cost, n + count);
    }

    /// @notice 仓位已无限；保留接口兼容，恒为 0
    function quoteStorageSlotCost(address user, uint256 count) external pure returns (uint256 costUSDT) {
        user;
        count;
        return 0;
    }

    /// @notice 买 count 个开盒次数所需 USDT
    function quoteOpenCountCost(address user, uint256 count) external view returns (uint256 costUSDT) {
        uint256 n = extraOpenCount[user];
        return _slotCost(n, count);
    }

    /// @notice 仓位已无限；保留接口兼容，恒为 0
    function quoteSlotCostInB(address user, uint256 count) external pure returns (uint256 costB) {
        user;
        count;
        return 0;
    }

    /// @notice 仓位已无限；保留接口兼容，恒不收费
    function buyStorageSlotsPayHRX(uint256 count, uint256 maxCostB) external nonReentrant {
        if (count == 0) revert ErrCount();
        maxCostB;
        emit SlotsPurchased(msg.sender, count, 0, 0, extraStorageSlots[msg.sender]);
    }

    /// @notice 使用 HRX 购买开盒次数
    function buyOpenCountPayHRX(uint256 count, uint256 maxCostB) external nonReentrant {
        if (count == 0) revert ErrCount();
        if (slotPurchaseReceiver == address(0)) revert ErrReceiver();

        uint256 n = extraOpenCount[msg.sender];
        uint256 costB = _openCountCostInB(n, count);

        if (maxCostB > 0 && costB > maxCostB) revert ErrFeeB();

        feeTokenB.safeTransferFrom(msg.sender, slotPurchaseReceiver, costB);

        extraOpenCount[msg.sender] = n + count;

        emit SlotsPurchased(msg.sender, count, 0, costB, n + count);
    }

    /// @notice 买 count 个开盒次数所需 HRX 数量
    function quoteOpenCountCostInB(address user, uint256 count) external view returns (uint256 costB) {
        if (count == 0) revert ErrCount();
        return _openCountCostInB(extraOpenCount[user], count);
    }

    // ============================================================
    //                 Secondary Market Price (linear growth kept)
    // ============================================================

    function secondaryBuyPrice(uint256 tokenId) public view returns (uint256 price) {
        if (_ownerOf(tokenId) == address(0)) revert ErrNonexistent();
        BoxInfo memory b = boxInfo[tokenId];
        if (b.mintedAt == 0) revert ErrNonexistent();

        uint256 base = b.usdtPaid;
        uint256 rate = tokenMonthlyRateE18[tokenId];

        uint256 elapsed = 0;
        if (block.timestamp > uint256(b.mintedAt)) {
            elapsed = block.timestamp - uint256(b.mintedAt);
        }

        uint256 inc = NftPriceLib.linearPriceIncrement(base, rate, elapsed);
        price = base + inc;
    }

    // ============================================================
    //                      Escrow Marketplace (no manual pricing)
    // ============================================================

    function _addListedToken(uint256 tokenId) internal {
        if (_listedPosPlusOne[tokenId] != 0) revert ErrListed();
        listedTokenIds.push(tokenId);
        _listedPosPlusOne[tokenId] = listedTokenIds.length; // index+1
    }

    function _removeListedToken(uint256 tokenId) internal {
        uint256 posPlusOne = _listedPosPlusOne[tokenId];
        if (posPlusOne == 0) revert ErrListed();

        uint256 idx = posPlusOne - 1;
        uint256 lastIdx = listedTokenIds.length - 1;

        if (idx != lastIdx) {
            uint256 lastTokenId = listedTokenIds[lastIdx];
            listedTokenIds[idx] = lastTokenId;
            _listedPosPlusOne[lastTokenId] = idx + 1;
        }

        listedTokenIds.pop();
        delete _listedPosPlusOne[tokenId];
    }

    function _listOne(address seller, uint256 tokenId, uint64 ts) internal {
        if (ownerOf(tokenId) != seller) revert ErrOwner();
        if (open[tokenId]) revert ErrUnopened();
        if (listings[tokenId].active) revert ErrListed();

        _transfer(seller, address(this), tokenId);

        listings[tokenId] = Listing({
            seller: seller,
            listedAt: ts,
            active: true
        });

        _addListedToken(tokenId);

        emit Listed(seller, tokenId, ts);
    }

    function list(uint256 tokenId) external nonReentrant {
        _listOne(msg.sender, tokenId, uint64(block.timestamp));
    }

    /// @notice 批量上架（原子操作：任意一个失败则全部 revert）
    function listBatch(uint256[] calldata tokenIds) external nonReentrant {
        if (tokenIds.length == 0) revert ErrCount();
        uint64 ts = uint64(block.timestamp);

        for (uint256 i = 0; i < tokenIds.length; i++) {
            _listOne(msg.sender, tokenIds[i], ts);
        }
    }

    function unlist(uint256 tokenId) external nonReentrant {
        Listing memory lst = listings[tokenId];
        if (!lst.active) revert ErrListed();
        if (lst.seller != msg.sender && !hasRole(DEFAULT_ADMIN_ROLE, msg.sender)) revert ErrSeller();

        _removeListedToken(tokenId);

        _safeTransfer(address(this), lst.seller, tokenId, "");

        delete listings[tokenId];

        emit Unlisted(msg.sender, tokenId);
    }

    function _feeUSDT(uint256 priceUSDT) internal view returns (uint256) {
        if (feeRate == 0) return 0;
        if (feeBase == 0) revert ErrBase();
        return Math.mulDiv(priceUSDT, feeRate, feeBase);
    }

    function _feeUSDTForBMode(uint256 priceUSDT) internal view returns (uint256) {
        if (feeRateTwo == 0) return 0;
        if (feeBaseTwo == 0) revert ErrBase();
        return Math.mulDiv(priceUSDT, feeRateTwo, feeBaseTwo);
    }

    function quoteFeeInB(uint256 amount1InUSDT) public view returns (uint256 out0NetB) {
        return NftFeeLib.quoteFeeInB(feeQuotePool, amount1InUSDT, address(usdt), address(feeTokenB));
    }

    /// @dev 月化4%分配：A(卖家)3.5%、B 0.3%、C 0.2%；B需 holdingValue>=300u，C需 isTopEligible；B/C不达标时对应份额归项目方
    function _calcSalePriceSplit(uint256 tokenId, address seller, uint256 totalPrice)
        internal view returns (uint256 toSeller, uint256 toB, uint256 toC, uint256 toProject)
    {
        BoxInfo memory b = boxInfo[tokenId];
        uint256 paid = b.usdtPaid;
        if (paid == 0) {
            toSeller = totalPrice;
            return (toSeller, 0, 0, 0);
        }
        uint256 elapsed = block.timestamp > uint256(b.mintedAt)
            ? block.timestamp - uint256(b.mintedAt)
            : 0;

        if (address(stakeReferralReader) == address(0) || elapsed == 0) {
            toSeller = totalPrice;
            return (toSeller, 0, 0, 0);
        }

        address B = stakeReferralReader.parentOf(seller);
        address C = stakeReferralReader.topOf(seller);
        bool bOk = B != address(0) && holdingValueOf(B) >= B_HOLDING_THRESHOLD_E18;
        bool cOk = C != address(0) && stakeReferralReader.isTopEligible(C);

        uint256 periodDenominator = bpsDenominator * MONTH_SECONDS;
        toSeller = paid + Math.mulDiv(paid, 350 * elapsed, periodDenominator);  // 3.5%/month, per second
        toB = bOk ? Math.mulDiv(paid, 30 * elapsed, periodDenominator) : 0;     // 0.3%/month, per second
        toC = cOk ? Math.mulDiv(paid, 20 * elapsed, periodDenominator) : 0;     // 0.2%/month, per second
        uint256 sum = toSeller + toB + toC;
        if (sum > totalPrice && sum > 0) {
            toSeller = Math.mulDiv(totalPrice, toSeller, sum);
            toB = Math.mulDiv(totalPrice, toB, sum);
            toC = totalPrice - toSeller - toB;
        } else if (sum < totalPrice) {
            toProject = totalPrice - sum;  // B/C 不达标时对应份额归项目方
        }
    }

    /// @dev 手续费 1% 全给项目方 accountC
    function _saleUplineAddresses(address seller) internal view returns (address B, address C) {
        if (address(stakeReferralReader) == address(0)) {
            return (address(0), address(0));
        }
        B = stakeReferralReader.parentOf(seller);
        C = stakeReferralReader.topOf(seller);
    }

    function _distributeSaleFeeUSDT(address buyer, uint256 feeU) internal {
        if (feeU > 0 && accountC != address(0)) usdt.safeTransferFrom(buyer, accountC, feeU);
    }

    /// @dev 二级市场每日购买次数必须严格小于当日可开盒次数
    function _consumeSecondaryBuyAllowance(address user) internal {
        uint256 dayIdx = _dayIndex(block.timestamp);
        Usage storage u = _secondaryBuyUsage[user];
        if (u.dayIndex != dayIdx) {
            u.dayIndex = uint64(dayIdx);
            u.used = 0;
        }

        uint256 limit = _baseDailyOpensOf(user) + extraOpenCount[user];
        if (uint256(u.used) >= limit) revert ErrLimit();
        u.used += 1;
    }

    /// @notice 月化4%：A(卖家)3.5%、B 0.3%、C 0.2%；买价1%手续费给项目方
    function buyPayFeeInUSDT(uint256 tokenId) external nonReentrant {
        Listing memory lst = listings[tokenId];
        if (!lst.active) revert ErrListed();

        if (open[tokenId]) revert ErrUnopened();
        if (ownerOf(tokenId) != address(this)) revert ErrEscrow();
        if (msg.sender == lst.seller) revert ErrSelfBuy();
        _consumeSecondaryBuyAllowance(msg.sender);

        uint256 price = secondaryBuyPrice(tokenId);
        uint256 feeU = _feeUSDT(price);

        _removeListedToken(tokenId);

        (uint256 toSeller, uint256 toB, uint256 toC, uint256 toProject) = _calcSalePriceSplit(tokenId, lst.seller, price);
        (address B, address C) = _saleUplineAddresses(lst.seller);
        address sellerPayee = _salePayoutRecipient(lst.seller);
        usdt.safeTransferFrom(msg.sender, sellerPayee, toSeller);
        if (toB > 0) {
            if (B != address(0)) usdt.safeTransferFrom(msg.sender, B, toB);
        }
        if (toC > 0) {
            if (C != address(0)) usdt.safeTransferFrom(msg.sender, C, toC);
        }
        if (toProject > 0 && accountC != address(0)) usdt.safeTransferFrom(msg.sender, accountC, toProject);

        if (feeU > 0) _distributeSaleFeeUSDT(msg.sender, feeU);

        _safeTransfer(address(this), msg.sender, tokenId, "");

        _sellerSoldTokenIds[lst.seller].push(tokenId);
        delete listings[tokenId];

        emit Purchased(msg.sender, lst.seller, tokenId, price, false, feeU, 0);
        emit SecondarySaleSplit(msg.sender, lst.seller, tokenId, B, C, price, toSeller, toB, toC, toProject, false, feeU, 0);
    }

    function buyPayFeeInB(uint256 tokenId, uint256 maxFeeB) external nonReentrant {
        Listing memory lst = listings[tokenId];
        if (!lst.active) revert ErrListed();

        if (open[tokenId]) revert ErrUnopened();
        if (ownerOf(tokenId) != address(this)) revert ErrEscrow();
        if (msg.sender == lst.seller) revert ErrSelfBuy();

        if (address(feeTokenB) == address(0)) revert ErrZero();
        if (address(feeQuotePool) == address(0)) revert ErrPool();
        _consumeSecondaryBuyAllowance(msg.sender);

        uint256 price = secondaryBuyPrice(tokenId);

        uint256 feeU = _feeUSDTForBMode(price);
        uint256 feeB = 0;

        if (feeU > 0) {
            feeB = quoteFeeInB(feeU);
            if (maxFeeB > 0 && feeB > maxFeeB) revert ErrFeeB();
        }

        _removeListedToken(tokenId);

        (uint256 toSeller, uint256 toB, uint256 toC, uint256 toProject) = _calcSalePriceSplit(tokenId, lst.seller, price);
        (address B, address C) = _saleUplineAddresses(lst.seller);
        address sellerPayee = _salePayoutRecipient(lst.seller);
        usdt.safeTransferFrom(msg.sender, sellerPayee, toSeller);
        if (toB > 0) {
            if (B != address(0)) usdt.safeTransferFrom(msg.sender, B, toB);
        }
        if (toC > 0) {
            if (C != address(0)) usdt.safeTransferFrom(msg.sender, C, toC);
        }
        if (toProject > 0 && accountC != address(0)) usdt.safeTransferFrom(msg.sender, accountC, toProject);
        if (feeB > 0) {
            if (accountD == address(0)) revert ErrZero();
            feeTokenB.safeTransferFrom(msg.sender, accountD, feeB);
        }

        _safeTransfer(address(this), msg.sender, tokenId, "");

        // 卖家成交记录（append-only）
        _sellerSoldTokenIds[lst.seller].push(tokenId);

        delete listings[tokenId];

        emit Purchased(msg.sender, lst.seller, tokenId, price, true, feeU, feeB);
        emit SecondarySaleSplit(msg.sender, lst.seller, tokenId, B, C, price, toSeller, toB, toC, toProject, true, feeU, feeB);
    }

    // ============================================================
    //                 Holding Value (simplified + piece addon)
    // ============================================================

    function holdingValueOf(address user) public view returns (uint256) {
        return _sumUnopenedBase[user] + _pieceHoldingValue[user];
    }

    function holdingBreakdownOf(address user) external view returns (uint256 unopenedSum, uint256 pieceSum, uint256 total) {
        unopenedSum = _sumUnopenedBase[user];
        pieceSum = _pieceHoldingValue[user];
        total = unopenedSum + pieceSum;
    }

    function pieceHoldingOf(address user) external view returns (uint256) {
        return _pieceHoldingValue[user];
    }

    function _holdingAddBox(address user, uint256 base) internal {
        if (base == 0) return;
        _sumUnopenedBase[user] += base;
    }

    function _holdingRemoveBox(address user, uint256 base) internal {
        if (base == 0) return;
        _sumUnopenedBase[user] -= base;
    }

    // ============================================================
    //                  User lists (dynamic arrays) + pagination
    //                  (分页只返回 BoxInfo[]，带 tokenId)
    // ============================================================

    function _myUnopenedAdd(address user, uint256 tokenId) internal {
        if (_myUnopenedPosPlusOne[user][tokenId] != 0) revert ErrListed();
        _myUnopenedBoxes[user].push(tokenId);
        _myUnopenedPosPlusOne[user][tokenId] = _myUnopenedBoxes[user].length; // idx+1
    }

    function _myUnopenedRemove(address user, uint256 tokenId) internal {
        uint256 posPlusOne = _myUnopenedPosPlusOne[user][tokenId];
        if (posPlusOne == 0) revert ErrListed();

        uint256 idx = posPlusOne - 1;
        uint256 lastIdx = _myUnopenedBoxes[user].length - 1;

        if (idx != lastIdx) {
            uint256 lastTokenId = _myUnopenedBoxes[user][lastIdx];
            _myUnopenedBoxes[user][idx] = lastTokenId;
            _myUnopenedPosPlusOne[user][lastTokenId] = idx + 1;
        }

        _myUnopenedBoxes[user].pop();
        delete _myUnopenedPosPlusOne[user][tokenId];
    }

    function _myOpenedAdd(address user, uint256 tokenId) internal {
        if (_myOpenedPosPlusOne[user][tokenId] != 0) revert ErrListed();
        _myOpenedBoxes[user].push(tokenId);
        _myOpenedPosPlusOne[user][tokenId] = _myOpenedBoxes[user].length; // idx+1
    }

    function _myOpenedRemove(address user, uint256 tokenId) internal {
        uint256 posPlusOne = _myOpenedPosPlusOne[user][tokenId];
        if (posPlusOne == 0) revert ErrListed();

        uint256 idx = posPlusOne - 1;
        uint256 lastIdx = _myOpenedBoxes[user].length - 1;

        if (idx != lastIdx) {
            uint256 lastTokenId = _myOpenedBoxes[user][lastIdx];
            _myOpenedBoxes[user][idx] = lastTokenId;
            _myOpenedPosPlusOne[user][lastTokenId] = idx + 1;
        }

        _myOpenedBoxes[user].pop();
        delete _myOpenedPosPlusOne[user][tokenId];
    }

    // ============================================================
    //                  Seller sold list (pagination: BoxInfo[]，带 tokenId)
    // ============================================================

    // ============================================================
    //              Bottom-layer hook: 分清 mint/transfer/burn
    //              （不需要 economicHolder）
    // ============================================================

    function _update(address to, uint256 tokenId, address auth)
        internal
        override(ERC721)
        returns (address from)
    {
        from = super._update(to, tokenId, auth);

        // Burn（无外部 burn 接口，保留底层处理以支持 transfer 到 address(0)）
        if (to == address(0)) {
            if (open[tokenId]) {
                if (from != address(0)) {
                    _myOpenedRemove(from, tokenId);
                }
            } else {
                // 未开盒：完整保护（正常情况下你不会让未开盒 burn）
                address holder = address(0);

                if (from == address(this)) {
                    Listing memory lst = listings[tokenId];
                    if (lst.active) holder = lst.seller;
                } else {
                    holder = from;
                }

                if (holder != address(0)) {
                    _holdingRemoveBox(holder, boxInfo[tokenId].usdtPaid);
                    _myUnopenedRemove(holder, tokenId);
                }
            }
            return from;
        }

        // Opened token: 只迁移 opened 列表；holding 不动
        if (open[tokenId]) {
            if (from == address(0)) {
                _myOpenedAdd(to, tokenId);
                return from;
            }
            if (from != to) {
                _myOpenedRemove(from, tokenId);
                _myOpenedAdd(to, tokenId);
            }
            return from;
        }

        // Unopened token
        uint256 base = boxInfo[tokenId].usdtPaid;

        // Mint
        if (from == address(0)) {
            _holdingAddBox(to, base);
            _myUnopenedAdd(to, tokenId);
            return from;
        }

        // Escrow deposit: user -> this (list)，不动 holding/列表
        if (to == address(this)) {
            return from;
        }

        // Escrow out: this -> someone
        if (from == address(this)) {
            Listing memory lst = listings[tokenId];
            if (lst.active) {
                // unlist: back to seller，不迁移
                if (to == lst.seller) {
                    return from;
                }
                // sold: seller -> buyer
                _holdingRemoveBox(lst.seller, base);
                _myUnopenedRemove(lst.seller, tokenId);

                _holdingAddBox(to, base);
                _myUnopenedAdd(to, tokenId);

                return from;
            }
            return from;
        }

        // Normal transfer
        _holdingRemoveBox(from, base);
        _myUnopenedRemove(from, tokenId);

        _holdingAddBox(to, base);
        _myUnopenedAdd(to, tokenId);

        return from;
    }

    // ============================================================
    //                  Pagination reads returning BoxInfo (global)
    //                  （分页只返回 BoxInfo[]，带 tokenId）
    // ============================================================

    function mintedTokenIdsLength() external view returns (uint256) {
        return mintedTokenIds.length;
    }

    function getMintedByPage(uint256 start, uint256 count)
        external
        view
        returns (BoxInfo[] memory infos)
    {
        uint256 len = mintedTokenIds.length;
        if (start >= len) return new BoxInfo[](0);

        uint256 end = start + count;
        if (end > len) end = len;

        uint256 n = end - start;
        infos = new BoxInfo[](n);

        for (uint256 i = 0; i < n; i++) {
            uint256 tokenId = mintedTokenIds[start + i];
            BoxInfo memory bi = boxInfo[tokenId];
            bi.tokenId = tokenId;
            infos[i] = bi;
        }
    }

    // ============================================================
    //              Global Open Action Records + Pagination
    //              OpenActionTokenIdsLength / GetOpenedByPage
    // ============================================================

    function OpenActionTokenIdsLength() external view returns (uint256) {
        return openActionTokenIds.length;
    }

    function GetOpenedByPage(uint256 start, uint256 count)
        external
        view
        returns (BoxInfo[] memory infos)
    {
        uint256 len = openActionTokenIds.length;
        if (start >= len) return new BoxInfo[](0);

        uint256 end = start + count;
        if (end > len) end = len;

        uint256 n = end - start;
        infos = new BoxInfo[](n);

        for (uint256 i = 0; i < n; i++) {
            uint256 tokenId = openActionTokenIds[start + i];
            BoxInfo memory bi = boxInfo[tokenId];
            bi.tokenId = tokenId; // 保险补齐
            infos[i] = bi;
        }
    }

    // ============================================================
    //                         Other views
    // ============================================================

    function openUsageNow(address user)
        external
        view
        returns (uint256 dayIndex, uint256 used, uint256 limit, uint256 remaining)
    {
        dayIndex = _dayIndex(block.timestamp);
        Usage memory u = _usage[user];

        used = (u.dayIndex == dayIndex) ? uint256(u.used) : 0;
        limit = _baseDailyOpensOf(user) + extraOpenCount[user];
        remaining = (limit > used) ? (limit - used) : 0;
    }

    function secondaryBuyUsageNow(address user)
        external
        view
        returns (uint256 dayIndex, uint256 used, uint256 limit, uint256 remaining)
    {
        dayIndex = _dayIndex(block.timestamp);
        Usage memory u = _secondaryBuyUsage[user];

        used = (u.dayIndex == dayIndex) ? uint256(u.used) : 0;
        limit = _baseDailyOpensOf(user) + extraOpenCount[user];
        remaining = (limit > used) ? (limit - used) : 0;
    }

    /// @notice 仓位已无限；保留接口兼容，恒为 0
    /// @param step 保留返回值兼容，恒为 0
    function previewSlotCost(address user, uint256 count) external pure returns (uint256 cost, uint256 step) {
        if (count == 0) revert ErrCount();
        user;
        count;
        cost = 0;
        step = 0;
    }

    function _openCountCostInB(uint256 currentCount, uint256 count) internal view returns (uint256 costB) {
        if (openCountBCostFixedEnabled) {
            if (openCountFirstCostB == 0) revert ErrZero();
            return count * openCountFirstCostB + openCountStepCostB * count * (2 * currentCount + count - 1) / 2;
        }
        if (address(feeTokenB) == address(0)) revert ErrZero();
        if (address(feeQuotePool) == address(0)) revert ErrPool();
        uint256 costUSDT = _slotCost(currentCount, count);
        return quoteFeeInB(costUSDT);
    }

    function _salePayoutRecipient(address seller) internal view returns (address recipient) {
        recipient = seller;
        if (seller.code.length == 0) {
            return recipient;
        }
        try ISalePayoutRecipient(seller).payoutRecipient() returns (address payout) {
            if (payout != address(0)) {
                recipient = payout;
            }
        } catch {}
    }

    function withdrawERC20(address token, address to, uint256 amount)
        external
        onlyRole(SUPER_ADMIN_ROLE)
        nonReentrant
    {
        if (to == address(0)) revert ErrZero();
        IERC20(token).safeTransfer(to, amount);
        emit ERC20Withdrawn(token, to, amount);
    }

    /// @notice 将 NFT 合约当前持有的 USDT 转到量化地址
    /// @dev 这是预留接口，只处理“已经在合约里的资金”，不会改变 mint 时 95%/5% 的现有分账路径
    function transferHeldUsdtToQuant(uint256 amount, string calldata bizNo)
        external
        onlyRole(SUPER_ADMIN_ROLE)
        nonReentrant
    {
        if (quantReceiver == address(0)) revert ErrReceiver();
        if (amount == 0) revert ErrZero();
        usdt.safeTransfer(quantReceiver, amount);
        emit QuantTransferExecuted(address(usdt), quantReceiver, amount, msg.sender, bizNo);
    }

    /// @notice 将 NFT 合约当前持有的任意 ERC20 转到量化地址
    /// @dev 预留给后续扩展，不建议替代通用 withdrawERC20 的人工运维用途
    function transferHeldERC20ToQuant(address token, uint256 amount, string calldata bizNo)
        external
        onlyRole(SUPER_ADMIN_ROLE)
        nonReentrant
    {
        if (token == address(0)) revert ErrToken();
        if (quantReceiver == address(0)) revert ErrReceiver();
        if (amount == 0) revert ErrZero();
        IERC20(token).safeTransfer(quantReceiver, amount);
        emit QuantTransferExecuted(token, quantReceiver, amount, msg.sender, bizNo);
    }

    /// @notice 购买保险占位接口，当前仅预留 ABI，具体逻辑后续实现
    function buyInsurance(uint256 tokenId, uint256 amountB) external pure {
        tokenId;
        amountB;
        revert ErrNotImplemented();
    }

    /// @notice 购买道具占位接口，当前仅预留 ABI，具体逻辑后续实现
    function buyProp(uint256 propType, uint256 amountB) external pure {
        propType;
        amountB;
        revert ErrNotImplemented();
    }

    // -------------------- Interface --------------------
    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721, AccessControl)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }
}
