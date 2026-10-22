/**
 * BSC 测试网盲盒合约部署（含 MockUSDT）
 *
 * 部署前：
 *   1. 从 BSC 测试网水龙头获取 tBNB: https://faucets.chain.link/bnb-chain-chapel
 *   2. 设置 box/.env: BSC_PRIVATE_KEY=你的私钥
 *
 * 运行：
 *   npx hardhat ignition deploy ./ignition/modules/BlindBoxTestnet.ts --network bscTestnet
 *
 * 部署完成后可调用 usdt.mint(地址, 数量) 给测试地址打 U
 */

import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

// BSC 测试网 WBNB
const WBNB_TESTNET = "0xae13d989daC2f0dEbFf460aC112a837C89BAa7cd";

const PIECE_SUPPLY = 1171350625n * 10n ** 18n;
const MOCK_USDT_INITIAL = 10_000_000n * 10n ** 18n; // 1000 万测试 U
const DEFAULT_MIN_STAKE_A = 300n * 10n ** 18n;

export default buildModule("BlindBoxTestnet", (m) => {
  const deployer = m.getAccount(0);

  const initialRecipient = m.getParameter("initialRecipient", deployer);
  const accountA = m.getParameter("accountA", deployer);
  const accountB = m.getParameter("accountB", deployer);
  const accountC = m.getParameter("accountC", deployer);
  const accountD = m.getParameter("accountD", deployer);
  const minStakeA = m.getParameter("minStakeA", DEFAULT_MIN_STAKE_A);

  // 0. 库
  const nftPriceLib = m.library("NftPriceLib");
  const nftFeeLib = m.library("NftFeeLib");

  // 1. MockUSDT（测试用，可 mint）
  const usdt = m.contract("MockUSDT", [deployer, MOCK_USDT_INITIAL]);

  // 2. Piece
  const piece = m.contract("Piece", [
    "Piece",
    "PIECE",
    initialRecipient,
    PIECE_SUPPLY,
  ]);

  // 3. Hrx
  const hrx = m.contract("Hrx", ["Hrx", "HRX"]);

  // 4. Market（用 MockUSDT）
  const market = m.contract("BondingCurvePrimaryMarket", [
    usdt,
    hrx,
    deployer,
    accountA,
  ]);

  // 5. Stake
  const stake = m.contract("UmbrellaStakeDividend", [
    hrx,
    usdt,
    accountA,
    minStakeA,
  ]);

  // 6. Nft
  const nft = m.contract(
    "BlindBoxNFT",
    [
      "BlindBox",
      "BOX",
      usdt,
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

  // 7. SwapHrxUsdt
  const swap = m.contract("SwapHrxUsdt", [
    hrx,
    usdt,
    deployer,
    accountA,
    3,
    100,
  ]);

  // 8. Router（测试网 WBNB）
  const router = m.contract("DealRouter", [WBNB_TESTNET]);

  // 9. ListTemplate
  const listTemplate = m.contract("BlindBoxUserVault");

  // 10. FactoryList
  const factoryList = m.contract("BlindBoxVaultFactory", [
    listTemplate,
    nft,
    usdt,
  ]);

  // 11. StakeTwo
  const stakeTwo = m.contract("BlindBoxPrincipalVault", [nft]);

  // --- 配置 ---
  m.call(swap, "setPrimaryMarket", [market]);
  m.call(nft, "setFeeQuotePool", [swap]);
  m.call(nft, "setPieceToken", [piece]);
  m.call(nft, "setPieceValueUsdtE18", [1n * 10n ** 18n]);
  m.call(nft, "setStakeReferralReader", [stake]);
  m.call(piece, "setNFT", [nft]);
  m.call(stake, "setNftContract", [nft]);
  m.call(stake, "setPieceToken", [piece]);

  return {
    usdt,
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
