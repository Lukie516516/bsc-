/**
 * 给指定地址 mint 测试 USDT（需在 BSC 测试网部署后使用）
 *
 * 用法: npx hardhat run scripts/mint-test-usdt.ts --network bscTestnet
 *
 * 环境变量: MOCK_USDT_ADDR, MINT_TO, MINT_AMOUNT (可选)
 * 或修改下方常量
 */

import hre from "hardhat";

const MOCK_USDT = process.env.MOCK_USDT_ADDR || "0x"; // 部署后的 MockUSDT 地址
const MINT_TO = process.env.MINT_TO || "0xE2D4ecEcb19B39D7e796fF9A4965543AeD903069";
const MINT_AMOUNT = process.env.MINT_AMOUNT || "10000000000000000000000"; // 10000 * 10^18

async function main() {
  if (MOCK_USDT === "0x" || MOCK_USDT.length < 10) {
    console.error("请设置 MOCK_USDT_ADDR 为部署后的 MockUSDT 合约地址");
    process.exit(1);
  }

  const [signer] = await hre.viem.getWalletClients();
  if (!signer) {
    console.error("未找到钱包，请配置 BSC_PRIVATE_KEY");
    process.exit(1);
  }

  const usdt = await hre.viem.getContractAt(
    "MockUSDT",
    MOCK_USDT as `0x${string}`,
    { client: { wallet: signer } }
  );

  const tx = await usdt.write.mint([
    MINT_TO as `0x${string}`,
    BigInt(MINT_AMOUNT),
  ]);
  console.log("Mint tx:", tx);
  console.log(`已给 ${MINT_TO} 打 ${MINT_AMOUNT} 个测试 U（18位小数）`);
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});
