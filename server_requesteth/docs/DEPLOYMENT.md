# 部署与配置指南

## 环境变量

### 必需（按功能）

| 功能 | 变量 | 说明 |
|------|------|------|
| 铸造分佣打款 | `REFERRAL_PAYOUT_PRIVATE_KEY` | 打款私钥 |
| 铸造分佣打款 | `REFERRAL_USDT_ADDR` | USDT 合约地址 |
| SetReward 接口 | `SET_REWARD_ADMIN_PRIVATE_KEY` | 调用 setRewardOnce 的私钥 |
| AICAT 接口 | `AICAT_ADMIN_PRIVATE_KEY` | AddLiquidity 等 |
| 盲盒结算 Worker | `BLIND_SETTLE_PRIVATE_KEY` | 见 blind_settle_worker |
| 调仓 Worker | `REBALANCE_PRIVATE_KEY` | 见 rebalance_worker |

### 链上配置（可选，有默认值）

| 变量 | 默认值 |
|------|--------|
| `ETH_BSC_RPC_URLS` | 公开 BSC 节点列表 |
| `ETH_BSC_RPC_URLS_POLLING` | 事件轮询用 RPC（可配置付费节点） |
| `ETH_NFT_CONTRACT` | 见 internal/conf/eth.go |
| `ETH_STAKE_CONTRACT` | 同上 |
| `ETH_USDT_CONTRACT` | BSC USDT |
| ... | 更多见 docs/INTERFACES.md |

### 当前正式版建议直接配置为

```env
ETH_NFT_CONTRACT=0x169c4E946909D4E115b724957e5D1c74d0971f70
ETH_STAKE_CONTRACT=0x907F84e4127049bB65C82F50CE29c1FF67f344D5
ETH_SWAP_CONTRACT=0x3c9Ea8d476eA06108283457B3Bb0dfbfD6cf1A1b
ETH_PRIMARY_MARKET_CONTRACT=0xdA5A5BC6b61B2eAB091A1b7E4d84C7eb537BDC0e
ETH_VAULT_FACTORY_CONTRACT=0x7AD5C447580A79e072bFD72E373f7b25dDE0F061
ETH_USDT_CONTRACT=0x55d398326f99059fF775485246999027B3197955
```

### Worker 相关变量也要一起换成正式版

```env
NFT_ADDR=0x169c4E946909D4E115b724957e5D1c74d0971f70
STAKE_ADDR=0x907F84e4127049bB65C82F50CE29c1FF67f344D5
USDT_ADDR=0x55d398326f99059fF775485246999027B3197955
EXCHANGE_ADDR=0x3c9Ea8d476eA06108283457B3Bb0dfbfD6cf1A1b
```

### 上线前特别注意

- 不要把本地 `.env` 原样上传到服务器
- 私钥只能在宝塔服务器环境变量里单独填写，禁止提交到仓库或打进镜像
- 如果服务器此前跑的是旧版 `BlindBoxMainnetTest`，必须同步替换 API 和 worker 使用的全部合约地址
- `REQUESTETH_ADMIN_API_KEY` 生产环境必须配置，否则管理接口默认拒绝访问

## 部署步骤

### 方式一：Docker Compose（推荐）

```bash
cd requestEth
cp .env.example .env   # 填写私钥、合约地址等
docker compose up -d
```

- 使用 `docker/configs/config.yaml`（mysql/redis 已配置为 compose 服务名）
- 服务：api、blind_settle_worker、referral_worker、rebalance_worker、mysql、redis
- 首次启动需执行数据库迁移（连接 mysql 后手动执行或使用迁移工具）

### 方式二：本地运行

1. 复制 `configs/config.yaml` 并修改数据库、Redis 等
2. 设置环境变量（私钥、合约地址等）
3. 执行数据库迁移
4. 启动 API：`go run ./cmd/requestEth -conf=./configs`
5. 启动 Worker：`blind_settle_worker`、`rebalance_worker`、`referral_worker`

## 安全建议

- 私钥仅通过环境变量传入，禁止写入配置文件
- 生产环境使用独立的 RPC 节点（付费节点更稳定）
- 定期轮换打款私钥

## 健康检查

- `GET /health` 返回 `{"status":"ok"}`，用于负载均衡/容器探活
- 不经过限流，优先响应

## 限流

HTTP 服务已启用 BBR 限流中间件（`ratelimit.Server()`），触发时返回 429。可配置自定义 Limiter 替换默认算法。

## 优雅关闭

`app.Run()` 监听 SIGINT/SIGTERM，收到信号后等待进行中请求完成再退出。部署时使用 `kill -TERM <pid>` 或 `docker stop` 触发。
