# 测试部署说明

## 主网测试（推荐）

在主网用 MockUSDT 测试，无需测试网 tBNB：

```bash
cd box
npx hardhat compile
npm run deploy:bsc:test
```

部署后 deployer 有 1000 万 MockUSDT，可 `usdt.mint(地址, 数量)` 给测试地址打 U。

打 U 脚本（主网）：
```bash
MOCK_USDT_ADDR=0x部署的MockUSDT地址 MINT_TO=0x目标地址 npx hardhat run scripts/mint-test-usdt.ts --network bsc
```

---

## 测试网部署

## 一、测试前需要准备

| 项目 | 说明 |
|------|------|
| **tBNB** | BSC 测试网原生代币，用于 gas。从水龙头领取 |
| **测试 U (MockUSDT)** | 部署时自动创建，部署者获得 1000 万枚，可 `mint()` 给其他地址 |
| **HRX** | 部署时自动 mint 10 亿枚给 deployer |
| **Piece** | 部署时 mint 给 initialRecipient |

---

## 二、获取 tBNB（测试网 BNB）

1. 打开 [Chainlink BSC Chapel Faucet](https://faucets.chain.link/bnb-chain-chapel)
2. 连接钱包，选择 BSC Testnet (Chapel)
3. 领取 tBNB（每次约 0.5 tBNB）

或使用 [BNB Chain 官方水龙头](https://www.bnbchain.org/en/faucet-smart)

---

## 三、部署步骤

### 1. 配置私钥

确保 `box/.env` 中有：

```
BSC_PRIVATE_KEY=你的64位私钥hex（可带或不带0x）
```

### 2. 编译

```bash
cd box
npx hardhat compile
```

### 3. 部署到 BSC 测试网

```bash
npx hardhat ignition deploy ./ignition/modules/BlindBoxTestnet.ts --network bscTestnet
```

### 4. 记录部署地址

部署完成后会输出各合约地址，保存到 `requestEth/.env` 和前端配置。

---

## 四、给测试地址打 U

MockUSDT 部署后，deployer 有 MINTER_ROLE，可调用：

```solidity
// 给地址 0x... 打 10000 个测试 U（18 位小数）
usdt.mint(0x目标地址, 10000 * 10**18);
```

或使用项目脚本：

```bash
# 部署后设置 MockUSDT 地址，然后执行
MOCK_USDT_ADDR=0x部署的地址 MINT_TO=0x目标地址 MINT_AMOUNT=10000000000000000000000 npx hardhat run scripts/mint-test-usdt.ts --network bscTestnet
```

---

## 五、requestEth 后端配置

部署完成后，更新 `requestEth/.env`：

```
# 测试网合约地址（替换为实际部署地址）
ETH_NFT_CONTRACT=0x...
ETH_STAKE_CONTRACT=0x...
ETH_SWAP_CONTRACT=0x...
ETH_PRIMARY_MARKET_CONTRACT=0x...
ETH_VAULT_FACTORY_CONTRACT=0x...
ETH_USDT_CONTRACT=0x...   # MockUSDT 地址
ETH_BSC_RPC_URLS=https://bsc-testnet-rpc.publicnode.com,https://bsc-testnet.bnbchain.org
```

### 团队长资格门槛

- `Stake.sol` 部署时需要显式传入 `minStakeA`
- 当前部署模块默认使用 `300e18`
- 规则为：
  - 成为 `top` 不等于自动拥有团队长收益资格
  - 只有满足 `stakedA >= minStakeA`，或被管理员 `setTopManualEligible(top, true)` 人工特批，才算合格团队长

---

## 六、前端配置

- **网络**：BSC Testnet (Chain ID 97)
- **RPC**：https://bsc-testnet-rpc.publicnode.com
- **区块浏览器**：https://testnet.bscscan.com
- **合约地址**：使用部署输出的地址

---

## 七、主网 vs 测试网

| 项目 | 主网 (chain-56) | 测试网 (chain-97) |
|------|-----------------|-------------------|
| USDT | 0x55d398...（真实） | MockUSDT（可 mint） |
| WBNB | 0xbb4CdB... | 0xae13d9... |
| 部署模块 | BlindBox.ts | BlindBoxTestnet.ts |
| 网络 | bsc | bscTestnet |
