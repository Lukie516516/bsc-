/**
 * 链上验证：BlindBoxNFT mint(tier=1) + openBoxPayHRX（自铸自开）
 *
 * 依赖 box/.env 的 BSC_PRIVATE_KEY（需有 MockUSDT、HRX、BNB）
 *
 * 用法: node scripts/verify-nft-flow.mjs
 */
import { readFileSync, existsSync } from "fs";
import { resolve, dirname } from "path";
import { fileURLToPath } from "url";
import {
  createPublicClient,
  createWalletClient,
  decodeEventLog,
  http,
  parseAbi,
  parseEther,
} from "viem";
import { privateKeyToAccount } from "viem/accounts";
import { bsc } from "viem/chains";

const __dirname = dirname(fileURLToPath(import.meta.url));
const boxRoot = resolve(__dirname, "..");

function loadEnv(path) {
  if (!existsSync(path)) return;
  const raw = readFileSync(path, "utf8");
  for (const line of raw.split(/\r?\n/)) {
    const t = line.trim();
    if (!t || t.startsWith("#")) continue;
    const eq = t.indexOf("=");
    if (eq <= 0) continue;
    const k = t.slice(0, eq).trim();
    let v = t.slice(eq + 1).trim();
    if (
      (v.startsWith('"') && v.endsWith('"')) ||
      (v.startsWith("'") && v.endsWith("'"))
    ) {
      v = v.slice(1, -1);
    }
    if (process.env[k] === undefined) process.env[k] = v;
  }
}

loadEnv(resolve(boxRoot, ".env"));

const NFT = "0x73F76b11ebF03718B71c163BE2B5445ECFC25Ce6";
const USDT = "0x2674DE1Fe0c884D45cd128b3F67Eb03BFD649ff4";
const HRX = "0xE3198bf629C91bd678190D22393Ba5DEc40E9106";

const pkRaw = (process.env.BSC_PRIVATE_KEY || "").trim();
const rpc = process.env.BSC_RPC_URL || "https://bsc-rpc.publicnode.com";

const erc20 = parseAbi([
  "function approve(address spender, uint256 amount) returns (bool)",
  "function balanceOf(address) view returns (uint256)",
]);
const nftAbi = parseAbi([
  "function mint(uint8 tier) returns (uint256 tokenId)",
  "function openBoxPayHRX(uint256 tokenId, uint256 maxFeeB)",
  "function ownerOf(uint256 tokenId) view returns (address)",
  "function open(uint256 tokenId) view returns (bool)",
  "function openFeeHrxReceiver() view returns (address)",
  "function feeQuotePool() view returns (address)",
  "event Minted(address indexed minter, uint256 indexed tokenId, uint8 indexed tier, uint256 usdtPaid)",
  "event Opened(address indexed owner, uint256 indexed tokenId, uint256 timestamp, uint256 feeU, uint256 feeB)",
]);

if (!pkRaw || pkRaw.length < 10) {
  console.error("需要 BSC_PRIVATE_KEY（box/.env）");
  process.exit(1);
}

const pkHex = pkRaw.startsWith("0x") ? pkRaw : `0x${pkRaw}`;
const account = privateKeyToAccount(/** @type {import("viem").Hex} */ (pkHex));
const transport = http(rpc, { timeout: 120_000 });
const wallet = createWalletClient({ account, chain: bsc, transport });
const publicClient = createPublicClient({ chain: bsc, transport });

async function nextNonce() {
  return publicClient.getTransactionCount({
    address: account.address,
    blockTag: "pending",
  });
}

async function main() {
  console.log("账户:", account.address);
  const pool = await publicClient.readContract({
    address: NFT,
    abi: nftAbi,
    functionName: "feeQuotePool",
  });
  const recv = await publicClient.readContract({
    address: NFT,
    abi: nftAbi,
    functionName: "openFeeHrxReceiver",
  });
  console.log("feeQuotePool:", pool);
  console.log("openFeeHrxReceiver:", recv);

  let n = await nextNonce();
  const approveU = await wallet.writeContract({
    address: USDT,
    abi: erc20,
    functionName: "approve",
    args: [NFT, parseEther("1000")],
    nonce: n,
  });
  console.log("approve USDT -> NFT tx:", approveU);
  await publicClient.waitForTransactionReceipt({ hash: approveU });

  n = await nextNonce();
  const mintTx = await wallet.writeContract({
    address: NFT,
    abi: nftAbi,
    functionName: "mint",
    args: [1],
    nonce: n,
  });
  console.log("mint(1) tx:", mintTx);
  const mintRc = await publicClient.waitForTransactionReceipt({ hash: mintTx });

  let tokenId = null;
  for (const log of mintRc.logs) {
    try {
      const ev = decodeEventLog({
        abi: nftAbi,
        data: log.data,
        topics: log.topics,
      });
      if (ev.eventName === "Minted") tokenId = ev.args.tokenId;
    } catch {
      /* skip */
    }
  }
  if (tokenId == null) throw new Error("未解析到 Minted.tokenId");
  console.log("tokenId:", tokenId.toString());

  const owner = await publicClient.readContract({
    address: NFT,
    abi: nftAbi,
    functionName: "ownerOf",
    args: [tokenId],
  });
  console.log("ownerOf:", owner);
  if (owner.toLowerCase() !== account.address.toLowerCase()) {
    throw new Error("owner 不匹配");
  }

  n = await nextNonce();
  const approveH = await wallet.writeContract({
    address: HRX,
    abi: erc20,
    functionName: "approve",
    args: [NFT, parseEther("1000000")],
    nonce: n,
  });
  console.log("approve HRX -> NFT tx:", approveH);
  await publicClient.waitForTransactionReceipt({ hash: approveH });

  n = await nextNonce();
  const maxFeeB = parseEther("100");
  const openTx = await wallet.writeContract({
    address: NFT,
    abi: nftAbi,
    functionName: "openBoxPayHRX",
    args: [tokenId, maxFeeB],
    nonce: n,
  });
  console.log("openBoxPayHRX tx:", openTx);
  await publicClient.waitForTransactionReceipt({ hash: openTx });

  const opened = await publicClient.readContract({
    address: NFT,
    abi: nftAbi,
    functionName: "open",
    args: [tokenId],
  });
  console.log("open[tokenId]:", opened);

  console.log("\n✅ NFT 铸造 + HRX 开盒：成功");
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});
