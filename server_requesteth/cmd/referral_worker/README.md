# 铸造分佣 Worker

监听 **铸造盲盒** 事件（nft_minted），计算 B/D 分佣待发记录，写入 `referral_pending` 表。**二级市场购买盲盒不分佣**。比例可由 config 或环境变量配置。

## 运行

```bash
# 必须设置
export NFT_ADDR=0x...   # NFT 合约地址
export STAKE_ADDR=0x... # 分红(Stake)合约地址

# 可选，默认从 configs/config.yaml 读取
export ETH_RPC_URL=https://...  # BSC RPC
export DATABASE_SOURCE=root:xxx@tcp(127.0.0.1:3306)/game?parseTime=true

cd cmd/referral_worker
go run . -conf=../../configs
```

## 依赖

- 数据库已有 `nft_minted` 表（由链上 Minted 事件监听写入）

## 表结构迁移（若已有 referral_pending）

若表来自旧版（基于 nft_market_purchase），需执行：

```sql
-- 新增 minted_id 列
ALTER TABLE referral_pending ADD COLUMN minted_id BIGINT UNSIGNED DEFAULT 0;
ALTER TABLE referral_pending ADD INDEX idx_minted (minted_id);
-- 若 purchase_id 有 unique 约束导致新记录(minted_id>0, purchase_id=0)冲突，需改为普通索引：
-- ALTER TABLE referral_pending DROP INDEX idx_purchase;
-- ALTER TABLE referral_pending ADD INDEX idx_purchase (purchase_id);
```

## 输出

- 表 `referral_pending`：每笔铸造一条记录，比例可配（默认 B 0.3%、D 0.2%）
- **配置**：config.yaml `referral` 或环境变量 `REFERRAL_B_BPS`、`REFERRAL_D_BPS`、`REFERRAL_HOLD_MIN_USD`
- 打款：由项目方根据 `ListPending` 结果，批量 USDT 打给 B（推广者）和 D（团队长）
