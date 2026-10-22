#!/bin/bash
# 挂单 / 撤单 / 市场买 / 注册 / 兑换（串行）。每分钟跑请用 sync-events-minutely.sh。
# 默认 max-time=58（与服务端单次扫链超时一致）。

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(dirname "$SCRIPT_DIR")"
LOG_DIR="$PROJECT_DIR/log"
mkdir -p "$LOG_DIR"

API_BASE="${API_BASE:-http://127.0.0.1:8003}"
MT="${SYNC_CURL_MAX_TIME:-58}"

log() {
  echo "[$(date '+%Y-%m-%d %H:%M:%S')] [rest] $*"
}

log "sync-rest start max-time=$MT"

curl -s -o /dev/null -w "%{http_code}" --connect-timeout 8 --max-time "$MT" "$API_BASE/api/get_box_list_event" | grep -q 200 && log "get_box_list_event ok" || log "get_box_list_event fail"

curl -s -o /dev/null -w "%{http_code}" --connect-timeout 8 --max-time "$MT" "$API_BASE/api/get_box_unlist_event" | grep -q 200 && log "get_box_unlist_event ok" || log "get_box_unlist_event fail"

curl -s -o /dev/null -w "%{http_code}" --connect-timeout 8 --max-time "$MT" "$API_BASE/api/get_box_buy_event" | grep -q 200 && log "get_box_buy_event ok" || log "get_box_buy_event fail"

curl -s -o /dev/null -w "%{http_code}" --connect-timeout 8 --max-time "$MT" "$API_BASE/api/get_user_r_event" | grep -q 200 && log "get_user_r_event ok" || log "get_user_r_event fail"

curl -s -o /dev/null -w "%{http_code}" --connect-timeout 8 --max-time "$MT" "$API_BASE/api/get_exchange_event" | grep -q 200 && log "get_exchange_event ok" || log "get_exchange_event fail"

log "sync-rest end"
