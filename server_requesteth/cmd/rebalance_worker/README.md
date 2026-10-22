# 调仓脚本

定时将热钱包（accountA）的 USDT 余额按比例转到交易所充值地址。

## 流程

1. 每 N 分钟轮询一次
2. 查询热钱包 USDT 余额
3. 若余额 ≥ 阈值，转出 95% 到交易所地址（5% 留作 staking 等）

## 运行

```bash
cd requestEth
go run ./cmd/rebalance_worker -conf=./configs
```

## 环境变量

| 变量 | 说明 |
|------|------|
| `ETH_RPC_URL` | BSC RPC |
| `USDT_ADDR` | USDT 合约（默认 BSC 主网 0x55d398...） |
| `EXCHANGE_ADDR` | 交易所充值地址 |
| `REBALANCE_PRIVATE_KEY` | 热钱包私钥（持有 USDT 的 accountA，hex 无 0x） |

## 配置（config.yaml `rebalance`）

| 配置 | 默认 | 说明 |
|------|------|------|
| `rebalance_bps` | 9500 | 转出比例 95% |
| `min_amount_usd` | 100 | 余额 ≥ 100 USDT 才执行 |
| `chain_id` | 56 | BSC 主网 |
| `interval_min` | 60 | 检查间隔（分钟） |

## 注意

- 私钥对应地址须为 NFT 铸造分账收款地址（accountA），即持有 USDT 的热钱包
- 确保热钱包有足够 BNB 支付 gas
