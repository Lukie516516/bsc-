// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/math/Math.sol";

/// @notice NFT 合约需要提供的接口：查询某地址“盲盒持有价值(USDT,18位)”
interface IBlindBoxValue {
    function holdingValueOf(address user) external view returns (uint256);
}

/// @title UmbrellaStakeDividend
/// @notice
/// - 推荐网：parentOf[user]、childrenOf[parent]、topOf[user]
/// - 质押：stake A（18位），minStakeA 必须为非 0 的有效门槛
/// - setReward：仅 NFT 合约可调用，传入 holder + profit(USDT)
/// - 分红：profit 的 refProfitBps% 进入“分配池”，其余归用户可提
///   分配池：
///   - 先由 1级上级B的“持有价值”决定给(B,C)的比例；剩余给(项目方+顶级)平分(顶级需满足条件)
///   - (B,C)之间按 (BValue*2) : (CValue*1) 分配（权重倍数可调）
/// - 顶级分红条件：
///   1) topEnabled[top]==true（管理员开关）
///   2) stakedA[top] >= minStakeA
///   3) 或管理员通过 setTopManualEligible(top, true) 人工放行
///
/// @dev 本版本：注册时不检查伞顶是否达到 minStakeA，也不会自动授权 topEnabled；
///      在 setReward 分配时才检查 top 是否合格，不合格则 top 不拿对应 top 份额（该份额归项目方）。
contract UmbrellaStakeDividend is AccessControl, ReentrancyGuard {
    using SafeERC20 for IERC20;

    bytes32 public constant SUPER_ADMIN_ROLE = keccak256("SUPER_ADMIN_ROLE");
    bytes32 public constant MANAGER_ROLE = keccak256("MANAGER_ROLE");
    bytes32 public constant MIGRATOR_ROLE = keccak256("MIGRATOR_ROLE");

    IERC20 public immutable aToken;   // A token (18 decimals)
    IERC20 public immutable usdt;     // USDT (18 decimals)

    /// @notice 项目方地址（拿项目份额，也可 claim）
    address public projectAddress;

    /// @notice 唯一允许调用 setReward 的 NFT 合约地址
    address public nftContract;

    /// @notice 分红使用的 NFT 合约接口（用于查 B/C 的持有价值）
    IBlindBoxValue public nftValueReader;

    // -------------------- staking --------------------
    uint256 public minStakeA; // 最低质押数量（18位，必须 > 0）
    mapping(address => uint256) public stakedA;
    uint256 public totalStakedA;

    // -------------------- Piece staking (提升 B/C 分红权重) --------------------
    IERC20 public pieceToken;
    mapping(address => uint256) public stakedPiece;
    uint256 public totalStakedPiece;

    // -------------------- referral graph --------------------
    mapping(address => bool) public registered;
    mapping(address => address) public parentOf; // 直属上级
    mapping(address => address) public topOf;    // 顶级(伞顶)

    mapping(address => address[]) private _childrenOf; // 下级数组

    function childrenLength(address u) external view returns (uint256) {
        return _childrenOf[u].length;
    }

    function getChildrenByPage(
        address u,
        uint256 start,
        uint256 count
    ) external view returns (address[] memory out) {
        uint256 len = _childrenOf[u].length;
        if (start >= len) return new address[](0);
        uint256 end = start + count;
        if (end > len) end = len;

        out = new address[](end - start);
        for (uint256 i = start; i < end; i++) {
            out[i - start] = _childrenOf[u][i];
        }
    }

    // -------------------- top eligibility --------------------
    mapping(address => bool) public topEnabled; // 管理员后台授权开关
    mapping(address => bool) public topManualEligible; // 管理员可直接授权团队长资格

    function isTopEligible(address top) public view returns (bool) {
        if (top == address(0)) return false;
        if (!topEnabled[top]) return false;
        if (topManualEligible[top]) return true;
        return stakedA[top] >= minStakeA;
    }

    /// @notice 运营/后端可直接读取团队长资格状态，明确看到是“达标”还是“人工授权”
    function getTopEligibilityStatus(address top)
        external
        view
        returns (
            bool enabled,
            bool manualEligible,
            uint256 stakedAmount,
            uint256 requiredStake,
            bool eligible
        )
    {
        enabled = topEnabled[top];
        manualEligible = topManualEligible[top];
        stakedAmount = stakedA[top];
        requiredStake = minStakeA;
        eligible = isTopEligible(top);
    }

    // -------------------- reward accounting (USDT) --------------------
    mapping(address => uint256) public pendingUsdt;
    mapping(address => uint256) public totalEarnedUsdt;

    // -------------------- configs --------------------
    uint256 public bpsDenominator = 10_000;

    /// @notice profit 的多少进入“分配池”（默认 10%）
    uint256 public refProfitBps = 1_000; // 10%

    /// @notice B(1级上级)持有价值阈值（USDT, 18位）
    uint256 public holdValueMin = 300e18;
    uint256 public holdValueMax = 10_000e18;

    /// @notice 当 B 从 holdValueMin 到 holdValueMax 时，给(B,C)的占利润比例从 2% 线性到 10%（占池子 20%~100%，池子=10%利润）
    uint256 public minUplinePortionBps = 2_000;  // 占池子 20% = 利润 2% (B=300u)
    uint256 public maxUplinePortionBps = 10_000; // 占池子 100% = 利润 10% (B=10000u)

    /// @notice (B,C)权重倍数（示例：B*2, C*1）
    uint256 public level1WeightMul = 2;
    uint256 public level2WeightMul = 1;

    bool public autoSettleProjectPending;
    bool public migrationClosed;

    // -------------------- events --------------------
    event ProjectAddressChanged(address indexed oldAddr, address indexed newAddr);
    event NftContractChanged(address indexed oldNft, address indexed newNft);

    event Registered(address indexed user, address indexed parent, address indexed top);
    event MigrationClosed(address indexed operator);
    event TopEnabledSet(address indexed top, bool enabled);
    event TopManualEligibleSet(address indexed top, bool enabled);

    event MinStakeSet(uint256 newMinStake);
    event Staked(address indexed user, uint256 amount, uint256 newTotal);
    event Unstaked(address indexed user, uint256 amount, uint256 newTotal);

    event PieceTokenSet(address indexed oldPiece, address indexed newPiece);
    event PieceStaked(address indexed user, uint256 amount, uint256 newTotal);
    event PieceUnstaked(address indexed user, uint256 amount, uint256 newTotal);

    event ConfigBpsDenominator(uint256 denom);
    event ConfigRefProfit(uint256 refProfitBps);
    event ProjectAutoSettlementSet(bool enabled);
    event ConfigHoldRange(uint256 minValue, uint256 maxValue, uint256 minUplinePortionBps, uint256 maxUplinePortionBps);
    event ConfigWeightMul(uint256 level1Mul, uint256 level2Mul);

    event RewardNotified(
        address indexed user,
        uint256 profit,
        uint256 userShare,
        address indexed l1,
        address indexed l2,
        address top,
        uint256 pool,
        uint256 uplinePortionBps,
        uint256 toL1,
        uint256 toL2,
        uint256 toTop,
        uint256 toProject
    );

    event Claimed(address indexed user, address indexed to, uint256 amount);

    constructor(
        address aToken_,
        address usdt_,
        address project_,
        uint256 minStakeA_
    ) {
        require(aToken_ != address(0) && usdt_ != address(0), "ZERO_TOKEN");
        require(project_ != address(0), "ZERO_PROJECT");
        require(minStakeA_ > 0, "MIN_STAKE_ZERO");

        aToken = IERC20(aToken_);
        usdt = IERC20(usdt_);
        projectAddress = project_;
        minStakeA = minStakeA_;

        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(SUPER_ADMIN_ROLE, msg.sender);
        _grantRole(MANAGER_ROLE, msg.sender);
        _grantRole(MIGRATOR_ROLE, msg.sender);
    }

    // -------------------- admin setters --------------------

    function setProjectAddress(address p) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(p != address(0), "ZERO_PROJECT");
        address old = projectAddress;
        projectAddress = p;
        emit ProjectAddressChanged(old, p);
    }

    /// @notice 设置允许调用 setReward 的 NFT 合约，并作为持有价值读取源
    function setNftContract(address nft) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(nft != address(0), "ZERO_NFT");
        address old = nftContract;
        nftContract = nft;
        nftValueReader = IBlindBoxValue(nft);
        emit NftContractChanged(old, nft);
    }

    function setMinStakeA(uint256 m) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(m > 0, "MIN_STAKE_ZERO");
        minStakeA = m;
        emit MinStakeSet(m);
    }

    function setPieceToken(address piece_) external onlyRole(DEFAULT_ADMIN_ROLE) {
        // 禁止设为 0：否则已质押用户无法 unstake，资金会被锁
        require(piece_ != address(0), "ZERO_PIECE");
        address old = address(pieceToken);
        pieceToken = IERC20(piece_);
        emit PieceTokenSet(old, piece_);
    }

    function setTopEnabled(address top, bool enabled) external onlyRole(MANAGER_ROLE) {
        require(top != address(0), "ZERO_TOP");
        topEnabled[top] = enabled;
        emit TopEnabledSet(top, enabled);
    }

    /// @notice 管理员可直接授权团队长资格，作为未达质押门槛时的例外放行
    function setTopManualEligible(address top, bool enabled) external onlyRole(MANAGER_ROLE) {
        require(top != address(0), "ZERO_TOP");
        topManualEligible[top] = enabled;
        emit TopManualEligibleSet(top, enabled);
    }

    function setBpsDenominator(uint256 denom) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(denom > 0, "DENOM=0");
        bpsDenominator = denom;
        emit ConfigBpsDenominator(denom);
    }

    function setRefProfitBps(uint256 bps) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(bps <= bpsDenominator, "BPS_TOO_BIG");
        refProfitBps = bps;
        emit ConfigRefProfit(bps);
    }

    /// @notice 项目方份额是否自动结算到 projectAddress；关闭则继续累计在 pendingUsdt[projectAddress]
    function setProjectAutoSettlement(bool enabled) external onlyRole(DEFAULT_ADMIN_ROLE) {
        autoSettleProjectPending = enabled;
        emit ProjectAutoSettlementSet(enabled);
    }

    function closeMigration() external onlyRole(DEFAULT_ADMIN_ROLE) {
        migrationClosed = true;
        emit MigrationClosed(msg.sender);
    }

    function setHoldValueConfig(uint256 minV, uint256 maxV, uint256 minPortionBps, uint256 maxPortionBps)
        external
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        require(maxV > 0 && maxV >= minV, "BAD_RANGE");
        require(minPortionBps <= bpsDenominator && maxPortionBps <= bpsDenominator, "BAD_PORTION");
        require(minPortionBps <= maxPortionBps, "MIN_GT_MAX");
        holdValueMin = minV;
        holdValueMax = maxV;
        minUplinePortionBps = minPortionBps;
        maxUplinePortionBps = maxPortionBps;
        emit ConfigHoldRange(minV, maxV, minPortionBps, maxPortionBps);
    }

    function setWeightMultipliers(uint256 l1Mul, uint256 l2Mul) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(l1Mul > 0 && l2Mul > 0, "MUL=0");
        level1WeightMul = l1Mul;
        level2WeightMul = l2Mul;
        emit ConfigWeightMul(l1Mul, l2Mul);
    }

    // -------------------- register --------------------

    /// @notice 注册：parent 为上级地址；若 parent==0 则自己成为顶级（伞顶），但仍需管理员授权 topEnabled
    /// @dev 本版本：注册不再检查伞顶是否达到 minStakeA；顶级分红资格在 setReward 时判断
    function register(address parent) external {
        _registerReferral(msg.sender, parent);
    }

    /// @notice 迁移期导入单个旧会员推荐关系；关闭迁移后不可再调用。
    function importReferral(address user, address parent) external onlyRole(MIGRATOR_ROLE) {
        require(!migrationClosed, "MIGRATION_CLOSED");
        _registerReferral(user, parent);
    }

    /// @notice 迁移期批量导入旧会员推荐关系；users[i] 对应 parents[i]。
    function importReferralBatch(address[] calldata users, address[] calldata parents) external onlyRole(MIGRATOR_ROLE) {
        require(!migrationClosed, "MIGRATION_CLOSED");
        require(users.length == parents.length, "LENGTH_MISMATCH");

        for (uint256 i = 0; i < users.length; i++) {
            _registerReferral(users[i], parents[i]);
        }
    }

    function _registerReferral(address user, address parent) internal {
        require(user != address(0), "ZERO_USER");
        require(!registered[user], "ALREADY_REGISTERED");
        require(parent != user, "PARENT_SELF");

        if (parent == address(0)) {
            registered[user] = true;
            parentOf[user] = address(0);
            topOf[user] = user;

            emit Registered(user, address(0), user);
            return;
        }

        require(registered[parent], "PARENT_NOT_REGISTERED");
        address top = topOf[parent];
        require(top != address(0), "BAD_TOP");

        registered[user] = true;
        parentOf[user] = parent;
        topOf[user] = top;

        _childrenOf[parent].push(user);

        emit Registered(user, parent, top);
    }

    // -------------------- staking A --------------------

    function stake(uint256 amount) external nonReentrant {
        require(amount > 0, "AMOUNT=0");
        aToken.safeTransferFrom(msg.sender, address(this), amount);
        stakedA[msg.sender] += amount;
        totalStakedA += amount;
        emit Staked(msg.sender, amount, stakedA[msg.sender]);
    }

    function unstake(uint256 amount) external nonReentrant {
        require(amount > 0, "AMOUNT=0");
        uint256 bal = stakedA[msg.sender];
        require(bal >= amount, "INSUFFICIENT_STAKE");
        stakedA[msg.sender] = bal - amount;
        totalStakedA -= amount;
        aToken.safeTransfer(msg.sender, amount);
        emit Unstaked(msg.sender, amount, stakedA[msg.sender]);
    }

    // -------------------- staking Piece（提升 B/C 分红权重） --------------------

    function stakePiece(uint256 amount) external nonReentrant {
        require(amount > 0, "AMOUNT=0");
        require(address(pieceToken) != address(0), "PIECE_NOT_SET");
        pieceToken.safeTransferFrom(msg.sender, address(this), amount);
        stakedPiece[msg.sender] += amount;
        totalStakedPiece += amount;
        emit PieceStaked(msg.sender, amount, stakedPiece[msg.sender]);
    }

    function unstakePiece(uint256 amount) external nonReentrant {
        require(amount > 0, "AMOUNT=0");
        uint256 bal = stakedPiece[msg.sender];
        require(bal >= amount, "INSUFFICIENT_STAKE");
        stakedPiece[msg.sender] = bal - amount;
        totalStakedPiece -= amount;
        pieceToken.safeTransfer(msg.sender, amount);
        emit PieceUnstaked(msg.sender, amount, stakedPiece[msg.sender]);
    }

    // -------------------- claim USDT --------------------

    function claimAll(address to) external nonReentrant returns (uint256 claimed) {
        require(to != address(0), "TO=0");
        claimed = pendingUsdt[msg.sender];
        require(claimed > 0, "NOTHING");
        pendingUsdt[msg.sender] = 0;
        usdt.safeTransfer(to, claimed);
        emit Claimed(msg.sender, to, claimed);
    }

    function claim(address to, uint256 amount) external nonReentrant {
        require(to != address(0), "TO=0");
        require(amount > 0, "AMOUNT=0");
        uint256 p = pendingUsdt[msg.sender];
        require(p >= amount, "INSUFFICIENT_PENDING");
        pendingUsdt[msg.sender] = p - amount;
        usdt.safeTransfer(to, amount);
        emit Claimed(msg.sender, to, amount);
    }

    /// @notice 将项目方已沉淀的 pending 直接结算到项目方地址
    function settleProjectPending() external nonReentrant returns (uint256 settled) {
        settled = pendingUsdt[projectAddress];
        require(settled > 0, "NOTHING");
        pendingUsdt[projectAddress] = 0;
        usdt.safeTransfer(projectAddress, settled);
        emit Claimed(projectAddress, projectAddress, settled);
    }

    // -------------------- core: setReward --------------------

    modifier onlyNft() {
        require(msg.sender == nftContract, "ONLY_NFT");
        _;
    }

    /// @notice NFT 合约调用：通知某 user 本次盈利 profit(USDT 18位)
    /// @dev 这里仅做“记账分配”，不从 NFT 拉钱；USDT 应由 NFT mint 转入本合约等方式保证余额充足
    function setReward(address user, uint256 profit) external onlyNft nonReentrant {
        require(user != address(0), "USER=0");
        require(profit > 0, "PROFIT=0");
        require(projectAddress != address(0), "PROJECT=0");
        require(address(nftValueReader) != address(0), "NFT_READER=0");

        // 如果用户没注册，则默认他是一个新的顶级伞（避免 setReward 失败）
        if (!registered[user]) {
            _registerReferral(user, address(0));
        }

        uint256 pool = Math.mulDiv(profit, refProfitBps, bpsDenominator);
        uint256 userShare = profit - pool;

        _addPending(user, userShare);

        address l1 = parentOf[user];
        address l2 = (l1 == address(0)) ? address(0) : parentOf[l1];
        address top = topOf[user];
        if (top == address(0)) top = user;

        // 给(B,C)的占比：由 B 持有价值决定，占池子的 20%~100%（=利润的 2%~10%）
        // B<300u=>0%，B=300u=>20%池子，B=10000u=>100%池子；分完无剩余则 D/项目方无
        uint256 l1Value = (l1 == address(0)) ? 0 : nftValueReader.holdingValueOf(l1);
        uint256 uplinePortionBps = _calcUplinePortionBps(l1Value);

        uint256 uplinesPart = Math.mulDiv(pool, uplinePortionBps, bpsDenominator);
        uint256 rest = pool - uplinesPart; // 剩余再分给 D 和项目方各 50%

        uint256 toL1 = 0;
        uint256 toL2 = 0;

        // 权重：(B的NFT价值+B的Piece质押)*2 : (C的NFT价值+C的Piece质押)*1
        if (uplinesPart > 0 && l1 != address(0)) {
            uint256 l2Value = (l2 == address(0)) ? 0 : nftValueReader.holdingValueOf(l2);
            uint256 l1Piece = stakedPiece[l1];
            uint256 l2Piece = (l2 == address(0)) ? 0 : stakedPiece[l2];

            uint256 w1 = (l1Value + l1Piece) * level1WeightMul;
            uint256 w2 = (l2Value + l2Piece) * level2WeightMul;
            uint256 wSum = w1 + w2;

            if (wSum == 0) {
                rest += uplinesPart;
            } else {
                toL1 = Math.mulDiv(uplinesPart, w1, wSum);
                toL2 = uplinesPart - toL1;

                if (toL1 > 0) _addPending(l1, toL1);
                if (toL2 > 0 && l2 != address(0)) {
                    _addPending(l2, toL2);
                } else if (toL2 > 0) {
                    // 没有 l2，则给项目方
                    _addProjectPending(toL2);
                    toL2 = 0;
                }
            }
        } else if (uplinesPart > 0) {
            // 没有 l1，则全并入 rest
            rest += uplinesPart;
        }

        uint256 toTop = 0;
        uint256 toProject = 0;

        if (rest > 0) {
            // ✅ 分红时检查 top 是否合格：不合格则 top 不拿 top 份额
            if (isTopEligible(top)) {
                toTop = rest / 2;
                toProject = rest - toTop;
                if (toTop > 0) _addPending(top, toTop);
                if (toProject > 0) _addProjectPending(toProject);
            } else {
                toProject = rest;
                _addProjectPending(toProject);
            }
        }

        emit RewardNotified(
            user,
            profit,
            userShare,
            l1,
            l2,
            top,
            pool,
            uplinePortionBps,
            toL1,
            toL2,
            toTop,
            toProject
        );
    }

    function _addPending(address u, uint256 amt) internal {
        if (amt == 0) return;
        pendingUsdt[u] += amt;
        totalEarnedUsdt[u] += amt;
    }

    function _addProjectPending(uint256 amt) internal {
        if (amt == 0) return;
        if (autoSettleProjectPending) {
            totalEarnedUsdt[projectAddress] += amt;
            usdt.safeTransfer(projectAddress, amt);
            emit Claimed(projectAddress, projectAddress, amt);
            return;
        }
        _addPending(projectAddress, amt);
    }

    /// @dev 规则：
    /// - 若 BValue < holdValueMin => 0%
    /// - B 从 300u 到 10000u：给(B,C)的比例从 2% 线性到 10%
    function _calcUplinePortionBps(uint256 bValue) internal view returns (uint256) {
        if (bValue < holdValueMin) return 0;
        if (holdValueMax <= holdValueMin) return minUplinePortionBps;

        if (bValue >= holdValueMax) return maxUplinePortionBps;

        uint256 range = holdValueMax - holdValueMin;
        uint256 bpsRange = maxUplinePortionBps - minUplinePortionBps;
        return minUplinePortionBps + Math.mulDiv(bValue - holdValueMin, bpsRange, range);
    }

    // -------------------- profile (avatar & nickname) --------------------
    struct Profile {
        string nickname; // 昵称
        string avatar;   // 头像（可存 URL / IPFS CID 等）
    }

    /// @notice 地址 -> 资料（公开可读）
    mapping(address => Profile) public profileOf;

    /// @dev 防止有人用超长字符串恶意占用存储/导致前端卡顿（可按需调整或删除）
    uint256 public constant MAX_NICKNAME_LEN = 32;   // 32 bytes
    uint256 public constant MAX_AVATAR_LEN   = 256;  // 256 bytes

    event ProfileUpdated(address indexed user, string nickname, string avatar);

    /// @notice 一次性设置昵称+头像（传空字符串可视为清空）
    function setProfile(string calldata nickname, string calldata avatar) external {
        require(bytes(nickname).length <= MAX_NICKNAME_LEN, "NICK_TOO_LONG");
        require(bytes(avatar).length <= MAX_AVATAR_LEN, "AVATAR_TOO_LONG");

        profileOf[msg.sender] = Profile({nickname: nickname, avatar: avatar});
        emit ProfileUpdated(msg.sender, nickname, avatar);
    }

    /// @notice 只改昵称
    function setNickname(string calldata nickname) external {
        require(bytes(nickname).length <= MAX_NICKNAME_LEN, "NICK_TOO_LONG");
        profileOf[msg.sender].nickname = nickname;
        emit ProfileUpdated(msg.sender, profileOf[msg.sender].nickname, profileOf[msg.sender].avatar);
    }

    /// @notice 只改头像
    function setAvatar(string calldata avatar) external {
        require(bytes(avatar).length <= MAX_AVATAR_LEN, "AVATAR_TOO_LONG");
        profileOf[msg.sender].avatar = avatar;
        emit ProfileUpdated(msg.sender, profileOf[msg.sender].nickname, profileOf[msg.sender].avatar);
    }

    function withdrawERC20(address token, address to, uint256 amount)
        external
        onlyRole(SUPER_ADMIN_ROLE)
        nonReentrant
    {
        require(to != address(0), "to=0");
        IERC20(token).safeTransfer(to, amount);
    }
}
