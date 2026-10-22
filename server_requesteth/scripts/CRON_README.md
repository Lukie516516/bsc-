# 链上事件同步 - 宝塔 Cron 配置

## 不需要在宝塔新建数据库

当前使用 **Docker MySQL**，库名 `game` 已由 docker-compose 自动创建。  
API 启动时会自动建表（AutoMigrate），**无需在宝塔里新建数据库**。

---

## 宝塔添加计划任务

1. 登录宝塔 → **计划任务**
2. 添加任务：
   - **任务类型**：Shell 脚本
   - **任务名称**：requestEth 链上事件同步
   - **执行周期**：N 分钟（建议 1 分钟）
   - **脚本内容**：

```bash
/www/wwwroot/NewrequestEth/requestEth/scripts/sync-events.sh
```

3. 可选：勾选「日志」保存输出，或手动指定日志路径：

```bash
/www/wwwroot/NewrequestEth/requestEth/scripts/sync-events.sh >> /www/wwwroot/NewrequestEth/requestEth/log/sync-events.log 2>&1
```

4. 保存后，可先「执行」一次测试

---

## 脚本说明

`sync-events.sh` 会依次调用以下 API，将链上事件写入数据库：

| API | 写入表 | 用途 |
|-----|--------|------|
| get_box_mint_event | nft_minted | referral_worker 铸造分佣（失败时自动调用 sync_mint_from_contract 回退） |
| get_box_open_event | nft_opened | blind_settle_worker 开盒结算 |
| get_box_list_event | nft_market_listed | 市场挂单 |
| get_box_unlist_event | nft_market_unlisted | 取消挂单 |
| get_box_buy_event | nft_market_purchase | 市场购买 |
| get_box_transfer_event | nft_transfer | NFT 转账 |
| **update_box** | **合并到 nft_minted** | **卖出/转让后更新 `to_addr`，否则仓库列表仍显示旧数据** |

---

## 修改 API 地址

若 API 端口不是 8003，可设置环境变量：

```bash
API_BASE=http://127.0.0.1:8003 /www/wwwroot/NewrequestEth/requestEth/scripts/sync-events.sh
```
