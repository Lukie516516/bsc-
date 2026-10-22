/**
 * 调用 MockUSDT.mint() 给指定地址打测试 U
 * 用法:
 *   TARGET_ADDRESS=0x... BSC_PRIVATE_KEY=0x... node scripts/mint-usdt-simple.mjs
 * 可选: AMOUNT_USDT=10000（默认 10000，18 位小数）
 */
import { createWalletClient, createPublicClient, http, parseAbi } from "viem";
import { privateKeyToAccount } from "viem/accounts";
import { bsc } from "viem/chains";

const MOCK_USDT = "0x2674DE1Fe0c884D45cd128b3F67Eb03BFD649ff4";
const TARGET =
  process.env.TARGET_ADDRESS?.trim() || "0xE2D4ecEcb19B39D7e796fF9A4965543AeD903069";
const amountUsdt = BigInt(process.env.AMOUNT_USDT || "10000");
const AMOUNT = amountUsdt * 10n ** 18n;
const PK = process.env.BSC_PRIVATE_KEY;
if (!PK) {
  throw new Error("BSC_PRIVATE_KEY is required");
}

const account = privateKeyToAccount(("0x" + PK.replace(/^0x/, "")));

const rpc = process.env.BSC_RPC_URL || "https://bsc-rpc.publicnode.com";
const transport = http(rpc, { timeout: 120_000 });

const client = createWalletClient({
  account,
  chain: bsc,
  transport,
});

const publicClient = createPublicClient({
  chain: bsc,
  transport,
});

const abi = parseAbi(["function mint(address to, uint256 amount) external"]);

async function main() {
  const hash = await client.writeContract({
    address: MOCK_USDT,
    abi,
    functionName: "mint",
    args: [TARGET, AMOUNT],
  });
  console.log("Mint tx:", hash);
  await publicClient.waitForTransactionReceipt({ hash });
  console.log("已给", TARGET, "打", amountUsdt.toString(), "测试 U");
}

main().catch(console.error);
