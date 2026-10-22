#!/bin/bash
# 向指定地址发送测试代币
# 项目代币：HRX、Piece（部署者/initialRecipient 有余额）
# 主网代币：BNB（gas）、USDT（需真有余额才能发）
#
# 用法: PRIVATE_KEY=你的私钥 ./send-test-tokens.sh [目标地址]
# 或: export PRIVATE_KEY=xxx && ./send-test-tokens.sh 0xE2D4ecEcb19B39D7e796fF9A4965543AeD903069

TARGET="${1:-0xE2D4ecEcb19B39D7e796fF9A4965543AeD903069}"
API_BASE="${API_BASE:-http://127.0.0.1:8003}"

# 合约地址（BSC 主网，见 box/ignition/deployments/chain-56/deployed_addresses.json）
HRX="0x057B67e0645D4db6A7373C7EA6fA6eeD4f218fFD"   # 项目代币，部署时 mint 10亿给 deployer
PIECE="0xa18144b788732760b037fff964E29E245C84692a" # 项目代币，mint 给 initialRecipient
USDT="0x55d398326f99059fF775485246999027B3197955"  # BSC 主网 USDT（非项目代币）

if [ -z "$PRIVATE_KEY" ]; then
  echo "错误: 请设置环境变量 PRIVATE_KEY"
  echo "用法: PRIVATE_KEY=0x你的64位私钥 ./send-test-tokens.sh $TARGET"
  echo ""
  echo "代币说明："
  echo "  - HRX:  部署者钱包有 10 亿枚，用于开盒/槽位/手续费"
  echo "  - Piece: initialRecipient 有约 11.7 亿枚，用于 burn 增加 holdingValue"
  echo "  - BNB:  gas，需主网真实 BNB"
  echo "  - USDT: 主网真实 USDT，需有余额"
  exit 1
fi

echo "目标地址: $TARGET"
echo "API: $API_BASE"
echo ""

# 1. BNB（gas，0.01 BNB = 10000000000000000 wei）
echo ">>> 发送 0.01 BNB..."
curl -s -X POST "$API_BASE/api/transaction_eth" \
  -H "Content-Type: application/json" \
  -d "{\"privateKey\":\"$PRIVATE_KEY\",\"toAddr\":\"$TARGET\",\"amount\":\"10000000000000000\"}" | jq .
echo ""

# 2. HRX（项目测试代币，1000 * 10^18）
echo ">>> 发送 1000 HRX..."
curl -s -X POST "$API_BASE/api/transaction" \
  -H "Content-Type: application/json" \
  -d "{\"privateKey\":\"$PRIVATE_KEY\",\"toAddr\":\"$TARGET\",\"token\":\"$HRX\",\"amount\":\"1000000000000000000000\"}" | jq .
echo ""

# 3. Piece（项目测试代币，1000 * 10^18）
echo ">>> 发送 1000 Piece..."
curl -s -X POST "$API_BASE/api/transaction" \
  -H "Content-Type: application/json" \
  -d "{\"privateKey\":\"$PRIVATE_KEY\",\"toAddr\":\"$TARGET\",\"token\":\"$PIECE\",\"amount\":\"1000000000000000000000\"}" | jq .
echo ""

# 4. USDT（主网真实代币，需有余额；10 * 10^18）
echo ">>> 发送 10 USDT（若钱包有余额）..."
curl -s -X POST "$API_BASE/api/transaction" \
  -H "Content-Type: application/json" \
  -d "{\"privateKey\":\"$PRIVATE_KEY\",\"toAddr\":\"$TARGET\",\"token\":\"$USDT\",\"amount\":\"10000000000000000000\"}" | jq .
echo ""
echo "完成。可在 BSCScan 查看: https://bscscan.com/address/$TARGET"
