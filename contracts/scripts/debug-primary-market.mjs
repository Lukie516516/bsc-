/**
 * 只读诊断 BondingCurvePrimaryMarket（内盘）初始化与余额
 * 用法: node scripts/debug-primary-market.mjs
 */
import { createPublicClient, http, parseAbi } from "viem";
import { bsc } from "viem/chains";

const MARKET = "0xE9E4C6c8Eb8b80420D94475C7278c693dd70b41a";
const USDT = "0x2674DE1Fe0c884D45cd128b3F67Eb03BFD649ff4";
const HRX = "0xE3198bf629C91bd678190D22393Ba5DEc40E9106";

const rpc = process.env.BSC_RPC_URL || "https://bsc-rpc.publicnode.com";
const publicClient = createPublicClient({
  chain: bsc,
  transport: http(rpc, { timeout: 60_000 }),
});

const erc20 = parseAbi(["function balanceOf(address) view returns (uint256)"]);
const market = parseAbi([
  "function initStatus() view returns (bool)",
  "function x1() view returns (uint256)",
  "function x2() view returns (uint256)",
  "function s1() view returns (uint256)",
  "function s2() view returns (uint256)",
  "function modeledReserve() view returns (uint256)",
  "function internalSupply() view returns (uint256)",
]);

async function main() {
  const initStatus = await publicClient.readContract({
    address: MARKET,
    abi: market,
    functionName: "initStatus",
  });
  const [x1, x2, s1, s2, modeled, internal] = await Promise.all([
    publicClient.readContract({ address: MARKET, abi: market, functionName: "x1" }),
    publicClient.readContract({ address: MARKET, abi: market, functionName: "x2" }),
    publicClient.readContract({ address: MARKET, abi: market, functionName: "s1" }),
    publicClient.readContract({ address: MARKET, abi: market, functionName: "s2" }),
    publicClient.readContract({
      address: MARKET,
      abi: market,
      functionName: "modeledReserve",
    }),
    publicClient.readContract({
      address: MARKET,
      abi: market,
      functionName: "internalSupply",
    }),
  ]);
  const usdtBal = await publicClient.readContract({
    address: USDT,
    abi: erc20,
    functionName: "balanceOf",
    args: [MARKET],
  });
  const hrxBal = await publicClient.readContract({
    address: HRX,
    abi: erc20,
    functionName: "balanceOf",
    args: [MARKET],
  });

  console.log(JSON.stringify({
    market: MARKET,
    initStatus,
    ledger: { x1: x1.toString(), x2: x2.toString(), s1: s1.toString(), s2: s2.toString() },
    modeledReserve: modeled.toString(),
    internalSupply: internal.toString(),
    marketUsdtBalance: usdtBal.toString(),
    marketHrxBalance: hrxBal.toString(),
  }, null, 2));

  if (!initStatus) {
    console.log("\n结论: initStatus=false，尚未执行 buyWithUsdtInit，需要跑 init-primary-market。");
  } else if (usdtBal === 0n && internal > 0n) {
    console.log("\n提示: 已初始化但 Market 上 MockUSDT 余额为 0；用户「卖 HRX 换 U」可能因合约里没有可转出的 U 而 revert。");
  }
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});
