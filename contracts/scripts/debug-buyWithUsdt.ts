/**
 * buyWithUsdt 调试脚本：检查池子状态、用户授权、报价等
 *
 * 用法: npx hardhat run scripts/debug-buyWithUsdt.ts --network bsc
 *
 * 环境变量（可选）:
 *   USER_ADDRESS - 要调试的用户地址，默认 0xE2D4ecEcb19B39D7e796fF9A4965543AeD903069
 *   USDT_AMOUNT  - 模拟购买的 USDT 数量（18位），默认 1000
 */

import { createPublicClient, http, parseAbi } from "viem";
import { bsc } from "viem/chains";

// BlindBoxMainnetTest 部署地址
const ADDRESSES = {
  Market: "0xE9E4C6c8Eb8b80420D94475C7278c693dd70b41a" as const,
  MockUSDT: "0x2674DE1Fe0c884D45cd128b3F67Eb03BFD649ff4" as const,
  Hrx: "0xE3198bf629C91bd678190D22393Ba5DEc40E9106" as const,
};

const USER = (process.env.USER_ADDRESS || "0xE2D4ecEcb19B39D7e796fF9A4965543AeD903069") as `0x${string}`;
const USDT_AMOUNT = BigInt(process.env.USDT_AMOUNT || "1000") * 10n ** 18n;

// 最小 ABI 用于读取
const erc20Abi = parseAbi([
  "function balanceOf(address) view returns (uint256)",
  "function allowance(address, address) view returns (uint256)",
]);
const marketAbi = parseAbi([
  "function initStatus() view returns (bool)",
  "function quoteBuyGivenUsdt(uint256) view returns (uint256, uint256, uint256)",
  "function x1() view returns (uint256)",
  "function x2() view returns (uint256)",
  "function s1() view returns (uint256)",
  "function s2() view returns (uint256)",
  "function feeRecipient() view returns (address)",
  "function buyWithUsdt(uint256, uint256, address) returns (uint256)",
]);

async function main() {
  const rpc = process.env.BSC_RPC_URL || "https://bsc.publicnode.com";
  const publicClient = createPublicClient({
    chain: bsc,
    transport: http(rpc, { timeout: 15_000 }),
  });

  console.log("========== buyWithUsdt 调试 ==========\n");
  console.log("合约地址:");
  console.log("  Market:", ADDRESSES.Market);
  console.log("  USDT:  ", ADDRESSES.MockUSDT);
  console.log("  HRX:   ", ADDRESSES.Hrx);
  console.log("用户地址:", USER);
  console.log("模拟金额:", USDT_AMOUNT.toString(), "wei (1000 USDT)\n");

  // 1. 池子是否已初始化
  const initStatus = await publicClient.readContract({
    address: ADDRESSES.Market,
    abi: marketAbi,
    functionName: "initStatus",
  });
  console.log("1. 池子初始化状态 initStatus:", initStatus);
  if (!initStatus) {
    console.log("   ⚠️ 池子未初始化！必须先调用 buyWithUsdtInit");
    console.log("   解决: 运行 npx hardhat run scripts/init-primary-market.ts --network bsc\n");
  }

  // 2. Market 的 HRX 余额
  const marketHrxBalance = await publicClient.readContract({
    address: ADDRESSES.Hrx,
    abi: erc20Abi,
    functionName: "balanceOf",
    args: [ADDRESSES.Market],
  });
  console.log("2. Market 合约 HRX 余额:", formatE18(marketHrxBalance));

  // 3. Market 的 USDT 余额
  const marketUsdtBalance = await publicClient.readContract({
    address: ADDRESSES.MockUSDT,
    abi: erc20Abi,
    functionName: "balanceOf",
    args: [ADDRESSES.Market],
  });
  console.log("3. Market 合约 USDT 余额:", formatE18(marketUsdtBalance));

  // 4. 用户 USDT 余额
  const userUsdtBalance = await publicClient.readContract({
    address: ADDRESSES.MockUSDT,
    abi: erc20Abi,
    functionName: "balanceOf",
    args: [USER],
  });
  console.log("4. 用户 USDT 余额:", formatE18(userUsdtBalance));

  // 5. 用户对 Market 的 USDT 授权
  const allowance = await publicClient.readContract({
    address: ADDRESSES.MockUSDT,
    abi: erc20Abi,
    functionName: "allowance",
    args: [USER, ADDRESSES.Market],
  });
  console.log("5. 用户 USDT 授权 (allowance):", formatE18(allowance));
  if (allowance < USDT_AMOUNT) {
    console.log("   ⚠️ 授权不足！用户需先 approve(Market, 金额)");
  }

  // 6. 报价
  try {
    const [ausdNetOut, ausdFee, ausdGrossOut] = await publicClient.readContract({
      address: ADDRESSES.Market,
      abi: marketAbi,
      functionName: "quoteBuyGivenUsdt",
      args: [USDT_AMOUNT],
    });
    console.log("6. 报价 quoteBuyGivenUsdt(1000 USDT):");
    console.log("   预计获得 HRX:", formatE18(ausdNetOut));
    console.log("   手续费 HRX: ", formatE18(ausdFee));
    console.log("   总铸造 HRX: ", formatE18(ausdGrossOut));
    if (marketHrxBalance < ausdGrossOut) {
      console.log("   ⚠️ Market HRX 不足！需要", formatE18(ausdGrossOut), "实际只有", formatE18(marketHrxBalance));
    }
  } catch (e: unknown) {
    console.log("6. 报价失败:", (e as Error).message);
  }

  // 7. 池子状态
  const [x1, x2, s1, s2, feeRecipient] = await Promise.all([
    publicClient.readContract({ address: ADDRESSES.Market, abi: marketAbi, functionName: "x1" }),
    publicClient.readContract({ address: ADDRESSES.Market, abi: marketAbi, functionName: "x2" }),
    publicClient.readContract({ address: ADDRESSES.Market, abi: marketAbi, functionName: "s1" }),
    publicClient.readContract({ address: ADDRESSES.Market, abi: marketAbi, functionName: "s2" }),
    publicClient.readContract({ address: ADDRESSES.Market, abi: marketAbi, functionName: "feeRecipient" }),
  ]);
  console.log("\n7. 池子状态:");
  console.log("   x1 (买轴):", x1.toString());
  console.log("   x2 (卖轴):", x2.toString());
  console.log("   s1 (入金):", s1.toString());
  console.log("   s2 (出金):", s2.toString());
  console.log("   feeRecipient:", feeRecipient);

  // 8. 模拟调用
  console.log("\n8. 模拟 buyWithUsdt 调用...");
  const [quoteNet] = await publicClient.readContract({
    address: ADDRESSES.Market,
    abi: marketAbi,
    functionName: "quoteBuyGivenUsdt",
    args: [USDT_AMOUNT],
  });
  const minOut = quoteNet * 95n / 100n; // 5% 滑点
  try {
    await publicClient.simulateContract({
      address: ADDRESSES.Market,
      abi: marketAbi,
      functionName: "buyWithUsdt",
      args: [USDT_AMOUNT, minOut, USER],
      account: USER,
    });
    console.log("   ✅ 模拟成功，交易应可执行");
  } catch (e: unknown) {
    const err = e as Error & { cause?: { data?: string } };
    console.log("   ❌ 模拟失败:", err.message);
    if (err.cause && typeof err.cause === "object" && "data" in err.cause) {
      console.log("   错误 data:", (err.cause as { data: string }).data);
    }
  }

  console.log("\n========== 调试结束 ==========");
}

function formatE18(n: bigint): string {
  const s = n.toString();
  if (s.length <= 18) return "0." + s.padStart(18, "0");
  return s.slice(0, -18) + "." + s.slice(-18).replace(/0+$/, "") || "0";
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});
