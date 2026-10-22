/**
 * 测试槽位购买报价：previewSlotCost(USDT) 和 quoteSlotCostInB(HRX)
 *
 * 用法: npx hardhat run scripts/test-slot-cost.ts --network bsc
 *
 * 环境变量（可选）:
 *   USER_ADDRESS - 要查询的用户地址，默认 0xE2D4ecEcb19B39D7e796fF9A4965543AeD903069
 *   SLOT_COUNT   - 购买槽位数量，默认 1
 */

import { createPublicClient, http, parseAbi } from "viem";
import { bsc } from "viem/chains";

const NFT_ADDRESS = "0x73F76b11ebF03718B71c163BE2B5445ECFC25Ce6" as const;
const USER = (process.env.USER_ADDRESS || "0xE2D4ecEcb19B39D7e796fF9A4965543AeD903069") as `0x${string}`;
const SLOT_COUNT = BigInt(process.env.SLOT_COUNT || "1");

const nftAbi = parseAbi([
  "function previewSlotCost(address user, uint256 count) view returns (uint256 cost, uint256 step)",
  "function quoteSlotCostInB(address user, uint256 count) view returns (uint256 costB)",
  "function extraStorageSlots(address) view returns (uint256)",
  "function slotFirstCost() view returns (uint256)",
  "function slotStepCost() view returns (uint256)",
  "function feeQuotePool() view returns (address)",
  "function feeTokenB() view returns (address)",
]);

const poolAbi = parseAbi([
  "function token0() view returns (address)",
  "function token1() view returns (address)",
  "function getReserves() view returns (uint112 reserve0, uint112 reserve1, uint32 ts)",
  "function liquidityInited() view returns (bool)",
  "function feeNum() view returns (uint32)",
  "function feeDen() view returns (uint32)",
]);

function formatE18(n: bigint): string {
  const s = n.toString();
  if (s.length <= 18) return "0." + s.padStart(18, "0");
  return s.slice(0, -18) + "." + s.slice(-18).replace(/0+$/, "") || "0";
}

async function main() {
  const rpc = process.env.BSC_RPC_URL || "https://bsc.publicnode.com";
  const client = createPublicClient({
    chain: bsc,
    transport: http(rpc, { timeout: 15_000 }),
  });

  console.log("========== 槽位购买报价测试 ==========\n");
  console.log("BlindBoxNFT:", NFT_ADDRESS);
  console.log("用户地址:  ", USER);
  console.log("购买数量:  ", SLOT_COUNT.toString(), "个槽位\n");

  // 1. 用户当前槽位数
  const extraSlots = await client.readContract({
    address: NFT_ADDRESS,
    abi: nftAbi,
    functionName: "extraStorageSlots",
    args: [USER],
  });
  console.log("1. 用户已购买槽位 extraStorageSlots:", extraSlots.toString());

  // 2. 槽位计价参数
  const [slotFirstCost, slotStepCost] = await Promise.all([
    client.readContract({ address: NFT_ADDRESS, abi: nftAbi, functionName: "slotFirstCost" }),
    client.readContract({ address: NFT_ADDRESS, abi: nftAbi, functionName: "slotStepCost" }),
  ]);
  console.log("2. 槽位计价 slotFirstCost:", formatE18(slotFirstCost), "USDT");
  console.log("   slotStepCost:        ", formatE18(slotStepCost), "USDT/档");

  // 3. previewSlotCost - 购买所需 USDT
  const [costUSDT] = await client.readContract({
    address: NFT_ADDRESS,
    abi: nftAbi,
    functionName: "previewSlotCost",
    args: [USER, SLOT_COUNT],
  });
  console.log("\n3. previewSlotCost(user, count) 购买所需 USDT:");
  console.log("   cost:", formatE18(costUSDT), "USDT");
  console.log("   raw: ", costUSDT.toString(), "(18位小数)");

  // 4. quoteSlotCostInB - 购买所需 HRX
  const costHRX = await client.readContract({
    address: NFT_ADDRESS,
    abi: nftAbi,
    functionName: "quoteSlotCostInB",
    args: [USER, SLOT_COUNT],
  });
  console.log("\n4. quoteSlotCostInB(user, count) 购买所需 HRX:");
  console.log("   costB:", formatE18(costHRX), "HRX");
  console.log("   raw:  ", costHRX.toString(), "(18位小数)");

  // 5. 检查 feeQuotePool 池子状态
  const poolAddr = await client.readContract({
    address: NFT_ADDRESS,
    abi: nftAbi,
    functionName: "feeQuotePool",
  });
  console.log("\n5. feeQuotePool 池子地址:", poolAddr);

  if (poolAddr && poolAddr !== "0x0000000000000000000000000000000000000000") {
    const [token0, token1, liquidityInited] = await Promise.all([
      client.readContract({ address: poolAddr as `0x${string}`, abi: poolAbi, functionName: "token0" }),
      client.readContract({ address: poolAddr as `0x${string}`, abi: poolAbi, functionName: "token1" }),
      client.readContract({ address: poolAddr as `0x${string}`, abi: poolAbi, functionName: "liquidityInited" }),
    ]);
    console.log("   token0 (应为 HRX):", token0);
    console.log("   token1 (应为 USDT):", token1);
    console.log("   liquidityInited:", liquidityInited);

    const [r0, r1] = await client.readContract({
      address: poolAddr as `0x${string}`,
      abi: poolAbi,
      functionName: "getReserves",
    });
    console.log("   reserve0 (HRX):", formatE18(r0), "| raw:", r0.toString());
    console.log("   reserve1 (USDT):", formatE18(r1), "| raw:", r1.toString());

    const [feeNum, feeDen] = await Promise.all([
      client.readContract({ address: poolAddr as `0x${string}`, abi: poolAbi, functionName: "feeNum" }),
      client.readContract({ address: poolAddr as `0x${string}`, abi: poolAbi, functionName: "feeDen" }),
    ]);
    console.log("   feeNum/feeDen:", feeNum.toString(), "/", feeDen.toString());

    // 理论计算: grossMax = r0 * amountIn / (r1 + amountIn)
    const amountIn = costUSDT;
    const grossMax = (r0 * amountIn) / (r1 + amountIn);
    const feeOut = (grossMax * BigInt(feeNum)) / BigInt(feeDen);
    const expectedOut = grossMax - feeOut;
    console.log("\n   理论 grossMax:", formatE18(grossMax));
    console.log("   理论 out(扣费后):", formatE18(expectedOut));
  }

  console.log("\n========== 测试结束 ==========");
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});
