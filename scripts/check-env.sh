#!/bin/bash

# 环境检查脚本
echo "🔍 检查 CampusLink 开发环境..."
echo ""

# 颜色定义
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 检查结果
PASSED=0
FAILED=0

# 检查 Go
echo -n "检查 Go 版本... "
if command -v go &> /dev/null; then
    GO_VERSION=$(go version | awk '{print $3}')
    echo -e "${GREEN}✓${NC} $GO_VERSION"
    ((PASSED++))
else
    echo -e "${RED}✗${NC} 未安装 Go"
    echo "   请访问: https://go.dev/doc/install"
    ((FAILED++))
fi

# 检查 Docker
echo -n "检查 Docker... "
if command -v docker &> /dev/null; then
    DOCKER_VERSION=$(docker --version | awk '{print $3}' | sed 's/,//')
    echo -e "${GREEN}✓${NC} $DOCKER_VERSION"
    ((PASSED++))
else
    echo -e "${RED}✗${NC} 未安装 Docker"
    echo "   请访问: https://docs.docker.com/get-docker/"
    ((FAILED++))
fi

# 检查 Docker Compose
echo -n "检查 Docker Compose... "
if command -v docker-compose &> /dev/null; then
    COMPOSE_VERSION=$(docker-compose --version | awk '{print $4}' | sed 's/,//')
    echo -e "${GREEN}✓${NC} $COMPOSE_VERSION"
    ((PASSED++))
else
    echo -e "${RED}✗${NC} 未安装 Docker Compose"
    echo "   Docker Desktop 通常已包含"
    ((FAILED++))
fi

# 检查 Make
echo -n "检查 Make... "
if command -v make &> /dev/null; then
    MAKE_VERSION=$(make --version | head -n 1)
    echo -e "${GREEN}✓${NC} 已安装"
    ((PASSED++))
else
    echo -e "${YELLOW}⚠${NC} 未安装 Make (可选)"
    echo "   macOS: brew install make"
    echo "   Ubuntu: sudo apt install make"
fi

# 检查 Protoc (可选)
echo -n "检查 Protobuf 编译器... "
if command -v protoc &> /dev/null; then
    PROTOC_VERSION=$(protoc --version)
    echo -e "${GREEN}✓${NC} $PROTOC_VERSION"
else
    echo -e "${YELLOW}⚠${NC} 未安装 protoc (可选)"
    echo "   若需修改 proto 文件，请安装: https://grpc.io/docs/protoc-installation/"
fi

# 检查 Wire (可选)
echo -n "检查 Wire... "
if command -v wire &> /dev/null; then
    echo -e "${GREEN}✓${NC} 已安装"
else
    echo -e "${YELLOW}⚠${NC} 未安装 Wire (可选)"
    echo "   安装: go install github.com/google/wire/cmd/wire@latest"
fi

# 检查 Docker 守护进程
echo -n "检查 Docker 守护进程... "
if docker info &> /dev/null; then
    echo -e "${GREEN}✓${NC} 运行中"
    ((PASSED++))
else
    echo -e "${RED}✗${NC} Docker 未运行"
    echo "   请启动 Docker Desktop"
    ((FAILED++))
fi

# 检查端口占用
echo ""
echo "检查端口占用情况..."
PORTS=(3306 6379 8500 4222 9200 9000 8000)
PORT_NAMES=("MySQL" "Redis" "Consul" "NATS" "Elasticsearch" "MinIO" "APISIX")

for i in "${!PORTS[@]}"; do
    PORT=${PORTS[$i]}
    NAME=${PORT_NAMES[$i]}
    echo -n "  端口 $PORT ($NAME)... "
    if lsof -Pi :$PORT -sTCP:LISTEN -t >/dev/null 2>&1 ; then
        echo -e "${YELLOW}⚠${NC} 已被占用"
        echo "    PID: $(lsof -Pi :$PORT -sTCP:LISTEN -t)"
    else
        echo -e "${GREEN}✓${NC} 可用"
    fi
done

# 总结
echo ""
echo "================================"
if [ $FAILED -eq 0 ]; then
    echo -e "${GREEN}✓ 环境检查通过！ ($PASSED 项检查)${NC}"
    echo ""
    echo "下一步:"
    echo "  1. 启动基础服务: docker-compose up -d"
    echo "  2. 运行认证服务: cd app/auth-srv && go run cmd/server/main.go"
    echo "  3. 查看文档: cat QUICKSTART.md"
else
    echo -e "${RED}✗ 环境检查失败！ ($FAILED 项需要修复)${NC}"
    echo ""
    echo "请先安装缺失的依赖，然后重新运行此脚本"
fi
echo "================================"


