# NFT 合约与 Go 后端同步说明

## 一、当前状态

### 已完成
1. **ABI 已更新**：从 `box/artifacts/contracts/Nft.sol/BlindBoxNFT.json` 提取的 ABI 已写入 `requestEth/abi/box.abi` 和 `requestEth/abi/nft.abi`。
2. **Makefile 目标**：已添加 `make nft-abi` 用于重新生成 Go 绑定。

### 待执行（需本机安装 Go）
在 `requestEth` 目录执行：
```bash
make nft-abi
```
将安装 abigen 并生成 `internal/service/nft.go`。

---

## 二、合约与 Go 的映射变更

| 合约（当前 Nft.sol） | Go（旧绑定） | 说明 |
|---------------------|-------------|------|
| `buyStorageSlots` / `buyStorageSlotsPayHRX` | `BuyOpenSlots` | 槽位与开盒次数拆分 |
| `buyOpenCount` / `buyOpenCountPayHRX` | `BuyOpenSlots` | 同上 |
| `extraStorageSlots` | `ExtraSlots` | 槽位映射 |
| `extraOpenCount` | `ExtraSlots` | 开盒次数映射 |
| `openBoxPayUsdt` | `OpenBox` (payable) | 开盒用 USDT |
| `openBoxPayHRX` | `OpenBox` (payable) | 开盒用 HRX |
| `quoteStorageSlotCost` / `quoteStorageSlotCostInB` | `PreviewSlotCost` | 报价接口拆分 |
| `quoteOpenCountCost` / `quoteOpenCountCostInB` | `PreviewSlotCost` | 同上 |
| `baseDailyOpens` (变量) | `BASE_DAILY_OPENS` (常量) | 类型变更 |
| `baseStorageSlots`, `maxStorageSlotsOf` | 无 | 新增 |
| `burnAfterReward`, `globalSlotStepCost`, `openFeeBNB` | 有 | 合约已删除 |

---

## 三、业务代码迁移

当前项目中，对旧 NFT API 的调用**仅出现在** `internal/service/nft.go`（由 abigen 生成的绑定）。**未见其他业务代码直接调用** `ExtraSlots`、`BuyOpenSlots`、`OpenBox`、`PreviewSlotCost` 等。

若存在通过接口/封装调用这些方法的情况，请在 `make nft-abi` 后按新 API 调整：

- **槽位**：`ExtraStorageSlots` + `BuyStorageSlots` / `BuyStorageSlotsPayHRX`
- **开盒次数**：`ExtraOpenCount` + `BuyOpenCount` / `BuyOpenCountPayHRX`
- **开盒**：`OpenBoxPayUsdt` / `OpenBoxPayHRX`（替代 `OpenBox`）
- **报价**：`QuoteStorageSlotCost` / `QuoteOpenCountCost`（替代 `PreviewSlotCost`）

---

## 四、合约更新后的 ABI 同步流程

1. 在 `box` 中修改 `contracts/Nft.sol` 并编译：
   ```bash
   cd box && npx hardhat compile
   ```
2. 将 ABI 提取到 requestEth：
   ```powershell
   cd d:\nft-blind\box
   node -e "const fs=require('fs'); const j=JSON.parse(fs.readFileSync('artifacts/contracts/Nft.sol/BlindBoxNFT.json','utf8')); fs.writeFileSync('../requestEth/abi/box.abi', JSON.stringify(j.abi, null, 2));"
   cp ../requestEth/abi/box.abi ../requestEth/abi/nft.abi
   ```
3. 重新生成 Go 绑定：
   ```bash
   cd requestEth && make nft-abi
   ```

---

## 五、事件与 transaction 解析

`transaction.go` 中的事件解析（Opened、Purchased 等）使用内联 ABI，与当前合约的事件定义一致，无需修改。
