# 重新部署 BlindBoxMainnetTest（BSC 主网）

在确认 **[链上与当前编译不一致](../../docs/链上与仓库差异-BlindBoxMainnetTest.md)** 后，若要让主网行为与 **当前仓库** 一致，需要完整重新部署。**会得到一整套新合约地址**，原地址上的状态与数据不会自动迁移。

---

## 1. 前置条件

- BSC 主网 **BNB**（支付 gas，建议预留充足）。
- 部署钱包私钥：`box/.env` 中配置 `BSC_PRIVATE_KEY`（格式见 `box/.env.example`）。
- Node 依赖：在 `box` 目录已 `npm install`（若尚未安装）。

---

## 2. 编译与 ABI 导出

```bash
cd box
npx hardhat compile
node scripts/export-frontend-abis.mjs
```

会将 ABI 写入 `requestEth/abi/frontend_all_contracts.json`。

---

## 3. 执行 Ignition 部署

模块路径：`ignition/modules/BlindBoxMainnetTest.ts`。

对 **同一模块** 做「全新一整套」部署时，使用 **`--reset`** 清除本地 Ignition 中该部署的状态后再部署（否则会认为已部署过而跳过）：

```bash
cd box
npx hardhat ignition deploy ./ignition/modules/BlindBoxMainnetTest.ts --network bsc --reset
```

> **安全提示**：`--reset` 仅影响本机 `ignition/deployments/` 中记录的部署状态；链上仍会 **新部署** 合约，**不会**删除旧合约。请确认使用正确的网络和私钥。

可选：

- `--verify`：若已配置 BscScan API key，可在验证成功后自动验证合约源码（需按 Hardhat 官方文档配置 `etherscan`）。
- `--parameters <file.json>`：覆盖模块中的 `m.getParameter(...)`（如 `accountA`、`initialRecipient` 等）。

部署结束后，Ignition 会更新 `ignition/deployments/chain-56/deployed_addresses.json`，其中 **`BlindBoxMainnetTest#*`** 即为新地址。

---

## 4. 部署后必做（回填清单）

1. **前端** `docs/前端合约地址-BlindBoxMainnetTest.md`  
   - 替换表中所有 `BlindBoxMainnetTest` 相关地址（`VITE_*` 与一览表）。

2. **后端 requestEth** `requestEth/docs/DEPLOYMENT.md` 中的示例 env：  
   - `ETH_NFT_CONTRACT`、`ETH_STAKE_CONTRACT`、`ETH_SWAP_CONTRACT`、`ETH_PRIMARY_MARKET_CONTRACT`、`ETH_VAULT_FACTORY_CONTRACT`、`ETH_USDT_CONTRACT` 等与新部署一致。

3. **事件同步起始块**  
   - `ETH_DEPLOY_BLOCK_NFT_MINT`、`ETH_DEPLOY_BLOCK_NFT_LISTED`（或项目中等价变量）改为 **新 `BlindBoxNFT` 合约的部署区块**（可从部署 receipt / BscScan 查询），否则 `get_address_box` 等会漏事件。

4. **测试用 MockUSDT**  
   - 部署模块会部署新的 `MockUSDT`；需对测试钱包执行 `mint`（若合约提供）或从部署者余额转出。

5. **单文件 ABI**（若仍引用路径方式）：  
   - 按 `docs/前端合约地址-BlindBoxMainnetTest.md` 第五节从 `artifacts` 再导出 `nft.abi`、`swap_hrx_usdt.abi` 等。

---

## 5. 本仓库无法代你执行的一步

实际 **上链部署** 必须在持有私钥与 BNB 的环境中由你或运维执行；CI 无默认私钥，故无法在仓库内替你完成交易。
