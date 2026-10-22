#!/bin/bash
# 全量 = fast + rest（单次耗时长）。生产请拆成两个计划任务，见 scripts/README.md
# fast：铸造/转账/开盒/合并；rest：挂单/市场等
#
#   * * * * * .../sync-events.sh >> .../log/sync-events.log 2>&1

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(dirname "$SCRIPT_DIR")"
LOG_DIR="$PROJECT_DIR/log"
mkdir -p "$LOG_DIR"

API_BASE="${API_BASE:-http://127.0.0.1:8003}"
export API_BASE

"$SCRIPT_DIR/sync-events-fast.sh"
"$SCRIPT_DIR/sync-events-rest.sh"
