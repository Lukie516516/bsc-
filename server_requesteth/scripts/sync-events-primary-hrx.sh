#!/bin/bash
# 一级市场 HRX：primary_buy / primary_sell。
# 已并入 sync-events-minutely.sh；本脚本供只跑一级市场时使用。默认 max-time=58。

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(dirname "$SCRIPT_DIR")"
LOG_DIR="$PROJECT_DIR/log"
mkdir -p "$LOG_DIR"

API_BASE="${API_BASE:-http://127.0.0.1:8003}"
MT="${SYNC_CURL_MAX_TIME:-58}"

log() {
  echo "[$(date '+%Y-%m-%d %H:%M:%S')] [primary-hrx] $*"
}

log "sync-primary-hrx start max-time=$MT"

code_buy="$(curl -sS -o /dev/null -w "%{http_code}" --connect-timeout 8 --max-time "$MT" "$API_BASE/api/get_buy_event" || echo "000")"
if echo "$code_buy" | grep -qE '^200$'; then
  log "get_buy_event ok (http $code_buy)"
else
  log "get_buy_event fail (http $code_buy)"
fi

code_sell="$(curl -sS -o /dev/null -w "%{http_code}" --connect-timeout 8 --max-time "$MT" "$API_BASE/api/get_sell_event" || echo "000")"
if echo "$code_sell" | grep -qE '^200$'; then
  log "get_sell_event ok (http $code_sell)"
else
  log "get_sell_event fail (http $code_sell)"
fi

log "sync-primary-hrx end"
