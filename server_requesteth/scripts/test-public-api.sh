#!/usr/bin/env bash
# 批量探测 HTTP 接口（建议在服务器或能访问公网 API 的机器上执行）
# 用法：API_BASE=http://47.x.x.x:8003 bash scripts/test-public-api.sh
set -euo pipefail
API_BASE="${API_BASE:-http://127.0.0.1:8003}"
ADDR="${ADDR:-0xE2D4ecEcb19B39D7e796fF9A4965543AeD903069}"

code() { curl -s -o /tmp/tpa_body.txt -w "%{http_code}" --connect-timeout 5 "$@" || echo "000"; }

echo "API_BASE=$API_BASE"
echo "---------- 快接口 ----------"
for path in \
  "/health" \
  "/api/get_box_all" \
  "/api/debug_mint_sync" \
  "/api/sync_mint_from_contract" \
  "/api/referral_pending?page=1&pageSize=10" \
  "/api/get_address_box?address=${ADDR}&page=1&numTwo=0&numThree=1&openStatus=0&num=0" \
  "/api/get_market_box_list?page=1&pageSize=5" \
  "/api/get_all_info" \
  ; do
  c=$(code --max-time 60 "${API_BASE}${path}")
  echo "[$c] GET $path"
  head -c 160 /tmp/tpa_body.txt; echo ""
done

echo ""
echo "---------- 慢接口（链上扫描，可能 1～3 分钟，勿用短超时）----------"
for path in \
  "/api/get_box_mint_event" \
  "/api/get_box_open_event" \
  "/api/get_box_list_event" \
  "/api/get_box_unlist_event" \
  "/api/get_box_buy_event" \
  "/api/get_box_transfer_event" \
  "/api/get_user_r_event" \
  "/api/get_exchange_event" \
  ; do
  c=$(code --max-time 200 "${API_BASE}${path}")
  echo "[$c] GET $path (max-time 200s)"
  head -c 120 /tmp/tpa_body.txt; echo ""
done

echo "完成。"
