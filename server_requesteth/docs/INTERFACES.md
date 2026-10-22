# 链上与中心化接口说明

## 1. 盲盒结算（开盒 → 平仓 → setReward）

### 链上接口

- **openBoxPayUsdt(tokenId)** / **openBoxPayHRX(tokenId, maxFeeB)**：用户开盒，合约标记 `open[tokenId]=true`，发出 `Opened(owner, tokenId, timestamp, feeU, feeB)`
- **setRewardOnce(tokenId, reward)**：管理员写入盈亏，仅 NFT 合约 `ADMIN_ROLE` 可调

### 中心化职责

1. 监听 `Opened` 事件
2. 调用交易所 API 平仓对应 NFT 仓位
3. 获取结算盈亏
4. 调用 `setRewardOnce(tokenId, profit)` 写回链上

### 结算随机规则（中心化实现）

- 在中心化服务按配置生成随机盈亏：胜率、盈利区间（如 5%–8%）、亏损区间（如 5%–8%）
- 生成结果后调用 `setRewardOnce(tokenId, profit)` 写入链上
- 链上仅做结果校验与状态更新，不参与随机逻辑

---

## 2. 量化相关（由外部系统负责）

### 新铸造盲盒启动量化

- 调仓后，新铸造的盲盒需在交易所启动量化
- 由中心化脚本/服务调用交易所 API 完成，链上无接口
- 需监听 `Minted` 事件，获取新盲盒信息

### 调仓

- 链上：铸造分账 95% accountA，5% staking
- 链上预留接口：`setQuantReceiver(receiver)`、`transferHeldUsdtToQuant(amount, bizNo)`、`transferHeldERC20ToQuant(token, amount, bizNo)`
- 上述预留接口只处理 **已经在 NFT 合约里的资金**，不改变当前 mint 时 95% 直达 `accountA`、5% 直达 `staking` 的现有主路径
- 中心化：定时脚本将 accountA 余额的 95% 转到交易所充值地址
- 具体调仓与量化策略由量化团队实现

### 结算/收款优化

- 二级市场普通钱包直卖：原本就是点对点直接到账
- Vault 批量挂卖：现已支持通过 `payoutRecipient()` 直接把成交款打到用户钱包，不再必须先回到 Vault 再手动提取
- `Stake.sol` 新增：
  - `setProjectAutoSettlement(enabled)`：项目方沉淀份额可改为自动结算到项目方地址
  - `settleProjectPending()`：手动把已沉淀到 `pendingUsdt[projectAddress]` 的金额一次性结算出去

### 管理权限与团队长资格

- `Nft.sol`、`Stake.sol` 均新增 `SUPER_ADMIN_ROLE`
- 超管与日常业务管理员可分离：紧急提币、量化转出等敏感操作只给超管
- `Stake.sol` 的团队长资格规则调整为：
  - 成为 `top` 不等于自动拥有团队长收益资格
  - `minStakeA` 为构造期强制传入的非零门槛，且运行期不可改回 `0`
  - `setTopManualEligible(top, enabled)` 仅作为未达质押门槛时的人工特批兜底
- `Stake.sol` 新增 `getTopEligibilityStatus(top)`：可直接查看启用状态、人工授权状态、当前质押、门槛值与最终资格结果

### 预留占位接口

- `buyInsurance(tokenId, amountB)`：保险购买占位接口，当前仅保留 ABI，调用会回退
- `buyProp(propType, amountB)`：道具购买占位接口，当前仅保留 ABI，调用会回退

### 开盒次数购买（HRX）

- `setOpenCountCostConfigB(firstCostB, stepCostB, enabled)`：后台可直接配置“用 HRX 购买开盒次数”的首项与步长
- `quoteOpenCountCostInB(user, count)`：查询用户购买若干开盒次数所需的 HRX 数量
- 当 `enabled=false` 时，仍沿用旧逻辑：按 USDT 成本经池子报价换算成 HRX

---

## 3. 铸造分佣（0.2% + 0.3%）

### 实现方式

- 在中心化服务器完成：监听**铸造**（nft_minted），计算 B/D 分配，写入 `referral_pending` 表；二级市场购买不分佣
- 项目方定期批量打款给 B（推广者）和 D（团队长）
- Worker: `cmd/referral_worker`

### 分配规则

- **团队长**：D = topOf(minter)，须满足 `isTopEligible`
- `isTopEligible` 的产品语义为：
  - 团队长默认不因“成为 top”自动获得收益资格
  - 只有满足质押门槛，或被管理员人工授权，才算合格团队长
- **推广者**：B = parentOf(minter)，须 holdingValueOf(B) ≥ 300 USDT

| 条件 | 结果 |
|------|------|
| 无上级（B 为空） | 0.5% 归项目方 |
| D 非团队长 | 0.5% 归项目方 |
| B 未达 300u | 0.5% 归项目方 |
| 正常 | B 0.3%，D 0.2% |

---

## 4. 环境变量（API 需配置）

### 私钥（敏感）
| 变量 | 用途 |
|------|------|
| `REFERRAL_PAYOUT_PRIVATE_KEY` | 铸造分佣打款（ReferralPayout） |
| `REFERRAL_USDT_ADDR` | 铸造分佣打款使用的 USDT 合约地址 |
| `AICAT_ADMIN_PRIVATE_KEY` | AICAT 相关接口（AddLiquidity、RemoveLiquidity、BuyAICAT、WithdrawAICAT、SendAICAT） |
| `SET_REWARD_ADMIN_PRIVATE_KEY` | `SetReward` / `SetRewardTwo` 共用的后台私钥，当前两者最终都调用 `BlindBoxNFT.setRewardOnce` |

### 链上配置（合约地址、RPC）
| 变量 | 用途 |
|------|------|
| `ETH_BSC_RPC_URLS` | BSC RPC 列表，逗号分隔（交易、查询用） |
| `ETH_BSC_RPC_URLS_POLLING` | 事件轮询用 RPC（可配置付费节点） |
| `ETH_NFT_CONTRACT` | BlindBoxNFT 合约地址 |
| `ETH_STAKE_CONTRACT` | UmbrellaStakeDividend 合约地址 |
| `ETH_USDT_CONTRACT` | USDT 合约地址 |
| `ETH_DIVIDEND_CONTRACT` | 分红合约地址 |
| `ETH_SWAP_CONTRACT` | SwapHrxUsdt 合约地址 |
| `ETH_PRIMARY_MARKET_CONTRACT` | BondingCurvePrimaryMarket 合约地址 |
| `ETH_REGISTER_CONTRACT` | 触发 Registered 的合约地址 |
| `ETH_VAULT_FACTORY_CONTRACT` | BlindBoxVaultFactory 合约地址 |
| `ETH_AICAT_ADMIN_CONTRACT` | AICAT Admin 合约地址 |
| `ETH_AICAT_ADMIN_NEW_CONTRACT` | AICAT AdminNew 合约地址 |

### 当前正式版应同步的核心地址

```env
ETH_NFT_CONTRACT=0x169c4E946909D4E115b724957e5D1c74d0971f70
ETH_STAKE_CONTRACT=0x907F84e4127049bB65C82F50CE29c1FF67f344D5
ETH_SWAP_CONTRACT=0x3c9Ea8d476eA06108283457B3Bb0dfbfD6cf1A1b
ETH_PRIMARY_MARKET_CONTRACT=0xdA5A5BC6b61B2eAB091A1b7E4d84C7eb537BDC0e
ETH_VAULT_FACTORY_CONTRACT=0x7AD5C447580A79e072bFD72E373f7b25dDE0F061
ETH_USDT_CONTRACT=0x55d398326f99059fF775485246999027B3197955
```

### 前端对接提醒

- 挂卖金库地址不是固定环境变量，要先通过 `BlindBoxVaultFactory.vaultOf(user)` 查询
- 当前版本持仓已无限制，`quoteStorageSlotCost` / `quoteSlotCostInB` 只是兼容保留，前端可直接视为 `0`
- 团队长资格展示建议优先调用 `getTopEligibilityStatus(top)`，不要只根据关系链推断
