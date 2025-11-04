#!/bin/bash

# 启动所有服务
set -e

echo "启动 CampusLink 所有服务..."

# 服务列表和端口
declare -A SERVICES=(
  ["auth-srv"]="9001:10001"
  ["user-srv"]="9002:10002"
  ["product-srv"]="9003:10003"
  ["task-srv"]="9004:10004"
  ["order-srv"]="9005:10005"
  ["payment-srv"]="9006:10006"
  ["search-srv"]="9007:10007"
  ["admin-srv"]="9008:10008"
  ["biz-srv"]="9009:10009"
)

# 创建日志目录
mkdir -p logs

# 启动每个服务
for service in "${!SERVICES[@]}"; do
  echo "启动 $service..."
  ports=${SERVICES[$service]}
  grpc_port=$(echo $ports | cut -d':' -f1)
  http_port=$(echo $ports | cut -d':' -f2)
  
  nohup ./bin/$service > logs/$service.log 2>&1 &
  echo $! > logs/$service.pid
  echo "✓ $service 已启动 (gRPC: $grpc_port, HTTP: $http_port, PID: $(cat logs/$service.pid))"
done

echo "✓ 所有服务已启动！"
echo "查看日志: tail -f logs/<service-name>.log"


