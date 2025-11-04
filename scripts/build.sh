#!/bin/bash

# 构建脚本
set -e

echo "开始构建 CampusLink 项目..."

# 服务列表
SERVICES=(
  "auth-srv"
  "user-srv"
  "product-srv"
  "task-srv"
  "order-srv"
  "payment-srv"
  "search-srv"
  "admin-srv"
  "biz-srv"
)

# 创建 bin 目录
mkdir -p bin

# 构建每个服务
for service in "${SERVICES[@]}"; do
  echo "构建 $service..."
  cd app/$service
  go build -o ../../bin/$service cmd/server/main.go
  cd ../..
  echo "✓ $service 构建完成"
done

echo "✓ 所有服务构建完成！"


