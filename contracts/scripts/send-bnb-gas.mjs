/**
 * 从赞助钱包向指定地址转少量原生 BNB（gas）
 *
 * 配置：复制 .env.gas.example 为 .env.gas 并填写 GAS_FUNDER_PRIVATE_KEY
 *
 * 用法（在 box 目录）:
 *   node scripts/send-bnb-gas.mjs
 */
import { readFileSync, existsSync } from "fs";
import { resolve, dirname } from "path";
import { fileURLToPath } from "url";
import {
  createPublicClient,
  createWalletClient,
  http,
  parseEther,
} from "viem";
import { privateKeyToAccount } from "viem/accounts";
import { bsc } from "viem/chains";

const __dirname = dirname(fileURLToPath(import.meta.url));
const boxRoot = resolve(__dirname, "..");
const envPath = resolve(boxRoot, ".env");
const envGasPath = resolve(boxRoot, ".env.gas");

function loadDotEnvFile(path, override = false) {
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
    if (override || process.env[key] === undefined) process.env[key] = val;
  }
}

loadDotEnvFile(envPath, false);
loadDotEnvFile(envGasPath, true);

const pkRaw = (
  process.env.GAS_FUNDER_PRIVATE_KEY ||
  process.env.BSC_PRIVATE_KEY ||
  ""
).trim();
const to = (
  process.env.GAS_TO_ADDRESS ||
  "0x7b257d247bDDF8Fc83B1aDe8e819e231ee4a81FE"
).trim();
const amountStr = (process.env.BNB_AMOUNT || "0.003").trim();
const rpc = process.env.BSC_RPC_URL || "https://bsc-rpc.publicnode.com";

if (!pkRaw || pkRaw.length < 10) {
  console.error(
    "请设置 GAS_FUNDER_PRIVATE_KEY（box/.env.gas）或 BSC_PRIVATE_KEY（box/.env）",
  );
  process.exit(1);
}
if (!to || !/^0x[a-fA-F0-9]{40}$/.test(to)) {
  console.error("请在 box/.env.gas 中设置有效的 GAS_TO_ADDRESS");
  process.exit(1);
}

const pkHex =
  pkRaw.startsWith("0x") ? pkRaw : `0x${pkRaw.replace(/^0x/i, "")}`;
const account = privateKeyToAccount(/** @type {import('viem').Hex} */ (pkHex));
const transport = http(rpc, { timeout: 120_000 });
const wallet = createWalletClient({ account, chain: bsc, transport });
const publicClient = createPublicClient({ chain: bsc, transport });

const value = parseEther(amountStr);

async function main() {
  const bal = await publicClient.getBalance({ address: account.address });
  console.log("赞助钱包:", account.address);
  console.log("当前 BNB:", (Number(bal) / 1e18).toFixed(6));
  console.log("将转出:", amountStr, "BNB ->", to);

  if (bal < value + parseEther("0.0001")) {
    console.error("赞助钱包 BNB 不足（需略高于转账额 + 手续费）");
    process.exit(1);
  }

  const hash = await wallet.sendTransaction({
    to: /** @type {import('viem').Address} */ (to),
    value,
  });
  console.log("tx:", hash);
  await publicClient.waitForTransactionReceipt({ hash });
  const balTo = await publicClient.getBalance({
    address: /** @type {import('viem').Address} */ (to),
  });
  console.log("收款地址 BNB 余额:", (Number(balTo) / 1e18).toFixed(6));
  console.log("完成");
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});
