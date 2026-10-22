/**
 * 调用 requestEth POST /api/set_reward，触发链上 BlindBoxNFT.setRewardOnce（需服务端配置 SET_REWARD_ADMIN_PRIVATE_KEY）。
 *
 * reward 为 int256 字符串，18 位精度 wei（与 blind_settle_worker 一致）：1 USDT = 1e18。
 *
 * 用法:
 *   API_BASE=http://127.0.0.1:8003 node scripts/set-reward-api.mjs <tokenId> <rewardWei>
 *   API_BASE=http://47.245.108.129:8003 node scripts/set-reward-api.mjs 123 5000000000000000000
 *   # 亏损示例（-0.5 USDT）:
 *   node scripts/set-reward-api.mjs 123 -500000000000000000
 */
const API_BASE = (process.env.API_BASE || "http://127.0.0.1:8003").replace(/\/$/, "");

const tokenId = process.argv[2];
const reward = process.argv[3];

if (!tokenId || reward === undefined) {
  console.error(
    "用法: node scripts/set-reward-api.mjs <tokenId> <rewardWei>\n" +
      "示例: node scripts/set-reward-api.mjs 42 1000000000000000000"
  );
  process.exit(1);
}

const body = JSON.stringify({
  tokenId: String(tokenId),
  reward: String(reward),
});

const url = `${API_BASE}/api/set_reward`;
const res = await fetch(url, {
  method: "POST",
  headers: { "Content-Type": "application/json" },
  body,
});

const text = await res.text();
let json;
try {
  json = JSON.parse(text);
} catch {
  console.error("HTTP", res.status, text);
  process.exit(1);
}

console.log(json);
const txHash =
  typeof json.res === "string" &&
  json.res.startsWith("0x") &&
  /^0x[0-9a-fA-F]{64}$/.test(json.res);
process.exit(res.ok && txHash ? 0 : 1);
