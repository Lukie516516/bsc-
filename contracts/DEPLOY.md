# BSC 主网部署指南

## 1. 环境准备

```bash
cd box
pnpm install   # 或 npm install
```

## 2. 配置私钥

在 `box` 目录下创建 `.env` 文件：

```bash
# 复制模板
cp .env.example .env

# 编辑 .env，填入私钥（hex，无 0x 前缀）
# BSC_PRIVATE_KEY=你的64位hex私钥
```

`.env` 已加入 .gitignore，不会被提交。

## 3. 执行部署

```bash
npx hardhat ignition deploy ./ignition/modules/BlindBox.ts --network bsc
```

### 自定义收款地址（可选）

默认 accountA/B/C/D 和 initialRecipient 均为 deployer。若要指定：

```bash
npx hardhat ignition deploy ./ignition/modules/BlindBox.ts --network bsc \
  --parameters '{"initialRecipient":"0x你的地址","accountA":"0x..."}'
```

## 4. 部署后操作

### 4.1 SwapHrxUsdt 添加流动性

Swap 需先初始化流动性才能交易：

1. 向 SwapHrxUsdt 合约转入 HRX 和 USDT
2. 调用 `mintInitial(token1BalanceInit)`（owner）

### 4.2 更新 requestEth 配置

将部署输出的合约地址填入环境变量：

```
ETH_NFT_CONTRACT=0x...
ETH_STAKE_CONTRACT=0x...
ETH_SWAP_CONTRACT=0x...
ETH_USDT_CONTRACT=0x55d398326f99059fF775485246999027B3197955
```

### 4.3 更新前端配置

rtxadmin、rtxdata 等前端需配置新的合约地址和 RPC。

## 5. BSC 主网常量

| 名称 | 地址 |
|------|------|
| USDT | 0x55d398326f99059fF775485246999027B3197955 |
| WBNB | 0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c |
| OnePiece 代币 | 0xe8575516Abc9E535d03cB9fCDAb13f18698F2117 |

## 6. 预估 Gas

全部部署约 0.05–0.15 BNB，建议准备 0.2–0.3 BNB。
