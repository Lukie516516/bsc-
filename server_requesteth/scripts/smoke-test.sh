#!/usr/bin/env bash
# 模拟前端关键接口自检（可对本机或公网 API_BASE）
# 用法：
#   API_BASE=http://127.0.0.1:8003 ./scripts/smoke-test.sh
#   API_BASE=http://47.x.x.x:8003 ADDR=0x你的地址 ./scripts/smoke-test.sh
set -euo pipefail

API_BASE="${API_BASE:-http://127.0.0.1:8003}"
ADDR="${ADDR:-0xE2D4ecEcb19B39D7e796fF9A4965543AeD903069}"

echo "=========================================="
echo "API_BASE=$API_BASE"
echo "ADDR=$ADDR"
echo "=========================================="

fail=0
http_get() {
  local url="$1"
  local code
  code=$(curl -s -o /tmp/smoke_body.json -w "%{http_code}" --connect-timeout 5 --max-time 60 "$url" || echo "000")
  echo "$code"
}

echo ""
echo "[1] GET /health"
code=$(http_get "$API_BASE/health")
echo "HTTP $code"
if [ "$code" != "200" ]; then fail=1; fi
head -c 200 /tmp/smoke_body.json 2>/dev/null || true
echo ""

echo ""
echo "[2] GET /api/get_address_box（检查 count 与 list 条数、是否重复 tokenId）"
code=$(http_get "$API_BASE/api/get_address_box?address=${ADDR}&page=1&numTwo=0&numThree=1&openStatus=0&num=0")
echo "HTTP $code"
if [ "$code" != "200" ]; then fail=1; fi

if command -v jq &>/dev/null; then
  # count 可能是 JSON 数字或字符串
  cnt=$(jq -r '.count | if type == "string" then tonumber else . end | tostring' /tmp/smoke_body.json 2>/dev/null || echo "")
  n=$(jq '.list | length' /tmp/smoke_body.json 2>/dev/null || echo "0")
  uniq=$(jq '[.list[]?.tokenId] | unique | length' /tmp/smoke_body.json 2>/dev/null || echo "0")
  echo "count=$cnt  list_len=$n  unique_tokenIds=$uniq"
  if [ -n "$cnt" ] && [ "$cnt" != "null" ]; then
    if [ "$cnt" != "$n" ]; then
      echo "FAIL: count 与 list 条数不一致（可能未部署 GetAddressBox 去重修复）"
      fail=1
    fi
  fi
  if [ "$n" != "$uniq" ] && [ "$n" != "0" ]; then
    echo "FAIL: list 中存在重复 tokenId（1234 循环问题）"
    fail=1
  fi
  if [ "$fail" = "0" ]; then echo "OK: 条数与去重检查通过"; fi
else
  echo "（未安装 jq，跳过 JSON 校验；建议 apt install jq）"
  head -c 800 /tmp/smoke_body.json
  echo ""
fi

echo ""
echo "[3] GET /api/debug_mint_sync"
code=$(http_get "$API_BASE/api/debug_mint_sync")
echo "HTTP $code"
jq . /tmp/smoke_body.json 2>/dev/null || head -c 400 /tmp/smoke_body.json
echo ""

echo ""
echo "[4] GET /api/sync_mint_from_contract（只读链上 view，可多次调用）"
code=$(http_get "$API_BASE/api/sync_mint_from_contract")
echo "HTTP $code"
jq . /tmp/smoke_body.json 2>/dev/null || head -c 400 /tmp/smoke_body.json
echo ""

echo ""
if [ "$fail" = "0" ]; then
  echo "=========================================="
  echo "smoke-test: 全部检查通过（或需 jq 做完整校验）"
  echo "=========================================="
else
  echo "=========================================="
  echo "smoke-test: 存在 FAIL，请检查部署镜像是否为最新"
  echo "=========================================="
  exit 1
fi
