#!/bin/bash

# 快速创建微服务脚本
# 用法: ./create-service.sh <service-name> <grpc-port> <http-port>

set -e

SERVICE_NAME=$1
GRPC_PORT=$2
HTTP_PORT=$3

if [ -z "$SERVICE_NAME" ] || [ -z "$GRPC_PORT" ] || [ -z "$HTTP_PORT" ]; then
  echo "用法: $0 <service-name> <grpc-port> <http-port>"
  echo "示例: $0 product-srv 9003 10003"
  exit 1
fi

BASE_DIR="/Users/fangzijie/fzj/project/service/CampusLink"
APP_DIR="$BASE_DIR/app/$SERVICE_NAME"

echo "🚀 开始创建 $SERVICE_NAME..."

# 1. 复制目录结构
echo "1️⃣  复制目录结构..."
cp -r "$BASE_DIR/app/user-srv" "$APP_DIR"

# 2. 修改 go.mod
echo "2️⃣  修改 go.mod..."
sed -i '' "s/user-srv/$SERVICE_NAME/g" "$APP_DIR/go.mod"

# 3. 创建配置文件
echo "3️⃣  创建配置文件..."
cat > "$BASE_DIR/configs/$SERVICE_NAME.yaml" << EOF
server:
  http:
    network: tcp
    addr: 0.0.0.0:$HTTP_PORT
    timeout: 10000000000
  grpc:
    network: tcp
    addr: 0.0.0.0:$GRPC_PORT
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

# 4. 替换代码中的服务名
echo "4️⃣  替换代码中的服务名..."
find "$APP_DIR" -type f -name "*.go" -exec sed -i '' "s/user-srv/$SERVICE_NAME/g" {} +
find "$APP_DIR" -type f -name "*.go" -exec sed -i '' "s/user\/v1/${SERVICE_NAME/-srv/}\/v1/g" {} +
find "$APP_DIR" -type f -name "*.go" -exec sed -i '' "s/UserService/${SERVICE_NAME^}Service/g" {} +
find "$APP_DIR" -type f -name "*.go" -exec sed -i '' "s/NewUserService/New${SERVICE_NAME^}Service/g" {} +

# 5. 生成 Proto 代码
echo "5️⃣  生成 Proto 代码..."
PROTO_DIR=$(echo $SERVICE_NAME | sed 's/-srv//')
cd "$BASE_DIR/api/$PROTO_DIR/v1"
export PATH=$PATH:$(go env GOPATH)/bin
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $PROTO_DIR.proto

# 6. 生成 Wire 代码
echo "6️⃣  生成 Wire 代码..."
cd "$APP_DIR"
go mod tidy
cd cmd/server
rm -f wire_gen.go
wire

# 7. 编译服务
echo "7️⃣  编译服务..."
cd "$APP_DIR"
go build -o "$BASE_DIR/bin/$SERVICE_NAME" ./cmd/server

echo "✅ $SERVICE_NAME 创建成功！"
echo ""
echo "启动命令:"
echo "  cd $BASE_DIR && ./bin/$SERVICE_NAME -conf ./configs/$SERVICE_NAME.yaml"
echo ""
echo "端口信息:"
echo "  gRPC: $GRPC_PORT"
echo "  HTTP: $HTTP_PORT"


