# 链上事件同步脚本说明

目标：**在可接受范围内**（通常约 1 分钟内）把铸造 / 转账 / 开盒 / 合并写入库；挂单与市场等可同频或低频。

## 脚本分工

| 脚本 | 内容 | 建议周期 |
|------|------|----------|
| `sync-events-fast.sh` | `get_box_mint_event`（失败则 `sync_mint_from_contract`）→ `get_box_transfer_event` → `get_box_open_event` → `update_box` | **每 1 分钟** |
| `sync-events-rest.sh` | `get_box_list_event` → `get_box_unlist_event` → `get_box_buy_event` → `get_user_r_event` → `get_exchange_event` | **每 1 分钟**（与挂单/市场同频）或 **每 5 分钟**（减负） |
| `sync-events-primary-hrx.sh` | `get_buy_event` → `get_sell_event`（`primary_buy` / `primary_sell`，首页 `get_buy_list` 折线图等） | **每 3～5 分钟**（勿 10s 叠跑；周期 ≥ `--max-time`） |
| `sync-events.sh` | 依次执行 fast + rest，**单次耗时长**，不推荐作为唯一的高频任务 | 仅手工或低频全量 |

## 宝塔计划任务（示例）

**API 在宿主机 `8003` 时**（按实际端口改）：

```bash
# ① 高频（必开）：关闭「进程锁」
* * * * * /www/wwwroot/NewrequestEth/requestEth/scripts/sync-events-fast.sh >> /www/wwwroot/NewrequestEth/requestEth/log/sync-events-fast.log 2>&1

# ② 挂单/市场：需要约 1 分钟内进库则同样每分钟；否则 */5 或 */10
* * * * * /www/wwwroot/NewrequestEth/requestEth/scripts/sync-events-rest.sh >> /www/wwwroot/NewrequestEth/requestEth/log/sync-events-rest.log 2>&1

# ③ 一级市场 HRX 铸造/销毁事件（get_buy_list 数据源）
*/5 * * * * /www/wwwroot/NewrequestEth/requestEth/scripts/sync-events-primary-hrx.sh >> /www/wwwroot/NewrequestEth/requestEth/log/sync-events-primary-hrx.log 2>&1
```

环境变量（可选）：

```bash
export API_BASE=http://127.0.0.1:8003
```

首次部署：

```bash
chmod +x /path/to/requestEth/scripts/sync-events-fast.sh
chmod +x /path/to/requestEth/scripts/sync-events-rest.sh
chmod +x /path/to/requestEth/scripts/sync-events-primary-hrx.sh
mkdir -p /path/to/requestEth/log
```

## 进程锁

- **fast**：建议 **不要** 开进程锁；若单次执行经常超过定时间隔，会被跳过或积压。
- **rest**：接口较慢时可开进程锁，并把 **执行周期 ≥ 单次实际耗时**（看 `sync-events-rest.log` 里 start→end）。

## 不能保证「严格小于 60 秒」的情况

单次 `curl` 带 `--max-time`，RPC 扫块慢时可能接近上限；属基础设施限制，非仅靠改 cron 可消除。

## 与 `.env` / 服务

- `ETH_BSC_RPC_URLS`、`ETH_BSC_RPC_URLS_POLLING` 使用稳定节点。
- API 容器需健康；生产库可设 `REQUESTETH_SKIP_AUTOMIGRATE=1`（见 `.env.example`）。
