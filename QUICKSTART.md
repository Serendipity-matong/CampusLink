# CampusLink 快速开始指南

## 前置要求

### 必需软件

- **Go**: 1.21+ ([安装指南](https://go.dev/doc/install))
- **Docker**: 20.10+ ([安装指南](https://docs.docker.com/get-docker/))
- **Docker Compose**: 2.0+ (通常随Docker安装)
- **Make**: 构建工具

### 可选软件

- **Protobuf编译器**: 用于生成gRPC代码 ([安装指南](https://grpc.io/docs/protoc-installation/))
- **Wire**: Google依赖注入工具 (`go install github.com/google/wire/cmd/wire@latest`)
- **kubectl**: Kubernetes命令行工具 (生产部署需要)

## 快速启动 (5分钟)

### 步骤1: 克隆项目

```bash
git clone https://github.com/campuslink/campuslink.git
cd campuslink
```

### 步骤2: 启动基础服务

使用Docker Compose启动所有依赖服务：

```bash
docker-compose up -d
```

这将启动：
- ✅ MySQL (端口 3306)
- ✅ Redis (端口 6379)
- ✅ Consul (端口 8500)
- ✅ NATS (端口 4222)
- ✅ Elasticsearch (端口 9200)
- ✅ MinIO (端口 9000, 9001)
- ✅ DTM (端口 36789, 36790)
- ✅ Jaeger (端口 16686)
- ✅ Prometheus (端口 9090)
- ✅ Grafana (端口 3000)
- ✅ APISIX (端口 8000)

查看服务状态：
```bash
docker-compose ps
```

### 步骤3: 配置数据库

等待MySQL启动完成后，数据库会自动初始化。可以连接MySQL验证：

```bash
mysql -h 127.0.0.1 -P 3306 -u root -ppassword campuslink
```

### 步骤4: 启动认证服务 (auth-srv)

```bash
# 下载依赖
cd app/auth-srv
go mod tidy

# 运行服务
go run cmd/server/main.go -conf ../../configs/auth-srv.yaml
```

服务启动后会监听：
- 🌐 HTTP: `http://localhost:10001`
- 📡 gRPC: `localhost:9001`

### 步骤5: 测试API

#### 5.1 注册用户

```bash
curl -X POST http://localhost:10001/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "student01",
    "password": "password123",
    "phone": "13800138888",
    "student_id": "2024001",
    "real_name": "张三",
    "school": "某某大学",
    "verification_code": "123456"
  }'
```

#### 5.2 用户登录

```bash
curl -X POST http://localhost:10001/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "student01",
    "password": "password123",
    "login_type": "username"
  }'
```

返回示例：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
    "expires_in": 7200,
    "user_info": {
      "user_id": 1,
      "username": "student01",
      "phone": "13800138888",
      "student_id": "2024001",
      "real_name": "张三",
      "school": "某某大学"
    }
  }
}
```

## 启动所有服务

### 方法1: 手动启动 (开发调试)

在不同的终端窗口启动每个服务：

```bash
# 终端1: auth-srv
cd app/auth-srv && go run cmd/server/main.go

# 终端2: user-srv  (需先创建)
cd app/user-srv && go run cmd/server/main.go

# 终端3: product-srv (需先创建)
cd app/product-srv && go run cmd/server/main.go

# ... 其他服务类似
```

### 方法2: 使用脚本 (推荐)

```bash
# 编译所有服务
./scripts/build.sh

# 启动所有服务
./scripts/run-all.sh

# 停止所有服务
./scripts/stop-all.sh
```

### 方法3: 使用Makefile

```bash
# 编译
make build

# 运行所有服务
make run-all

# 停止所有服务
make stop-all
```

## 访问各项服务

### 管理界面

| 服务 | 地址 | 用户名 | 密码 |
|------|------|--------|------|
| Consul | http://localhost:8500 | - | - |
| MinIO | http://localhost:9001 | minioadmin | minioadmin |
| Jaeger | http://localhost:16686 | - | - |
| Prometheus | http://localhost:9090 | - | - |
| Grafana | http://localhost:3000 | admin | admin |
| APISIX | http://localhost:9180 | - | - |

### API端点

| 服务 | HTTP端口 | gRPC端口 | 健康检查 |
|------|----------|----------|----------|
| auth-srv | 10001 | 9001 | http://localhost:10001/health |
| user-srv | 10002 | 9002 | http://localhost:10002/health |
| product-srv | 10003 | 9003 | http://localhost:10003/health |
| task-srv | 10004 | 9004 | http://localhost:10004/health |
| order-srv | 10005 | 9005 | http://localhost:10005/health |
| payment-srv | 10006 | 9006 | http://localhost:10006/health |
| search-srv | 10007 | 9007 | http://localhost:10007/health |
| admin-srv | 10008 | 9008 | http://localhost:10008/health |
| biz-srv | 10009 | 9009 | http://localhost:10009/health |

## 开发新服务

### 使用Kratos CLI (推荐)

```bash
# 安装Kratos CLI
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest

# 创建新服务 (以order-srv为例)
cd app
kratos new order-srv

# 创建API定义
cd order-srv
kratos proto add api/order/v1/order.proto

# 生成代码
kratos proto client api/order/v1/order.proto

# 生成Wire依赖注入
cd cmd/server
wire
```

### 手动创建 (参考auth-srv)

1. 复制`app/auth-srv`目录结构
2. 修改`go.mod`中的模块名
3. 调整`internal/biz`、`internal/data`、`internal/service`的业务逻辑
4. 创建对应的配置文件 `configs/xxx-srv.yaml`

## 常见问题

### Q1: 服务启动失败，提示连接MySQL失败？

**A**: 确保MySQL已启动并完成初始化：

```bash
# 查看MySQL日志
docker-compose logs mysql

# 等待MySQL完全启动 (大约30秒)
docker-compose ps mysql
```

### Q2: gRPC调用失败？

**A**: 检查proto文件是否已生成：

```bash
# 重新生成proto代码
cd api/auth/v1
protoc --go_out=. --go-grpc_out=. --go-http_out=. auth.proto
```

### Q3: Wire依赖注入错误？

**A**: 重新生成Wire代码：

```bash
cd app/auth-srv/cmd/server
wire
```

### Q4: 端口被占用？

**A**: 修改配置文件中的端口或关闭占用端口的进程：

```bash
# 查找占用端口的进程 (Mac/Linux)
lsof -i :10001

# 杀死进程
kill -9 <PID>
```

### Q5: Redis连接失败？

**A**: 检查Redis是否启动：

```bash
docker-compose ps redis
docker-compose logs redis
```

## 下一步

1. ✅ [阅读架构设计文档](./ARCHITECTURE.md)
2. ✅ [查看API使用示例](./API_EXAMPLES.md)
3. ✅ [学习如何部署到Kubernetes](./deployments/k8s/README.md)
4. ✅ [了解如何贡献代码](./CONTRIBUTING.md)

## 获取帮助

- 📖 [完整文档](./README.md)
- 💬 [提交Issue](https://github.com/campuslink/campuslink/issues)
- 📧 联系我们: dev@campuslink.com


