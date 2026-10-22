# 盲盒结算 Worker

监听 `nft_opened` 表中 `check_status=0` 的开盒记录，按随机规则生成盈亏，调用 NFT 合约的 `setRewardOnce`，并更新数据库。

## 流程

1. 轮询 `nft_opened WHERE check_status=0`
2. 根据 `nft_minted` 获取该 token 的 `usdt_paid`（铸造价）
3. 按胜率与盈亏区间生成随机 reward（int256）
4. 调用 `Nft.setRewardOnce(tokenId, reward)`
5. 更新 `nft_opened.check_status=1`、`reward`，以及 `nft_minted.open_status=1`、`opened_at`

## 运行

```bash
cd requestEth
go run ./cmd/blind_settle_worker -conf=./configs
```

## 环境变量

| 变量 | 说明 |
|------|------|
| `ETH_RPC_URL` | BSC RPC（或从 config.yaml 读取） |
| `NFT_ADDR` | NFT 合约地址 |
| `BLIND_SETTLE_PRIVATE_KEY` | 具备 ADMIN_ROLE 的私钥（hex，无 0x 前缀） |
| `DATABASE_SOURCE` | MySQL 连接串（或从 config.yaml 读取） |

## 随机规则（config.yaml `blind_settle`）

| 配置 | 默认 | 说明 |
|------|------|------|
| `win_rate_bps` | 5000 | 胜率 50% (5000/10000) |
| `profit_min_bps` | 5000 | 盈利下限：usdtPaid 的 50% |
| `profit_max_bps` | 20000 | 盈利上限：usdtPaid 的 200% |
| `loss_min_bps` | 5000 | 亏损下限：usdtPaid 的 50% |
| `loss_max_bps` | 10000 | 亏损上限：usdtPaid 的 100% |
| `chain_id` | 56 | 链 ID（56=BSC 主网，97=测试网） |

## 前置条件

- `nft_opened` 数据需由主服务的 `GetBoxOpenEvent`（或等价逻辑）同步链上 Opened 事件而来
- 私钥对应地址须在 NFT 合约中拥有 `ADMIN_ROLE`
