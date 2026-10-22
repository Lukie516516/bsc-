/**
 * BSC 主网盲盒合约部署（使用 MockUSDT 测试）
 *
 * 与主网正式版区别：用 MockUSDT 替代真实 USDT，可 mint 测试用 U
 * 部署后需给测试地址 mint：usdt.mint(地址, 数量)
 *
 * 部署前：确保 deployer 有足够 BNB 支付 gas
 *
 * 运行：npx hardhat ignition deploy ./ignition/modules/BlindBoxMainnetTest.ts --network bsc
 */

import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const WBNB = "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c";

const PIECE_SUPPLY = 1171350625n * 10n ** 18n;
const MOCK_USDT_INITIAL = 10_000_000n * 10n ** 18n; // 1000 万测试 U
const DEFAULT_MIN_STAKE_A = 300n * 10n ** 18n;

export default buildModule("BlindBoxMainnetTest", (m) => {
  const deployer = m.getAccount(0);

  const initialRecipient = m.getParameter("initialRecipient", deployer);
  const accountA = m.getParameter("accountA", deployer);
  const accountB = m.getParameter("accountB", deployer);
  const accountC = m.getParameter("accountC", deployer);
  const accountD = m.getParameter("accountD", deployer);
  const minStakeA = m.getParameter("minStakeA", DEFAULT_MIN_STAKE_A);

  const nftPriceLib = m.library("NftPriceLib");
  const nftFeeLib = m.library("NftFeeLib");

  // MockUSDT（主网测试用）
  const usdt = m.contract("MockUSDT", [deployer, MOCK_USDT_INITIAL]);

  const piece = m.contract("Piece", [
    "Piece",
    "PIECE",
    initialRecipient,
    PIECE_SUPPLY,
  ]);

  const hrx = m.contract("Hrx", ["Hrx", "HRX"]);

  const market = m.contract("BondingCurvePrimaryMarket", [
    usdt,
    hrx,
    deployer,
    accountA,
  ]);

  const stake = m.contract("UmbrellaStakeDividend", [
    hrx,
    usdt,
    accountA,
    minStakeA,
  ]);

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

  const swap = m.contract("SwapHrxUsdt", [
    hrx,
    usdt,
    deployer,
    accountA,
    3,
    100,
  ]);

  const router = m.contract("DealRouter", [WBNB]);

  const listTemplate = m.contract("BlindBoxUserVault");

  const factoryList = m.contract("BlindBoxVaultFactory", [
    listTemplate,
    nft,
    usdt,
  ]);

  const stakeTwo = m.contract("BlindBoxPrincipalVault", [nft]);

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
