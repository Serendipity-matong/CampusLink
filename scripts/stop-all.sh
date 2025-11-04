#!/bin/bash

# 停止所有服务
set -e

echo "停止 CampusLink 所有服务..."

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

# 停止每个服务
for service in "${SERVICES[@]}"; do
  if [ -f "logs/$service.pid" ]; then
    pid=$(cat logs/$service.pid)
    if ps -p $pid > /dev/null 2>&1; then
      echo "停止 $service (PID: $pid)..."
      kill $pid
      rm logs/$service.pid
      echo "✓ $service 已停止"
    else
      echo "⚠ $service 未运行"
      rm logs/$service.pid
    fi
  else
    echo "⚠ $service PID 文件不存在"
  fi
done

echo "✓ 所有服务已停止！"


