# 盲盒项目 API 调用文档

**Base URL**：`http://47.245.108.129:8003`  
**Content-Type**：`application/json`  
**跨域**：已配置 CORS，支持 `*` 跨域

---

## 一、健康检查

| 接口 | 方法 | 说明 |
|------|------|------|
| `/health` | GET | 健康检查，返回 `{"status":"ok"}` |

**示例**：
```bash
curl http://47.245.108.129:8003/health
```

---

## 二、盲盒相关

### 2.1 盲盒总数

| 接口 | 方法 | 参数 | 说明 |
|------|------|------|------|
| `/api/get_box_all` | GET | 无 | 获取盲盒总数（未开 / 已开） |

**响应示例**：
```json
{
  "newLength": "100",
  "openLength": "50"
}
```

---

### 2.2 未开盲盒列表

| 接口 | 方法 | 参数 | 说明 |
|------|------|------|------|
| `/api/get_box_new` | GET | start, end | 按 tokenId 范围获取未开盲盒（从链上读取） |

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| start | string | 是 | tokenId 起始（十进制） |
| end | string | 是 | tokenId 结束（十进制） |

**响应示例**：
```json
{
  "list": [
    {
      "tokenId": "1",
      "price": "20000000000000000000",
      "openedAt": 0,
      "mintAt": 1700000000,
      "reward": "0",
      "rewardAt": 0
    }
  ]
}
```

---

### 2.3 已开盲盒列表

| 接口 | 方法 | 参数 | 说明 |
|------|------|------|------|
| `/api/get_box_open` | GET | start, end | 按 tokenId 范围获取已开盲盒（从链上读取） |

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| start | string | 是 | tokenId 起始 |
| end | string | 是 | tokenId 结束 |

**响应示例**：
```json
{
  "list": [
    {
      "tokenId": "1",
      "price": "20000000000000000000",
      "openedAt": 1700000100,
      "mintAt": 1700000000,
      "reward": "25000000000000000000",
      "rewardAt": 1700000200
    }
  ]
}
```

---

### 2.4 某用户盲盒列表（我的盲盒）

| 接口 | 方法 | 参数 | 说明 |
|------|------|------|------|
| `/api/get_address_box` | GET | 见下表 | 某地址持有的盲盒，支持筛选、分页、排序 |

| 参数 | 类型 | 必填 | 默认 | 说明 |
|------|------|------|------|------|
| address | string | 是 | - | 用户钱包地址 |
| page | uint64 | 是 | - | 页码（从 1 起） |
| num | uint64 | 否 | - | 0=未上架，1=上架中，2=全部 |
| numTwo | uint64 | 否 | - | 0=全部档位，1=20u，2=50u，3=100u |
| numThree | uint64 | 否 | - | 0=倒序，1=正序 |
| numFour | uint64 | 否 | - | 0=默认，1=倒序，2=正序 |
| openStatus | uint64 | 否 | - | 0=全部，1=未开，2=已开 |
| start | uint64 | 否 | - | 时间戳起始（秒） |
| end | uint64 | 否 | - | 时间戳结束（秒） |

**响应示例**：
```json
{
  "count": 10,
  "list": [
    {
      "tokenId": 1,
      "priceUsdt": 20,
      "tier": 1,
      "priceUsdtNow": 22,
      "blockTime": 1700000000,
      "listTime": 1700000100,
      "openStatus": 0,
      "openTime": 0,
      "minter": "0x..."
    }
  ]
}
```

---

### 2.5 二级市场在售盲盒

| 接口 | 方法 | 参数 | 说明 |
|------|------|------|------|
| `/api/get_market_box_list` | GET | 见下表 | 二级市场在售盲盒列表 |

| 参数 | 类型 | 必填 | 默认 | 说明 |
|------|------|------|------|------|
| page | uint64 | 是 | - | 页码 |
| num | uint64 | 否 | - | 时间排序：0=倒序，1=正序 |
| numTwo | uint64 | 否 | - | 0=全部，1=20u，2=50u，3=100u |
| numThree | uint64 | 否 | - | 价格排序：0=倒序，1=正序 |

**响应示例**：
```json
{
  "count": 5,
  "list": [
    {
      "tokenId": 2,
      "priceUsdt": 50,
      "tier": 2,
      "priceUsdtNow": 55,
      "blockTime": 1700000000,
      "listTime": 1700000200,
      "minter": "0x..."
    }
  ]
}
```

---

### 2.6 用户购买盲盒记录

| 接口 | 方法 | 参数 | 说明 |
|------|------|------|------|
| `/api/get_buy_box_list` | GET | address, page | 某用户二级市场购买盲盒记录 |

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| address | string | 是 | 用户钱包地址 |
| page | uint64 | 是 | 页码 |

**响应示例**：
```json
{
  "count": 20,
  "list": [
    {
      "user": "0x...",
      "tokenId": 3,
      "priceUsdt": 100,
      "blockTime": 1700000300,
      "feeB": 1.5,
      "feeUsdt": 2
    }
  ]
}
```

---

### 2.7 用户出售盲盒记录

| 接口 | 方法 | 参数 | 说明 |
|------|------|------|------|
| `/api/get_sell_box_list` | GET | address, page | 某用户二级市场出售盲盒记录 |

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| address | string | 是 | 用户钱包地址 |
| page | uint64 | 是 | 页码 |

**响应示例**：
```json
{
  "count": 10,
  "list": [
    {
      "tokenId": 4,
      "priceUsdt": 50,
      "blockTime": 1700000400,
      "feeB": 0.5,
      "feeUsdt": 1
    }
  ]
}
```

---

## 三、全局统计

### 3.1 大盘数据

| 接口 | 方法 | 参数 | 说明 |
|------|------|------|------|
| `/api/get_all_info` | GET | start, end | 大盘汇总数据 |

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| start | uint64 | 是 | 开始时间戳（秒） |
| end | uint64 | 是 | 结束时间戳（秒） |

**响应示例**：
```json
{
  "userCount": 1000,
  "todayUserCount": 50,
  "todayMintedCount": 20,
  "todayMintedSum": "1000",
  "mintedCount": 500,
  "mintedSum": "25000",
  "mintedCountNoOpen": 100,
  "mintedSumNoOpen": "5000",
  "mintedCountOne": 200,
  "mintedSumOne": "4000",
  "mintedCountTwo": 200,
  "mintedSumTwo": "10000",
  "mintedCountThree": 100,
  "mintedSumThree": "10000",
  "buyCount": 300,
  "buySum": "15000",
  "todayBuyCount": 10,
  "todayBuySum": "500",
  "openReward": "12000"
}
```

---

## 四、分红 / 收益

### 4.1 用户分红记录

| 接口 | 方法 | 参数 | 说明 |
|------|------|------|------|
| `/api/get_reward_list` | GET | address, page, reason | 某用户分红记录 |

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| address | string | 是 | 用户钱包地址 |
| page | uint64 | 是 | 页码（每页 20 条） |
| reason | uint64 | 否 | 1=盲盒收益，2=L1，3=L2，4=Top，5=项目方；0=全部 |

**响应示例**：
```json
{
  "count": 15,
  "list": [
    {
      "user": "0x...",
      "reason": 1,
      "amount": 25.5,
      "blockTime": 1700000500
    }
  ]
}
```

---

## 五、HRX 交易 / K 线

### 5.1 HRX 兑换记录（K 线数据）

| 接口 | 方法 | 参数 | 说明 |
|------|------|------|------|
| `/api/get_exchange_list` | GET | start, end | HRX/USDT 兑换记录，用于 K 线 |

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| start | uint64 | 是 | 开始时间戳（秒） |
| end | uint64 | 是 | 结束时间戳（秒） |

**响应示例**：
```json
{
  "list": [
    {
      "blockTime": 1700000000,
      "price": 0.5
    }
  ]
}
```

---

### 5.2 HRX 一级市场购买记录

| 接口 | 方法 | 参数 | 说明 |
|------|------|------|------|
| `/api/get_buy_list` | GET | start, end | 一级市场 HRX 购买记录 |

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| start | uint64 | 是 | 开始时间戳 |
| end | uint64 | 是 | 结束时间戳 |

**响应示例**：
```json
{
  "list": [
    {
      "blockTime": 1700000000,
      "price": 0.5,
      "usdtUse": 100,
      "ausdGrossOut": 200
    }
  ]
}
```

> ⚠️ **注意**：`get_buy_event`、`get_sell_event` 未在 sync-events 中定时调用，`primary_buy` / `primary_sell` 表可能为空，`get_buy_list` 可能返回空数据。

---

## 六、链上 / 余额

### 6.1 USDT 余额

| 接口 | 方法 | 参数 | 说明 |
|------|------|------|------|
| `/api/usdt_balance` | GET | address, token | 某地址代币余额 |

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| address | string | 是 | 钱包地址 |
| token | string | 否 | 代币合约地址，不传则用默认 USDT |

**响应示例**：
```json
{
  "balance": "1000.5",
  "err": ""
}
```

---

### 6.2 BNB 余额

| 接口 | 方法 | 参数 | 说明 |
|------|------|------|------|
| `/api/eth_balance` | GET | address | 某地址 BNB 余额 |

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| address | string | 是 | 钱包地址 |

**响应示例**：
```json
{
  "balance": "0.5",
  "err": ""
}
```

---

### 6.3 池子储备量

| 接口 | 方法 | 参数 | 说明 |
|------|------|------|------|
| `/api/get_reserves` | GET | pair | 交易对储备量 |

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| pair | string | 是 | 交易对合约地址 |

**响应示例**：
```json
{
  "reservesOne": "1000000",
  "reservesTwo": "500000"
}
```

---

### 6.4 验证签名

| 接口 | 方法 | 参数 | 说明 |
|------|------|------|------|
| `/api/verify_sig` | GET | sign, content | 验证签名，返回签名地址 |

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| sign | string | 是 | 签名 |
| content | string | 是 | 被签名的内容 |

**响应示例**：
```json
{
  "res": true,
  "address": "0x..."
}
```

---

## 七、其他（流动性、一级市场等）

### 7.1 用户 LP

| 接口 | 方法 | 参数 | 说明 |
|------|------|------|------|
| `/api/get_user_lp` | GET | address | 用户 LP 数量 |

---

### 7.2 日手续费

| 接口 | 方法 | 参数 | 说明 |
|------|------|------|------|
| `/api/get_daily_fee` | GET | ts | 某日手续费 |

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| ts | string | 否 | 当天 0 点时间戳（十进制秒）；不传或格式非法时返回安全默认值 |

---

### 7.3 按 orderId 查询

| 接口 | 方法 | 参数 | 说明 |
|------|------|------|------|
| `/api/get_lp` | GET | orderId | 按 orderId 查 LP |
| `/api/get_buy` | GET | orderId, contract | 按 orderId 查购买 |
| `/api/get_buy_aicat_by_order_id` | GET | orderId | 按 orderId 查 AICAT 购买 |

---

## 八、POST 接口（操作类）

### 8.1 结算（开盒 reward）

| 接口 | 方法 | 请求体 | 说明 |
|------|------|--------|------|
| `/api/set_reward` | POST | JSON | 为盲盒设置 reward，发放本金+通知分红 |
| `/api/set_reward_two` | POST | JSON | 兼容保留接口，当前内部同样转发到 `BlindBoxNFT.setRewardOnce` |

**请求体**：
```json
{
  "sendBody": {
    "tokenId": "123",
    "reward": "150"
  }
}
```

> `reward` 为字符串形式的金额值，按当前业务约定传 USDT 口径的 18 位精度字符串。  
> 通常由 `blind_settle_worker` 或管理后台调用，前端一般不直接调用。

---

### 8.2 代币 / BNB 转账

| 接口 | 方法 | 请求体 | 说明 |
|------|------|--------|------|
| `/api/transaction` | POST | JSON | 代币转账 |
| `/api/transaction_eth` | POST | JSON | BNB 转账 |

**transaction 请求体**：
```json
{
  "sendBody": {
    "privateKey": "...",
    "toAddr": "0x...",
    "token": "0x...",
    "amount": "100"
  }
}
```

**transaction_eth 请求体**：
```json
{
  "sendBody": {
    "privateKey": "...",
    "toAddr": "0x...",
    "amount": "0.1"
  }
}
```

---

### 8.3 流动性 / 一级市场

| 接口 | 方法 | 请求体 | 说明 |
|------|------|--------|------|
| `/api/add_liquidity` | POST | JSON | 添加流动性 |
| `/api/remove_liquidity` | POST | JSON | 移除流动性 |
| `/api/buy` | POST | JSON | 一级市场买 HRX |
| `/api/withdraw` | POST | JSON | 提现 |

**add_liquidity 请求体**：
```json
{
  "sendBody": {
    "orderId": "xxx",
    "user": "0x...",
    "amountUsdt": "100",
    "amountAicat": "200"
  }
}
```

**remove_liquidity 请求体**：
```json
{
  "sendBody": {
    "orderId": "xxx",
    "user": "0x...",
    "liquidityAmount": "50"
  }
}
```

**buy 请求体**：
```json
{
  "sendBody": {
    "orderId": "xxx",
    "usdtAmount": "100",
    "contract": "0x..."
  }
}
```

---

## 九、事件同步接口（内部用）

以下接口由 `sync-events.sh` 定时调用，用于同步链上事件到数据库，前端一般**不直接调用**：

| 接口 | 说明 |
|------|------|
| `/api/get_box_mint_event` | 铸造事件 → nft_minted |
| `/api/get_box_open_event` | 开盒事件 → nft_opened |
| `/api/get_box_list_event` | 挂单事件 |
| `/api/get_box_unlist_event` | 取消挂单事件 |
| `/api/get_box_buy_event` | 市场购买事件 → nft_market_purchase |
| `/api/get_box_transfer_event` | 转账事件 |
| `/api/get_user_r_event` | 用户注册事件 |
| `/api/get_exchange_event` | HRX 兑换事件 → swap_trade |

---

## 十、补充接口

以下接口当前已经可用，但更偏后台/运营使用，前端一般不会直接暴露到普通用户页面：

| 接口 | 方法 | 说明 |
|------|------|------|
| `/api/get_opened_box_profit_stats` | GET | 开盒利润统计 |
| `/api/get_fee_income_stats` | GET | 二级市场手续费统计 |
| `/api/referral_pending` | GET | 铸造分佣待发列表 |
| `/api/referral_payout` | POST | 铸造分佣批量打款 |

其中 `POST` 类接口如果由前端调试调用，仍需按 proto 结构传 `sendBody` 包装字段。

---

## 十一、前端典型调用示例

### 我的盲盒

```javascript
const API_BASE = 'http://47.245.108.129:8003';

// 某用户全部盲盒，第 1 页
fetch(`${API_BASE}/api/get_address_box?address=${userAddress}&page=1&num=2&numTwo=0&numThree=0&numFour=0&openStatus=0&start=0&end=9999999999`)
  .then(r => r.json())
  .then(data => {
    console.log(data.list, data.count);
  });
```

### 二级市场在售

```javascript
fetch(`${API_BASE}/api/get_market_box_list?page=1&num=0&numTwo=0&numThree=0`)
  .then(r => r.json())
  .then(data => {
    console.log(data.list, data.count);
  });
```

### 大盘数据

```javascript
const start = Math.floor(Date.now() / 1000) - 86400 * 7; // 7 天前
const end = Math.floor(Date.now() / 1000);
fetch(`${API_BASE}/api/get_all_info?start=${start}&end=${end}`)
  .then(r => r.json())
  .then(data => {
    console.log(data.userCount, data.mintedSum, data.openReward);
  });
```

### 用户分红记录

```javascript
fetch(`${API_BASE}/api/get_reward_list?address=${userAddress}&page=1&reason=0`)
  .then(r => r.json())
  .then(data => {
    console.log(data.list, data.count);
  });
```

### K 线（HRX 兑换）

```javascript
fetch(`${API_BASE}/api/get_exchange_list?start=${start}&end=${end}`)
  .then(r => r.json())
  .then(data => {
    console.log(data.list);
  });
```

---

## 十二、错误与限流

- **限流**：使用 BBR 算法，高并发时可能返回 429
- **错误**：业务错误时返回 4xx/5xx，响应体可能包含 `err` 字段
- **CORS**：已允许 `*` 跨域，支持 GET/POST/PUT/HEAD/OPTIONS

---

*文档版本：2025-03-16*
