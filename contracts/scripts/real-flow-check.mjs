import { createPublicClient, createWalletClient, decodeEventLog, getAddress, http, maxUint256, parseAbi, parseUnits } from "viem";
import { privateKeyToAccount } from "viem/accounts";
import { bsc } from "viem/chains";

const rpc = process.env.BSC_RPC_URL || "https://bnb56743.allnodes.me:8545/or6NBNSPsU9uDxP3";
const pk = (process.env.TEST_PRIVATE_KEY || "").trim().replace(/^0x/i, "");
if (!pk) {
  throw new Error("TEST_PRIVATE_KEY is required");
}

const account = privateKeyToAccount(`0x${pk}`);
const publicClient = createPublicClient({ chain: bsc, transport: http(rpc) });
const walletClient = createWalletClient({ account, chain: bsc, transport: http(rpc) });
const targetUsdtRaw = process.env.TARGET_USDT || "30";

const ADDR = {
  hrx: getAddress("0x057B67e0645D4db6A7373C7EA6fA6eeD4f218fFD"),
  usdt: getAddress("0x55d398326f99059fF775485246999027B3197955"),
  market: getAddress("0xdA5A5BC6b61B2eAB091A1b7E4d84C7eb537BDC0e"),
  nft: getAddress("0x169c4E946909D4E115b724957e5D1c74d0971f70"),
};

const erc20Abi = parseAbi([
  "function balanceOf(address) view returns (uint256)",
  "function allowance(address,address) view returns (uint256)",
  "function approve(address,uint256) returns (bool)",
  "function transfer(address,uint256) returns (bool)",
]);

const marketAbi = parseAbi([
  "function quoteSellForExactUsdt(uint256 usdtOut) view returns (uint256 ausdGrossIn, uint256 ausdFee, uint256 ausdBurn)",
  "function sellExactUsdt(uint256 usdtOut, uint256 maxAusdGrossIn, address to) returns (uint256 ausdGrossIn)",
]);

const nftAbi = parseAbi([
  "function mint(uint8 tier) returns (uint256 tokenId)",
  "function quoteFeeInB(uint256 amount1InUSDT) view returns (uint256 out0NetB)",
  "function openBoxPayHRX(uint256 tokenId,uint256 maxFeeB)",
  "event Minted(address indexed minter,uint256 indexed tokenId,uint8 indexed tier,uint256 usdtPaid)",
  "event Opened(address indexed owner,uint256 indexed tokenId,uint256 timestamp,uint256 usdtFeePaid,uint256 feeBPaid)",
]);

async function ensureApproval(token, spender, minAmount) {
  const allowance = await publicClient.readContract({
    address: token,
    abi: erc20Abi,
    functionName: "allowance",
    args: [account.address, spender],
  });
  if (allowance >= minAmount) {
    return null;
  }
  const hash = await walletClient.writeContract({
    address: token,
    abi: erc20Abi,
    functionName: "approve",
    args: [spender, maxUint256],
    account,
  });
  await publicClient.waitForTransactionReceipt({ hash });
  return hash;
}

function printStep(name, data) {
  console.log(JSON.stringify({ step: name, ...data }, null, 2));
}

async function main() {
  printStep("account", { address: account.address, rpc });

  const before = {
    hrx: await publicClient.readContract({ address: ADDR.hrx, abi: erc20Abi, functionName: "balanceOf", args: [account.address] }),
    usdt: await publicClient.readContract({ address: ADDR.usdt, abi: erc20Abi, functionName: "balanceOf", args: [account.address] }),
  };
  printStep("balances_before", { hrx: before.hrx.toString(), usdt: before.usdt.toString() });

  const targetUsdt = parseUnits(targetUsdtRaw, 18);
  const [quotedHrxIn] = await publicClient.readContract({
    address: ADDR.market,
    abi: marketAbi,
    functionName: "quoteSellForExactUsdt",
    args: [targetUsdt],
  });
  const maxHrxIn = quotedHrxIn * 2n;
  await ensureApproval(ADDR.hrx, ADDR.market, maxHrxIn);
  const swapHash = await walletClient.writeContract({
    address: ADDR.market,
    abi: marketAbi,
    functionName: "sellExactUsdt",
    args: [targetUsdt, maxHrxIn, account.address],
    account,
  });
  const swapReceipt = await publicClient.waitForTransactionReceipt({ hash: swapHash });
  printStep("swap", { hash: swapHash, status: swapReceipt.status, targetUsdt: targetUsdt.toString(), quotedHrxIn: quotedHrxIn.toString(), maxHrxIn: maxHrxIn.toString() });

  const afterSwapUsdt = await publicClient.readContract({
    address: ADDR.usdt,
    abi: erc20Abi,
    functionName: "balanceOf",
    args: [account.address],
  });
  printStep("balances_after_swap", { usdt: afterSwapUsdt.toString() });

  const mintCost = parseUnits("20", 18);
  await ensureApproval(ADDR.usdt, ADDR.nft, mintCost);
  const mintHash = await walletClient.writeContract({
    address: ADDR.nft,
    abi: nftAbi,
    functionName: "mint",
    args: [1],
    account,
  });
  const mintReceipt = await publicClient.waitForTransactionReceipt({ hash: mintHash });
  let tokenId = null;
  for (const log of mintReceipt.logs) {
    try {
      const parsed = decodeEventLog({ abi: nftAbi, data: log.data, topics: log.topics });
      if (parsed.eventName === "Minted") {
        tokenId = parsed.args.tokenId;
      }
    } catch {}
  }
  if (tokenId == null) {
    throw new Error("Minted tokenId not found in receipt logs");
  }
  printStep("mint", { hash: mintHash, status: mintReceipt.status, tokenId: tokenId.toString() });

  const feeUsdt = parseUnits("0.2", 18);
  const feeB = await publicClient.readContract({
    address: ADDR.nft,
    abi: nftAbi,
    functionName: "quoteFeeInB",
    args: [feeUsdt],
  });
  const maxFeeB = feeB * 2n;
  await ensureApproval(ADDR.hrx, ADDR.nft, maxFeeB);
  const openHash = await walletClient.writeContract({
    address: ADDR.nft,
    abi: nftAbi,
    functionName: "openBoxPayHRX",
    args: [tokenId, maxFeeB],
    account,
  });
  const openReceipt = await publicClient.waitForTransactionReceipt({ hash: openHash });
  printStep("open", { hash: openHash, status: openReceipt.status, feeB: feeB.toString(), maxFeeB: maxFeeB.toString() });

  const fundAmount = parseUnits("5", 18);
  const fundHash = await walletClient.writeContract({
    address: ADDR.usdt,
    abi: erc20Abi,
    functionName: "transfer",
    args: [ADDR.nft, fundAmount],
    account,
  });
  const fundReceipt = await publicClient.waitForTransactionReceipt({ hash: fundHash });
  printStep("fund_nft", { hash: fundHash, status: fundReceipt.status, amount: fundAmount.toString() });

  const after = {
    hrx: await publicClient.readContract({ address: ADDR.hrx, abi: erc20Abi, functionName: "balanceOf", args: [account.address] }),
    usdt: await publicClient.readContract({ address: ADDR.usdt, abi: erc20Abi, functionName: "balanceOf", args: [account.address] }),
    nftUsdt: await publicClient.readContract({ address: ADDR.usdt, abi: erc20Abi, functionName: "balanceOf", args: [ADDR.nft] }),
  };
  printStep("balances_after_all", {
    hrx: after.hrx.toString(),
    usdt: after.usdt.toString(),
    nftUsdt: after.nftUsdt.toString(),
    tokenId: tokenId.toString(),
  });
}

main().catch((err) => {
  console.error(err);
  process.exitCode = 1;
});
