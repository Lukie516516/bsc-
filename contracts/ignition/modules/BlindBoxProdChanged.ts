/**
 * BSC 主网正式版：只重部署改动过的合约
 *
 * 复用旧正式版地址：
 * - Hrx
 * - Piece
 * - BondingCurvePrimaryMarket
 * - SwapHrxUsdt
 * - DealRouter
 * - BSC USDT
 *
 * 重部署：
 * - UmbrellaStakeDividend
 * - BlindBoxNFT
 * - BlindBoxUserVault
 * - BlindBoxVaultFactory
 * - BlindBoxPrincipalVault
 *
 * 运行：
 * npx hardhat ignition deploy ./ignition/modules/BlindBoxProdChanged.ts --network bsc
 */

import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const USDT = "0x55d398326f99059fF775485246999027B3197955";

// 旧正式版保留地址
const OLD_HRX = "0x057B67e0645D4db6A7373C7EA6fA6eeD4f218fFD";
const OLD_PIECE = "0xa18144b788732760b037fff964E29E245C84692a";
const OLD_MARKET = "0xdA5A5BC6b61B2eAB091A1b7E4d84C7eb537BDC0e";
const OLD_SWAP = "0x3c9Ea8d476eA06108283457B3Bb0dfbfD6cf1A1b";
const OLD_ROUTER = "0xeB9ED1fefF998f4B56a112713D149564616d7690";

// 旧正式版部署记录里 accountA/B/C/D 都是同一个地址
const DEFAULT_ACCOUNT = "0x6807Ed8c44AE2a89DDB7706fA22A5b274EcD8785";
const DEFAULT_MIN_STAKE_A = 300n * 10n ** 18n;

export default buildModule("BlindBoxProdChanged", (m) => {
  const accountA = m.getParameter("accountA", DEFAULT_ACCOUNT);
  const accountB = m.getParameter("accountB", DEFAULT_ACCOUNT);
  const accountC = m.getParameter("accountC", DEFAULT_ACCOUNT);
  const accountD = m.getParameter("accountD", DEFAULT_ACCOUNT);
  const minStakeA = m.getParameter("minStakeA", DEFAULT_MIN_STAKE_A);

  const hrx = m.contractAt("Hrx", OLD_HRX);
  const piece = m.contractAt("Piece", OLD_PIECE);
  const market = m.contractAt("BondingCurvePrimaryMarket", OLD_MARKET);
  const swap = m.contractAt("SwapHrxUsdt", OLD_SWAP);
  const router = m.contractAt("DealRouter", OLD_ROUTER);

  const nftPriceLib = m.library("NftPriceLib");
  const nftFeeLib = m.library("NftFeeLib");

  const stake = m.contract("UmbrellaStakeDividend", [
    hrx,
    USDT,
    accountA,
    minStakeA,
  ]);

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

  const listTemplate = m.contract("BlindBoxUserVault");

  const factoryList = m.contract("BlindBoxVaultFactory", [
    listTemplate,
    nft,
    USDT,
  ]);

  const stakeTwo = m.contract("BlindBoxPrincipalVault", [nft]);

  m.call(nft, "setFeeQuotePool", [swap]);
  m.call(nft, "setPieceToken", [piece]);
  m.call(nft, "setPieceValueUsdtE18", [1n * 10n ** 18n]);
  m.call(nft, "setStakeReferralReader", [stake]);

  m.call(piece, "setNFT", [nft]);
  m.call(stake, "setNftContract", [nft]);
  m.call(stake, "setPieceToken", [piece]);

  return {
    hrx,
    piece,
    market,
    swap,
    router,
    nftPriceLib,
    nftFeeLib,
    stake,
    nft,
    listTemplate,
    factoryList,
    stakeTwo,
  };
});
