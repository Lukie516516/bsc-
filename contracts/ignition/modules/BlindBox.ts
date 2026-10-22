/**
 * BSC 主网盲盒合约部署
 *
 * 部署前设置环境变量：
 *   BSC_PRIVATE_KEY=你的私钥hex无0x
 *   BSC_RPC_URL=https://bsc-dataseed.binance.org/  (可选，有默认)
 *
 * 可选覆盖（不设则用 deployer）：
 *   DEPLOY_INITIAL_RECIPIENT=Piece代币接收地址
 *   DEPLOY_ACCOUNT_A=铸造分账收款
 *   DEPLOY_ACCOUNT_B=槽位购买收款
 *   DEPLOY_ACCOUNT_C=二级市场USDT手续费
 *   DEPLOY_ACCOUNT_D=二级市场HRX手续费
 *
 * 运行：npx hardhat ignition deploy ./ignition/modules/BlindBox.ts --network bsc
 */

import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

// BSC 主网常量
const USDT = "0x55d398326f99059fF775485246999027B3197955";
const WBNB = "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c";

const PIECE_SUPPLY = 1171350625n * 10n ** 18n;
const DEFAULT_MIN_STAKE_A = 300n * 10n ** 18n;

export default buildModule("BlindBox", (m) => {
  const deployer = m.getAccount(0);

  // 可选：部署时传参 --parameters '{"initialRecipient":"0x..."}' 覆盖
  // 不传则用 deployer
  const initialRecipient = m.getParameter("initialRecipient", deployer);
  const accountA = m.getParameter("accountA", deployer);
  const accountB = m.getParameter("accountB", deployer);
  const accountC = m.getParameter("accountC", deployer);
  const accountD = m.getParameter("accountD", deployer);
  const minStakeA = m.getParameter("minStakeA", DEFAULT_MIN_STAKE_A);

  // 0. NftPriceLib + NftFeeLib（Nft 依赖，减小体积）
  const nftPriceLib = m.library("NftPriceLib");
  const nftFeeLib = m.library("NftFeeLib");

  // 1. Piece (HRX 持有价值代币)
  const piece = m.contract("Piece", [
    "Piece",
    "PIECE",
    initialRecipient,
    PIECE_SUPPLY,
  ]);

  // 2. Hrx (一级市场 bonding curve 代币)
  const hrx = m.contract("Hrx", ["Hrx", "HRX"]);

  // 3. Market (一级市场)
  const market = m.contract("BondingCurvePrimaryMarket", [
    USDT,
    hrx,
    deployer,
    accountA,
  ]);

  // 4. Stake (质押分红)
  const stake = m.contract("UmbrellaStakeDividend", [
    hrx,
    USDT,
    accountA,
    minStakeA,
  ]);

  // 5. Nft (盲盒，依赖 NftPriceLib)
  const nft = m.contract(
    "BlindBoxNFT",
    [
      "BlindBox",
      "BOX",
      USDT,
      hrx,
      accountA,
      accountB,
      accountC,
      accountD,
      0,
      stake,
    ],
    { libraries: { NftPriceLib: nftPriceLib, NftFeeLib: nftFeeLib } }
  );

  // 6. SwapHrxUsdt (HRX/USDT AMM)
  const swap = m.contract("SwapHrxUsdt", [
    hrx,
    USDT,
    deployer,
    accountA,
    3,
    100,
  ]);

  // 7. Router
  const router = m.contract("DealRouter", [WBNB]);

  // 8. ListTemplate (BlindBoxUserVault 实现)
  const listTemplate = m.contract("BlindBoxUserVault");

  // 9. FactoryList
  const factoryList = m.contract("BlindBoxVaultFactory", [
    listTemplate,
    nft,
    USDT,
  ]);

  // 10. StakeTwo (本金可提)
  const stakeTwo = m.contract("BlindBoxPrincipalVault", [nft]);

  // --- 后续配置 ---

  // SwapHrxUsdt 设置 primaryMarket
  m.call(swap, "setPrimaryMarket", [market]);

  // Nft 设置 feeQuotePool、pieceToken、出售分佣 Reader
  m.call(nft, "setFeeQuotePool", [swap]);
  m.call(nft, "setPieceToken", [piece]);
  m.call(nft, "setPieceValueUsdtE18", [1n * 10n ** 18n]); // 1 Piece = 1 USDT
  m.call(nft, "setStakeReferralReader", [stake]); // 二级市场出售：B 0.3% + D 0.2%

  // Piece 设置 nft
  m.call(piece, "setNFT", [nft]);

  // Stake 设置 nft 和 piece
  m.call(stake, "setNftContract", [nft]);
  m.call(stake, "setPieceToken", [piece]);

  return {
    nftPriceLib,
    nftFeeLib,
    piece,
    hrx,
    market,
    stake,
    nft,
    swap,
    router,
    listTemplate,
    factoryList,
    stakeTwo,
  };
});
