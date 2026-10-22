/**
 * BSC 主网：mint(tier=1) 一笔 → 触发远端 requestEth 同步 → 校验 get_address_box
 *
 * 用法:
 *   API_BASE=http://47.245.108.129:8003 node scripts/mint-and-verify-api.mjs
 *   node scripts/mint-and-verify-api.mjs --verify-apis-only   # 只跑读接口验收，不铸造
 *
 * 依赖 box/.env: BSC_PRIVATE_KEY, 可选 BSC_RPC_URL
 */
import { execFile } from "node:child_process";
import { promisify } from "node:util";
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
const execFileAsync = promisify(execFile);
const curlBin = process.platform === "win32" ? "curl.exe" : "curl";

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

const NFT = process.env.VERIFY_NFT_ADDR || "0x73F76b11ebF03718B71c163BE2B5445ECFC25Ce6";
const USDT = process.env.VERIFY_USDT_ADDR || "0x2674DE1Fe0c884D45cd128b3F67Eb03BFD649ff4";
const API_BASE = (process.env.API_BASE || "http://127.0.0.1:8003").replace(/\/$/, "");
// get_box_new / get_box_open 的 start、end 为链上 tokenId，不是 Unix 时间；勿用 end=9999999999
const BOX_TOKEN_END = process.env.VERIFY_GET_BOX_TOKEN_END || "2000";

const pkRaw = (process.env.BSC_PRIVATE_KEY || "").trim();
const rpc = process.env.BSC_RPC_URL || "https://bsc-rpc.publicnode.com";

const erc20 = parseAbi([
  "function approve(address spender, uint256 amount) returns (bool)",
]);
const nftAbi = parseAbi([
  "function mint(uint8 tier) returns (uint256 tokenId)",
  "function ownerOf(uint256 tokenId) view returns (address)",
  "event Minted(address indexed minter, uint256 indexed tokenId, uint8 indexed tier, uint256 usdtPaid)",
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

async function sleep(ms) {
  await new Promise((r) => setTimeout(r, ms));
}

/** 优先 curl（部分代理下 Node fetch 易断连）；失败则回退 fetch */
async function httpGet(url) {
  try {
    const { stdout } = await execFileAsync(
      curlBin,
      [
        "-sS",
        "--connect-timeout",
        "30",
        "--max-time",
        "120",
        "-w",
        "\n%{http_code}",
        url,
      ],
      { maxBuffer: 50 * 1024 * 1024 },
    );
    const norm = stdout.replace(/\r\n/g, "\n");
    const m = norm.match(/^([\s\S]*)\n(\d{3})$/);
    const text = m ? m[1] : norm;
    const status = m ? parseInt(m[2], 10) : 0;
    if (status > 0) return { ok: status >= 200 && status < 300, status, text };
  } catch {
    /* fall through */
  }
  const r = await fetch(url, { signal: AbortSignal.timeout(120_000) });
  const t = await r.text();
  return { ok: r.ok, status: r.status, text: t };
}

async function httpGetRetry(url, attempts = 3, gapMs = 800) {
  let lastErr;
  for (let i = 0; i < attempts; i++) {
    try {
      return await httpGet(url);
    } catch (e) {
      lastErr = e;
      if (i < attempts - 1) await sleep(gapMs * (i + 1));
    }
  }
  throw lastErr;
}

/** 主读接口 + 余额：HTTP 200 且可解析为 JSON（请求间间隔，避免服务端限流断连） */
async function verifyReadApis(minterAddr, phaseLabel) {
  const a = encodeURIComponent(minterAddr);
  const paths = [
    "/health",
    "/api/debug_mint_sync",
    "/api/get_box_all",
    `/api/get_box_new?start=0&end=${BOX_TOKEN_END}`,
    `/api/get_box_open?start=0&end=${BOX_TOKEN_END}`,
    "/api/get_market_box_list?page=1&num=10&numTwo=0&numThree=0",
    `/api/get_buy_box_list?address=${a}&page=1`,
    `/api/get_sell_box_list?address=${a}&page=1`,
    "/api/get_all_info?start=0&end=9999999999",
    `/api/eth_balance?address=${a}`,
    `/api/usdt_balance?address=${a}`,
    `/api/get_address_box?address=${a}&page=1&numTwo=0&numThree=1&openStatus=2&num=0&start=0&end=9999999999`,
    `/api/address_box_v2?address=${a}&page=1&numTwo=0&numThree=1&openStatus=2&num=0&start=0&end=9999999999`,
  ];
  const pace = Number(process.env.API_VERIFY_PACE_MS || 600);
  console.log(`\n========== 读接口验收 ${phaseLabel} ==========`);
  for (const p of paths) {
    const url = `${API_BASE}${p}`;
    const res = await httpGetRetry(url);
    if (!res.ok) {
      throw new Error(`${phaseLabel} ${p} HTTP ${res.status}: ${res.text.slice(0, 200)}`);
    }
    try {
      JSON.parse(res.text);
    } catch {
      throw new Error(`${phaseLabel} ${p} 非 JSON: ${res.text.slice(0, 200)}`);
    }
    console.log("OK", p.split("?")[0]);
    await sleep(pace);
  }
  console.log(`========== 读接口验收 ${phaseLabel} 全部通过 ==========\n`);
}

async function main() {
  console.log("账户:", account.address);
  console.log("NFT:", NFT);
  console.log("API_BASE:", API_BASE);

  await verifyReadApis(account.address, "铸造前");

  let n = await publicClient.getTransactionCount({
    address: account.address,
    blockTag: "pending",
  });

  const approveU = await wallet.writeContract({
    address: USDT,
    abi: erc20,
    functionName: "approve",
    args: [NFT, parseEther("1000")],
    nonce: n,
  });
  console.log("approve USDT tx:", approveU);
  await publicClient.waitForTransactionReceipt({ hash: approveU });

  n = await publicClient.getTransactionCount({
    address: account.address,
    blockTag: "pending",
  });
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
  console.log("新 tokenId:", tokenId.toString());

  const owner = await publicClient.readContract({
    address: NFT,
    abi: nftAbi,
    functionName: "ownerOf",
    args: [tokenId],
  });
  console.log("链上 ownerOf:", owner);

  async function syncOnce(label) {
    console.log(`\n--- 同步一轮 ${label} ---`);
    for (const path of [
      "/api/get_box_mint_event",
      "/api/get_box_transfer_event",
      "/api/update_box",
      "/api/sync_mint_from_contract",
    ]) {
      const url = `${API_BASE}${path}`;
      const res = await httpGet(url);
      console.log(path, "HTTP", res.status, res.text.slice(0, 160));
    }
  }

  const addr = account.address;
  const q = `${API_BASE}/api/get_address_box?address=${encodeURIComponent(addr)}&page=1&numTwo=0&numThree=1&openStatus=0&num=2`;

  async function fetchBox() {
    const boxRes = await httpGet(q);
    const data = JSON.parse(boxRes.text);
    const list = data.list || [];
    const ids = list.map((x) => String(x.tokenId ?? x.token_id ?? ""));
    return { data, list, ids, httpOk: boxRes.ok };
  }

  await syncOnce("(mint 后立即)");

  const waitMs = Number(process.env.POST_MINT_WAIT_MS || 120_000);
  console.log(`\n--- 等待 ${waitMs / 1000}s（计划任务 / 游标推进） ---`);
  await new Promise((r) => setTimeout(r, waitMs));

  const maxRounds = Number(process.env.POST_MINT_POLL_ROUNDS || 12);
  const pollGapMs = Number(process.env.POST_MINT_POLL_GAP_MS || 10_000);
  let has = false;
  let lastIds = [];
  for (let round = 1; round <= maxRounds; round++) {
    await syncOnce(`(轮询 ${round}/${maxRounds})`);
    const { data, ids } = await fetchBox();
    lastIds = ids;
    has = ids.includes(tokenId.toString());
    console.log(
      "get_address_box count=",
      data.count,
      "ids=",
      ids.join(","),
      "含新tokenId?",
      has,
    );
    if (has) break;
    if (round < maxRounds) await new Promise((r) => setTimeout(r, pollGapMs));
  }

  const dbg = await httpGet(`${API_BASE}/api/debug_mint_sync`);
  console.log("\n--- debug_mint_sync ---\n", dbg.text.slice(0, 400));

  if (!has) {
    console.log("最终 tokenIds:", lastIds.join(","));
    throw new Error(
      "等待+轮询后 get_address_box 仍未包含刚铸造的 tokenId；请查库 nft_minted.to_addr / 游标 / 环境 ETH_NFT_CONTRACT",
    );
  }
  await verifyReadApis(account.address, "铸造并同步后");
  console.log("\n✅ 链上 mint 成功，get_address_box 已含新 token，读接口均 200+JSON");
}

const argv = new Set(process.argv.slice(2));
if (argv.has("--verify-apis-only")) {
  verifyReadApis(account.address, "仅读接口")
    .then(() => {
      console.log("✅ --verify-apis-only 完成");
      process.exit(0);
    })
    .catch((e) => {
      console.error(e);
      process.exit(1);
    });
} else {
  main().catch((e) => {
    console.error(e);
    process.exit(1);
  });
}
