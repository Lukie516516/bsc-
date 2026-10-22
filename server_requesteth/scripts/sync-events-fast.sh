#!/bin/bash
# 轻量同步：铸造 + 转账 + 开盒 + update_box（串行）。
# 每分钟跑请优先用 sync-events-minutely.sh（并行，总耗时更短）。
# curl --max-time 须 ≥ 服务端 syncCronCtxTimeout（默认 55s），此处默认 58。
#
#   * * * * * .../sync-events-fast.sh >> .../log/sync-events-fast.log 2>&1

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(dirname "$SCRIPT_DIR")"
LOG_DIR="$PROJECT_DIR/log"
mkdir -p "$LOG_DIR"

API_BASE="${API_BASE:-http://127.0.0.1:8003}"
MT="${SYNC_CURL_MAX_TIME:-58}"

log() {
  echo "[$(date '+%Y-%m-%d %H:%M:%S')] [fast] $*"
}

log "sync-fast start max-time=$MT"

if curl -s -o /dev/null -w "%{http_code}" --connect-timeout 8 --max-time "$MT" "$API_BASE/api/get_box_mint_event" | grep -q 200; then
  log "get_box_mint_event ok"
else
  log "get_box_mint_event fail, try sync_mint_from_contract"
  if curl -s -o /dev/null -w "%{http_code}" --connect-timeout 8 --max-time 300 "$API_BASE/api/sync_mint_from_contract" | grep -q 200; then
    log "sync_mint_from_contract ok"
  else
    log "sync_mint_from_contract fail"
  fi
fi

curl -s -o /dev/null -w "%{http_code}" --connect-timeout 8 --max-time "$MT" "$API_BASE/api/get_box_transfer_event" | grep -q 200 && log "get_box_transfer_event ok" || log "get_box_transfer_event fail"

if curl -s -o /dev/null -w "%{http_code}" --connect-timeout 8 --max-time "$MT" "$API_BASE/api/get_box_open_event" | grep -q 200; then
  log "get_box_open_event ok"
else
  log "get_box_open_event fail"
fi

if curl -s -o /dev/null -w "%{http_code}" --connect-timeout 8 --max-time "$MT" "$API_BASE/api/update_box" | grep -q 200; then
  log "update_box ok"
else
  log "update_box fail"
fi

log "sync-fast end"
