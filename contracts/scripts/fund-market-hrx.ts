/**
 * 向 Market 转入 HRX，使 buyWithUsdt 可执行
 * 用法: npx hardhat run scripts/fund-market-hrx.ts --network bsc
 */
import { createWalletClient, createPublicClient, http, parseAbi } from "viem";
import { privateKeyToAccount } from "viem/accounts";
import { bsc } from "viem/chains";

const ADDRESSES = {
  Market: "0xE9E4C6c8Eb8b80420D94475C7278c693dd70b41a" as const,
  Hrx: "0xE3198bf629C91bd678190D22393Ba5DEc40E9106" as const,
};

const FUND_AMOUNT = 1_000_000n * 10n ** 18n; // 100 万 HRX

const erc20Abi = parseAbi([
  "function balanceOf(address) view returns (uint256)",
  "function transfer(address, uint256) returns (bool)",
]);

async function main() {
  const pk = process.env.BSC_PRIVATE_KEY?.trim();
  if (!pk || pk.length < 10) {
    console.error("请设置 BSC_PRIVATE_KEY");
    process.exit(1);
  }

  const account = privateKeyToAccount(("0x" + pk.replace(/^0x/, "")) as `0x${string}`);
  const rpc = process.env.BSC_RPC_URL || "https://bsc-rpc.publicnode.com";
  const transport = http(rpc, { timeout: 120_000 });
  const walletClient = createWalletClient({ account, chain: bsc, transport });
  const publicClient = createPublicClient({ chain: bsc, transport });

  const balance = await publicClient.readContract({
    address: ADDRESSES.Hrx,
    abi: erc20Abi,
    functionName: "balanceOf",
    args: [account.address],
  });

  if (balance < FUND_AMOUNT) {
    console.error("HRX 不足，需要:", FUND_AMOUNT.toString());
    process.exit(1);
  }

  console.log("向 Market 转入", FUND_AMOUNT / 10n ** 18n, "HRX...");
  const tx = await walletClient.writeContract({
    address: ADDRESSES.Hrx,
    abi: erc20Abi,
    functionName: "transfer",
    args: [ADDRESSES.Market, FUND_AMOUNT],
  });
  console.log("tx:", tx);
  await publicClient.waitForTransactionReceipt({ hash: tx });
  console.log("✅ 完成");
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});
