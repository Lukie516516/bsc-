# 前端 API 调用说明

后端 base URL 示例：`https://your-api-domain.com`（以实际部署为准）

---

## 一、查看数据（GET）

### 1. 盲盒相关

| 接口 | 方法 | 参数 | 查看内容 |
|------|------|------|----------|
| `/api/get_box_all` | GET | 无 | 盲盒总数：newLength（未开）、openLength（已开） |
| `/api/get_box_new` | GET | start, end（时间戳） | 未开盲盒列表：tokenId、价格、铸造时间等 |
| `/api/get_box_open` | GET | start, end | 已开盲盒列表：tokenId、reward、开盒时间等 |
| `/api/get_address_box` | GET | address, page, num, numTwo, numThree, numFour, openStatus, start, end | **某用户盲盒**：未上架/上架中/全部，按档位、时间排序 |
| `/api/get_market_box_list` | GET | page, num, numTwo, numThree | **二级市场在售盲盒**：可购买列表 |
| `/api/get_sell_box_list` | GET | address, page | 某用户已售盲盒记录 |
| `/api/get_buy_box_list` | GET | address, page | 某用户已购盲盒记录 |

### 2. 全局统计

| 接口 | 方法 | 参数 | 查看内容 |
|------|------|------|----------|
| `/api/get_all_info` | GET | start, end（时间戳） | **大盘数据**：用户数、今日铸造、总铸造、销售总量/金额、盲盒总盈利等 |
| `/api/get_opened_box_profit_stats` | GET | start, end, page, page_size | **开盒利润统计**：盈利/亏损明细、利润总和、盈利/亏损数量 |
| `/api/get_fee_income_stats` | GET | start, end | **手续费收入统计**：USDT/HRX 手续费总和、交易笔数 |

### 3. 分红 / 收益记录

| 接口 | 方法 | 参数 | 查看内容 |
|------|------|------|----------|
| `/api/get_reward_list` | GET | address, page, reason | 某用户分红记录：盲盒收益、L1/L2/Top/项目方收益 |
| `/api/referral_pending` | GET | page, page_size, paid | 铸造分佣待发/已发列表（后台用） |

### 4. 链上 / 其他

| 接口 | 方法 | 参数 | 查看内容 |
|------|------|------|----------|
| `/api/usdt_balance` | GET | address, token | 某地址 USDT 余额 |
| `/api/eth_balance` | GET | address | 某地址 BNB 余额 |
| `/api/get_reserves` | GET | pair | 池子储备量 |
| `/api/verify_sig` | GET | sign, content | 验证签名，返回地址 |

---

## 二、修改 / 操作（POST）

### 1. 结算与打款（需后端私钥，慎重）

| 接口 | 方法 | 请求体 | 作用 |
|------|------|--------|------|
| `/api/set_reward` | POST | `{"sendBody":{"tokenId":"123","reward":"150"}}` | 为盲盒设置 reward（`BlindBoxNFT.setRewardOnce`），发放本金+通知分红 |
| `/api/set_reward_two` | POST | 同上 | 兼容保留接口，当前内部同样转发到 `BlindBoxNFT.setRewardOnce` |
| `/api/referral_payout` | POST | `{"sendBody":{"ids":[1,2,3]}}` | 铸造分佣批量打款 |

> 通常由 `blind_settle_worker` 定时调用，或管理后台「手动触发」按钮调用，前端一般不直接调用。

### 2. 交易 / 链上操作（示例，部分可能需配合钱包）

| 接口 | 方法 | 请求体 | 作用 |
|------|------|--------|------|
| `/api/transaction` | POST | `sendBody.privateKey`, `sendBody.toAddr`, `sendBody.token`, `sendBody.amount` | 代币转账 |
| `/api/transaction_eth` | POST | `sendBody.privateKey`, `sendBody.toAddr`, `sendBody.amount` | BNB 转账 |
| `/api/add_liquidity` | POST | `sendBody.orderId`, `sendBody.user`, `sendBody.amountUsdt`, `sendBody.amountAicat` | 添加流动性 |
| `/api/remove_liquidity` | POST | `sendBody.orderId`, `sendBody.user`, `sendBody.liquidityAmount` | 移除流动性 |
| `/api/buy` | POST | `sendBody.orderId`, `sendBody.usdtAmount`, `sendBody.contract` | 一级市场买 HRX |

> 用户铸造、开盒、上架、购买等，**一般由前端连接钱包直接调合约**，不一定走这些 API。

---

## 三、前端典型调用示例

### 1. 用户「我的盲盒」页

```javascript
// 获取某地址的盲盒（未上架、上架中、全部）
fetch(`${API_BASE}/api/get_address_box?address=${userAddress}&page=1&num=2&numTwo=0&numThree=0&numFour=0&openStatus=0&start=0&end=9999999999`)
  .then(r => r.json())
  .then(data => {
    // data.list: 盲盒列表
    // data.count: 总数
  });
```

### 2. 二级市场「在售盲盒」页

```javascript
fetch(`${API_BASE}/api/get_market_box_list?page=1&num=0&numTwo=0&numThree=0`)
  .then(r => r.json())
  .then(data => {
    // data.list: 在售盲盒
  });
```

### 3. 大盘 / 统计页

```javascript
// 全局数据
fetch(`${API_BASE}/api/get_all_info?start=${startTs}&end=${endTs}`)
  .then(r => r.json())
  .then(data => {
    // data.userCount, todayMintedCount, mintedSum, buySum, openReward...
  });

// 开盒利润（按时间段）
fetch(`${API_BASE}/api/get_opened_box_profit_stats?start=${start}&end=${end}&page=1&page_size=20`)
  .then(r => r.json())
  .then(data => {
    // data.total_profit, profit_count, loss_count, list
  });

// 手续费收入
fetch(`${API_BASE}/api/get_fee_income_stats?start=${start}&end=${end}`)
  .then(r => r.json())
  .then(data => {
    // data.fee_usdt_sum, fee_b_sum, count
  });
```

### 4. 用户分红记录

```javascript
fetch(`${API_BASE}/api/get_reward_list?address=${userAddress}&page=1&reason=0`)
  .then(r => r.json())
  .then(data => {
    // data.list: reason 1盲盒 2L1 3L2 4Top 5项目方
  });
```

### 5. 管理员触发结算（可选）

```javascript
// 为某个 tokenId 设置 reward（通常由 worker 自动执行）
fetch(`${API_BASE}/api/set_reward`, {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({
    sendBody: {
      tokenId: '123',
      reward: '150'   // 开盒收益，正数
    }
  })
});
```

---

## 四、数据流与职责划分

| 数据类型 | 来源 | 前端做什么 |
|----------|------|------------|
| 盲盒列表、市场、用户持有 | 后端 API | 调 GET 接口展示 |
| 统计（利润、手续费、大盘） | 后端 API | 调 GET 接口做图表/表格 |
| 分红记录、铸造分佣 | 后端 API | 调 GET 接口展示 |
| 铸造、开盒、上架、购买、质押 | **链上合约** | 前端用 web3/ethers 调合约 + 钱包签名 |
| 结算 setReward | 后端 API 或 worker | 管理后台可提供「手动触发」按钮，调用 POST `/api/set_reward` |

---

## 五、常见参数说明

- **address**：用户钱包地址  
- **page**：页码（从 1 起）  
- **start / end**：Unix 时间戳（秒）  
- **num**（get_address_box）：0=未上架，1=上架中，2=全部  
- **numTwo**：0=全部档位，1=20u，2=50u，3=100u  
- **numThree / numFour**：0=倒序，1=正序  
- **openStatus**：0=全部，1=未开，2=已开  
- 所有 `POST` 接口优先按 `sendBody` 包装字段传参，和 proto 定义保持一致
