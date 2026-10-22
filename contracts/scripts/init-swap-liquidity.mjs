/**
 * 为 SwapHrxUsdt（VITE_devis）注入双边代币并调用 mintInitial。
 * 部署后若未执行，liquidityInited=false、储备为 0，前端池子报价/路由会失败。
 *
 * 依赖：box/.env 中 BSC_PRIVATE_KEY 须为 Swap.owner()（部署账号）
 *
 * 用法（在 box 目录）:
 *   node scripts/init-swap-liquidity.mjs
 *
 * 可选环境变量（人类可读数量，18 位）:
 *   INIT_SWAP_HRX=100000
 *   INIT_SWAP_USDT=100000
 */
import { readFileSync, existsSync } from "fs";
import { resolve, dirname } from "path";
import { fileURLToPath } from "url";
import {
  createPublicClient,
  createWalletClient,
  http,
  parseAbi,
  parseEther,
} from "viem";
import { privateKeyToAccount } from "viem/accounts";
import { bsc } from "viem/chains";

const __dirname = dirname(fileURLToPath(import.meta.url));
const boxRoot = resolve(__dirname, "..");

function loadDotEnvFile(path) {
  if (!existsSync(path)) return;
  const raw = readFileSync(path, "utf8");
  for (const line of raw.split(/\r?\n/)) {
    const t = line.trim();
    if (!t || t.startsWith("#")) continue;
    const eq = t.indexOf("=");
    if (eq <= 0) continue;
    const key = t.slice(0, eq).trim();
    let val = t.slice(eq + 1).trim();
    if (
      (val.startsWith('"') && val.endsWith('"')) ||
      (val.startsWith("'") && val.endsWith("'"))
    ) {
      val = val.slice(1, -1);
    }
    if (process.env[key] === undefined) process.env[key] = val;
  }
}

loadDotEnvFile(resolve(boxRoot, ".env"));

const SWAP = "0x3Fc7A08C7af8B8300a6C56475cc32c72534eE7E4";
const HRX = "0xE3198bf629C91bd678190D22393Ba5DEc40E9106";
const USDT = "0x2674DE1Fe0c884D45cd128b3F67Eb03BFD649ff4";

const pkRaw = (process.env.BSC_PRIVATE_KEY || "").trim();
const hrxHuman = process.env.INIT_SWAP_HRX || "100000";
const usdtHuman = process.env.INIT_SWAP_USDT || "100000";
const hrxAmt = parseEther(hrxHuman);
const usdtAmt = parseEther(usdtHuman);
const rpc = process.env.BSC_RPC_URL || "https://bsc-rpc.publicnode.com";

const erc20 = parseAbi([
  "function transfer(address to, uint256 amount) returns (bool)",
]);
const swapAbi = parseAbi([
  "function liquidityInited() view returns (bool)",
  "function owner() view returns (address)",
  "function mintInitial(uint256 _b1) returns (uint256 amountDL, uint256 amountOther)",
]);

if (!pkRaw || pkRaw.length < 10) {
  console.error("需要 box/.env 中的 BSC_PRIVATE_KEY（须为 Swap owner）");
  process.exit(1);
}

const pkHex = pkRaw.startsWith("0x") ? pkRaw : `0x${pkRaw}`;
const account = privateKeyToAccount(/** @type {import('viem').Hex} */ (pkHex));
const transport = http(rpc, { timeout: 120_000 });
const wallet = createWalletClient({ account, chain: bsc, transport });
const publicClient = createPublicClient({ chain: bsc, transport });

async function main() {
  const inited = await publicClient.readContract({
    address: SWAP,
    abi: swapAbi,
    functionName: "liquidityInited",
  });
  if (inited) {
    console.log("SwapHrxUsdt 已 liquidityInited，跳过");
    return;
  }

  const owner = await publicClient.readContract({
    address: SWAP,
    abi: swapAbi,
    functionName: "owner",
  });
  if (owner.toLowerCase() !== account.address.toLowerCase()) {
    console.error("当前私钥地址", account.address, "不是 Swap owner", owner);
    process.exit(1);
  }

  console.log("向 Swap 注入 HRX", hrxHuman, "、USDT", usdtHuman);
  let nonce = await publicClient.getTransactionCount({
    address: account.address,
    blockTag: "pending",
  });
  const tx1 = await wallet.writeContract({
    address: HRX,
    abi: erc20,
    functionName: "transfer",
    args: [SWAP, hrxAmt],
    nonce,
  });
  console.log("HRX transfer tx:", tx1);
  await publicClient.waitForTransactionReceipt({ hash: tx1 });

  nonce = await publicClient.getTransactionCount({
    address: account.address,
    blockTag: "pending",
  });
  const tx2 = await wallet.writeContract({
    address: USDT,
    abi: erc20,
    functionName: "transfer",
    args: [SWAP, usdtAmt],
    nonce,
  });
  console.log("USDT transfer tx:", tx2);
  await publicClient.waitForTransactionReceipt({ hash: tx2 });

  nonce = await publicClient.getTransactionCount({
    address: account.address,
    blockTag: "pending",
  });
  const tx3 = await wallet.writeContract({
    address: SWAP,
    abi: swapAbi,
    functionName: "mintInitial",
    args: [0n],
    nonce,
  });
  console.log("mintInitial tx:", tx3);
  await publicClient.waitForTransactionReceipt({ hash: tx3 });

  const ok = await publicClient.readContract({
    address: SWAP,
    abi: swapAbi,
    functionName: "liquidityInited",
  });
  console.log("liquidityInited:", ok);
  console.log("完成");
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});
