#!/bin/bash

# 快速创建所有剩余服务的简化脚本
set -e

BASE_DIR="/Users/fangzijie/fzj/project/service/CampusLink"
cd $BASE_DIR

echo "🚀 开始批量创建微服务..."
echo ""

# 服务列表：名称 gRPC端口 HTTP端口
services=(
  "product-srv 9003 10003"
  "task-srv 9004 10004"
  "order-srv 9005 10005"
)

for service_info in "${services[@]}"; do
  read -r name grpc http <<< "$service_info"
  
  echo "📦 创建 $name..."
  
  # 1. 复制结构
  if [ -d "app/$name" ]; then
    echo "   ⚠️  $name 已存在，跳过"
    continue
  fi
  
  cp -r app/user-srv "app/$name"
  
  # 2. 修改 go.mod
  sed -i '' "s/user-srv/$name/g" "app/$name/go.mod"
  
  # 3. 创建配置
  cat > "configs/$name.yaml" << EOF
server:
  http:
    network: tcp
    addr: 0.0.0.0:$http
    timeout: 10000000000
  grpc:
    network: tcp
    addr: 0.0.0.0:$grpc
    timeout: 10000000000

data:
  database:
    driver: mysql
    source: root:123@tcp(127.0.0.1:3306)/campuslink?charset=utf8mb4&parseTime=True&loc=Local
    max_idle_conns: 10
    max_open_conns: 100
    conn_max_lifetime: 3600000000000
  redis:
    addr: 127.0.0.1:6379
    password: ""
    db: 0
    dial_timeout: 5000000000
    read_timeout: 3000000000
    write_timeout: 3000000000
EOF

  # 4. 替换代码
  proto_name=$(echo $name | sed 's/-srv//')
  find "app/$name" -type f -name "*.go" -exec sed -i '' "s/user-srv/$name/g" {} +
  find "app/$name" -type f -name "*.go" -exec sed -i '' "s/user\\/v1/$proto_name\\/v1/g" {} +
  find "app/$name" -type f -name "*.go" -exec sed -i '' "s/UserService/${proto_name^}Service/g" {} +
  find "app/$name" -type f -name "*.go" -exec sed -i '' "s/NewUserService/New${proto_name^}Service/g" {} +
  find "app/$name" -type f -name "*.go" -exec sed -i '' "s/user\\.proto/$proto_name.proto/g" {} +
  
  # 5. 生成 Proto 代码
  cd "api/$proto_name/v1"
  export PATH=$PATH:$(go env GOPATH)/bin
  protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $proto_name.proto 2>/dev/null || true
  
  # 6. 编译
  cd "$BASE_DIR/app/$name"
  go mod tidy 2>&1 | grep -v "finding module" || true
  cd cmd/server
  rm -f wire_gen.go
  wire 2>/dev/null || true
  cd ../..
  go build -o "$BASE_DIR/bin/$name" ./cmd/server 2>&1 | grep -v "finding module" || true
  
  if [ -f "$BASE_DIR/bin/$name" ]; then
    echo "   ✅ $name 创建成功！(端口: $grpc/$http)"
  else
    echo "   ❌ $name 创建失败"
  fi
  
  cd $BASE_DIR
  echo ""
done

echo "🎉 服务创建完成！"
echo ""
echo "查看已编译的服务:"
ls -lh bin/*-srv 2>/dev/null | awk '{print "  " $9 " (" $5 ")"}'

