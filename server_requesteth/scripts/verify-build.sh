#!/bin/bash
# 在 Linux 服务器上运行此脚本验证 Docker 构建
# 用法: bash scripts/verify-build.sh

set -e
cd "$(dirname "$0")/.."

echo "=== 验证 requestEth Docker 构建 ==="
echo ""

# 检查 Docker
if ! command -v docker &> /dev/null; then
    echo "错误: 未找到 docker 命令"
    exit 1
fi

echo "1. 清理旧镜像..."
docker rmi requesteth-test 2>/dev/null || true

echo ""
echo "2. 开始构建 (预计 2-5 分钟)..."
if docker build -t requesteth-test . ; then
    echo ""
    echo "=== 构建成功 ==="
    echo "镜像: requesteth-test"
    echo ""
    echo "可运行: docker run --rm requesteth-test ./requestEth -conf /data/conf"
    exit 0
else
    echo ""
    echo "=== 构建失败 ==="
    exit 1
fi
