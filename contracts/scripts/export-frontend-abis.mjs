/**
 * 从 Hardhat artifacts 导出「前端用」多合约 ABI 汇总。
 * 用法（在 box 目录）: node scripts/export-frontend-abis.mjs
 */
import fs from "fs";
import path from "path";
import { fileURLToPath } from "url";

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const root = path.resolve(__dirname, "..");
const artifactsDir = path.join(root, "artifacts", "contracts");
const outDir = path.join(root, "..", "requestEth", "abi");
const outFile = path.join(outDir, "frontend_all_contracts.json");

/** 合约逻辑名 -> artifacts 内相对路径（相对于 artifacts/contracts） */
const ARTIFACTS = [
  ["BlindBoxNFT", "Nft.sol/BlindBoxNFT.json"],
  ["BondingCurvePrimaryMarket", "Market.sol/BondingCurvePrimaryMarket.json"],
  ["SwapHrxUsdt", "SwapHrxUsdt.sol/SwapHrxUsdt.json"],
  ["DealRouter", "Router.sol/DealRouter.json"],
  ["UmbrellaStakeDividend", "Stake.sol/UmbrellaStakeDividend.json"],
  ["BlindBoxVaultFactory", "FactoryList.sol/BlindBoxVaultFactory.json"],
  ["BlindBoxPrincipalVault", "StakeTwo.sol/BlindBoxPrincipalVault.json"],
  ["BlindBoxUserVault", "ListTemplate.sol/BlindBoxUserVault.json"],
  ["MockUSDT", "MockUSDT.sol/MockUSDT.json"],
  ["Hrx", "Hrx.sol/Hrx.json"],
  ["Piece", "Piece.sol/Piece.json"],
  ["NftPriceLib", "NftPriceLib.sol/NftPriceLib.json"],
  ["NftFeeLib", "NftFeeLib.sol/NftFeeLib.json"],
  ["IERC20", "interfaces/IERC20.sol/IERC20.json"],
  ["IDealPair", "interfaces/IDealPair.sol/IDealPair.json"],
];

function loadAbi(relPath) {
  const full = path.join(artifactsDir, relPath);
  if (!fs.existsSync(full)) {
    throw new Error(`Missing artifact: ${full}`);
  }
  const j = JSON.parse(fs.readFileSync(full, "utf8"));
  if (!Array.isArray(j.abi)) {
    throw new Error(`No abi array in ${full}`);
  }
  return j.abi;
}

const contracts = {};
for (const [name, rel] of ARTIFACTS) {
  contracts[name] = loadAbi(rel);
}

const payload = {
  _meta: {
    description:
      "BlindBox 项目前端用 ABI 汇总：每个 key 为合约逻辑名，value 为该合约完整 abi 数组。请先执行 npx hardhat compile 再生成。",
    generatedAt: new Date().toISOString(),
    source: "box/artifacts/contracts (Hardhat)",
  },
  contracts,
};

fs.mkdirSync(outDir, { recursive: true });
fs.writeFileSync(outFile, JSON.stringify(payload, null, 2), "utf8");
console.log("Wrote", outFile);
console.log("Keys:", Object.keys(contracts).join(", "));
