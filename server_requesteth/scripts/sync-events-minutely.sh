#!/bin/bash
# 每分钟抓链上事件：先铸造（失败则 view 回退），其余接口并行，整轮墙钟约等于单次 max-time（默认 58s）。
#
# 依赖：服务端已按 MaxBlocksPerCronTick（约 6000 块/次）+ 55s context 限制单次请求，避免叠跑拖死 RPC。
#
# 宝塔 / crontab（每 1 分钟）：
#   * * * * * /bin/bash /path/to/requestEth/scripts/sync-events-minutely.sh >> /path/to/requestEth/log/sync-minutely.log 2>&1
#
# 可选：export API_BASE=http://127.0.0.1:8003
#
# 防叠跑（Linux 有 flock 时）：上一分钟未跑完则跳过本轮
#   * * * * * flock -n /tmp/nft-sync-minutely.lock -c '/bin/bash /path/to/sync-events-minutely.sh >> ...'

set -u

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(dirname "$SCRIPT_DIR")"
LOG_DIR="$PROJECT_DIR/log"
mkdir -p "$LOG_DIR"

API_BASE="${API_BASE:-http://127.0.0.1:8003}"
MT="${SYNC_CURL_MAX_TIME:-58}"
CONNECT="${SYNC_CURL_CONNECT:-8}"

ts() { date '+%Y-%m-%d %H:%M:%S'; }
log() { echo "[$(ts)] [minutely] $*"; }

log "start API_BASE=$API_BASE max-time=$MT"

# ---------- 1) 铸造（串行；失败则 sync_mint，HTTP 侧 5 分钟超时）----------
if curl -sS -o /dev/null -w "%{http_code}" --connect-timeout "$CONNECT" --max-time "$MT" \
  "$API_BASE/api/get_box_mint_event" | grep -q 200; then
  log "get_box_mint_event ok"
else
  log "get_box_mint_event fail, try sync_mint_from_contract"
  if curl -sS -o /dev/null -w "%{http_code}" --connect-timeout "$CONNECT" --max-time 300 \
    "$API_BASE/api/sync_mint_from_contract" | grep -q 200; then
    log "sync_mint_from_contract ok"
  else
    log "sync_mint_from_contract fail"
  fi
fi

# ---------- 2) 其余扫块接口并行 ----------
run_bg() {
  local name="$1" path="$2"
  (
    code="$(curl -sS -o /dev/null -w "%{http_code}" --connect-timeout "$CONNECT" --max-time "$MT" \
      "$API_BASE$path" 2>/dev/null || echo "000")"
    echo "[$(ts)] [minutely] $name http=$code"
  ) &
}

run_bg get_box_transfer_event   /api/get_box_transfer_event
run_bg get_box_open_event       /api/get_box_open_event
run_bg update_box               /api/update_box
run_bg get_box_list_event       /api/get_box_list_event
run_bg get_box_unlist_event     /api/get_box_unlist_event
run_bg get_box_buy_event        /api/get_box_buy_event
run_bg get_user_r_event         /api/get_user_r_event
run_bg get_exchange_event       /api/get_exchange_event
run_bg get_box_reward_event     /api/get_box_reward_event
run_bg get_buy_event            /api/get_buy_event
run_bg get_sell_event           /api/get_sell_event

wait
log "end"
