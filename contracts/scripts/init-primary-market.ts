/**
 * 初始化 BondingCurvePrimaryMarket 池子（buyWithUsdtInit）
 *
 * 池子部署后必须执行一次 init，否则 buyWithUsdt 会因 Market 无 HRX 而失败。
 *
 * 用法: npx hardhat run scripts/init-primary-market.ts --network bsc
 *
 * 环境变量:
 *   BSC_PRIVATE_KEY - deployer 私钥（必填）
 *   BSC_RPC_URL     - RPC 地址，默认 bsc.publicnode.com
 *   INIT_USDT       - 初始化投入的 USDT 数量（18位），默认 100
 *   INIT_TO         - 接收 HRX 的地址，默认 deployer
 */

import { createWalletClient, createPublicClient, http, parseAbi } from "viem";
import { privateKeyToAccount } from "viem/accounts";
import { bsc } from "viem/chains";

const ADDRESSES = {
  Market: "0xE9E4C6c8Eb8b80420D94475C7278c693dd70b41a" as const,
  MockUSDT: "0x2674DE1Fe0c884D45cd128b3F67Eb03BFD649ff4" as const,
  Hrx: "0xE3198bf629C91bd678190D22393Ba5DEc40E9106" as const,
};

const INIT_USDT = BigInt(process.env.INIT_USDT || "100") * 10n ** 18n;
const INIT_TO = (process.env.INIT_TO || "") as `0x${string}`;

const erc20Abi = parseAbi([
  "function balanceOf(address) view returns (uint256)",
  "function transfer(address, uint256) returns (bool)",
  "function approve(address, uint256) returns (bool)",
]);
const marketAbi = parseAbi([
  "function initStatus() view returns (bool)",
  "function quoteBuyGivenUsdt(uint256) view returns (uint256, uint256, uint256)",
  "function buyWithUsdtInit(uint256, address) returns (uint256)",
]);

async function main() {
  const pk = process.env.BSC_PRIVATE_KEY?.trim();
  if (!pk || pk.length < 10) {
    console.error("请设置 BSC_PRIVATE_KEY 环境变量");
    process.exit(1);
  }

  const account = privateKeyToAccount(("0x" + pk.replace(/^0x/, "")) as `0x${string}`);
  const deployer = account.address;
  const to = INIT_TO || deployer;

  const rpc = process.env.BSC_RPC_URL || "https://bsc-rpc.publicnode.com";
  const transport = http(rpc, { timeout: 120_000 });
  const walletClient = createWalletClient({
    account,
    chain: bsc,
    transport,
  });
  const publicClient = createPublicClient({ chain: bsc, transport });

  const initStatus = await publicClient.readContract({
    address: ADDRESSES.Market,
    abi: marketAbi,
    functionName: "initStatus",
  });
  if (initStatus) {
    console.log("池子已初始化，无需重复执行");
    process.exit(0);
  }

  const [, , ausdGrossOut] = await publicClient.readContract({
    address: ADDRESSES.Market,
    abi: marketAbi,
    functionName: "quoteBuyGivenUsdt",
    args: [INIT_USDT],
  });
  const hrxNeeded = ausdGrossOut;

  console.log("初始化 BondingCurvePrimaryMarket");
  console.log("  Deployer:", deployer);
  console.log("  USDT 投入:", INIT_USDT.toString());
  console.log("  HRX 需转给 Market:", hrxNeeded.toString());
  console.log("  接收 HRX 地址:", to);

  const deployerHrx = await publicClient.readContract({
    address: ADDRESSES.Hrx,
    abi: erc20Abi,
    functionName: "balanceOf",
    args: [deployer],
  });
  const deployerUsdt = await publicClient.readContract({
    address: ADDRESSES.MockUSDT,
    abi: erc20Abi,
    functionName: "balanceOf",
    args: [deployer],
  });

  if (deployerHrx < hrxNeeded) {
    console.error("Deployer HRX 不足，需要:", hrxNeeded.toString());
    process.exit(1);
  }
  if (deployerUsdt < INIT_USDT) {
    console.error("Deployer USDT 不足，需要:", INIT_USDT.toString());
    process.exit(1);
  }

  // 1. 转 HRX 给 Market
  console.log("\n1. 转移 HRX 到 Market...");
  const tx1 = await walletClient.writeContract({
    address: ADDRESSES.Hrx,
    abi: erc20Abi,
    functionName: "transfer",
    args: [ADDRESSES.Market, hrxNeeded],
  });
  console.log("   tx:", tx1);
  await publicClient.waitForTransactionReceipt({ hash: tx1 });

  // 2. Approve USDT
  console.log("\n2. Approve USDT...");
  const tx2 = await walletClient.writeContract({
    address: ADDRESSES.MockUSDT,
    abi: erc20Abi,
    functionName: "approve",
    args: [ADDRESSES.Market, INIT_USDT],
  });
  console.log("   tx:", tx2);
  await publicClient.waitForTransactionReceipt({ hash: tx2 });

  // 3. buyWithUsdtInit
  console.log("\n3. 调用 buyWithUsdtInit...");
  const tx3 = await walletClient.writeContract({
    address: ADDRESSES.Market,
    abi: marketAbi,
    functionName: "buyWithUsdtInit",
    args: [INIT_USDT, to],
  });
  console.log("   tx:", tx3);
  await publicClient.waitForTransactionReceipt({ hash: tx3 });

  console.log("\n✅ 池子初始化完成");
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});
