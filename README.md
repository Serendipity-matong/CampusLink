# CampusLink - 校园互联平台

> 基于 Go 微服务架构的校园生活服务平台，对标 mall4cloud 的 Go 语言实现

## 📋 项目简介

CampusLink 是一个面向校园场景的综合服务平台，支持：
- 🏃 **跑腿服务**：校园代取快递、代买商品
- 📚 **二手交易**：二手书籍、电子产品买卖
- 🎒 **物品租借**：自行车、运动器材、电子设备租赁
- 👨‍🏫 **代课服务**：课程代签、笔记分享
- 🛍️ **校园商城**：日用品、零食饮料等

## 🏗️ 系统架构

### 技术栈

| 类型 | 技术选型 |
|------|---------|
| **核心框架** | Go-Kratos |
| **服务治理** | Consul + gRPC |
| **网关** | Apache APISIX |
| **数据库** | MySQL + GORM |
| **缓存** | Redis (go-redis) |
| **消息队列** | NATS JetStream |
| **搜索引擎** | Elasticsearch |
| **分布式事务** | DTM |
| **定时任务** | gocron |
| **对象存储** | MinIO |
| **文档** | go-swagger |

### 微服务列表

| 服务名 | 端口 (gRPC/HTTP) | 功能 |
|--------|------------------|------|
| `gateway` | -/8000 | API网关 |
| `auth-srv` | 9001/10001 | 认证授权 |
| `user-srv` | 9002/10002 | 用户管理 |
| `product-srv` | 9003/10003 | 二手书/租借品 |
| `task-srv` | 9004/10004 | 跑腿/代课 |
| `order-srv` | 9005/10005 | 订单管理 |
| `payment-srv` | 9006/10006 | 支付服务 |
| `search-srv` | 9007/10007 | 搜索服务 |
| `admin-srv` | 9008/10008 | 后台管理 |
| `biz-srv` | 9009/10009 | 通用业务(短信/上传) |

## 📁 项目结构

```
campus-link/
├── api/                # Proto 定义和生成的代码
│   ├── auth/v1/
│   ├── user/v1/
│   ├── product/v1/
│   ├── task/v1/
│   ├── order/v1/
│   ├── payment/v1/
│   ├── search/v1/
│   ├── admin/v1/
│   └── biz/v1/
│
├── app/                # 微服务实现
│   ├── gateway/
│   ├── auth-srv/
│   ├── user-srv/
│   ├── product-srv/
│   ├── task-srv/
│   ├── order-srv/
│   ├── payment-srv/
│   ├── search-srv/
│   ├── admin-srv/
│   └── biz-srv/
│
├── common/             # 公共库
│   ├── common-core/    # 核心工具
│   ├── common-db/      # 数据库
│   ├── common-cache/   # 缓存
│   ├── common-auth/    # 认证
│   ├── common-mq/      # 消息队列
│   └── common-dtm/     # 分布式事务
│
├── configs/            # 配置文件
├── deployments/        # 部署文件
│   ├── docker/
│   ├── k8s/
│   └── apisix/
│
└── scripts/            # 脚本工具
```

## 🚀 快速开始

### 环境要求

- Go 1.21+
- MySQL 8.0+
- Redis 6.0+
- Consul 1.16+
- NATS 2.10+
- Elasticsearch 7.x
- Docker & Docker Compose (可选)

### 本地开发

1. **克隆项目**
```bash
git clone https://github.com/campuslink/campuslink.git
cd campuslink
```

2. **安装依赖**
```bash
go mod download
```

3. **生成 Proto 代码**
```bash
make api
```

4. **启动基础服务** (使用 Docker Compose)
```bash
docker-compose up -d mysql redis consul nats elasticsearch
```

5. **启动微服务** (以 user-srv 为例)
```bash
cd app/user-srv
go run cmd/server/main.go
```

6. **启动网关**
```bash
cd app/gateway
go run main.go
```

### Docker 部署

```bash
# 构建所有服务
make docker-build

# 启动所有服务
docker-compose up -d
```

### Kubernetes 部署

```bash
# 应用 K8s 配置
kubectl apply -f deployments/k8s/

# 查看服务状态
kubectl get pods -n campuslink
```

## 📖 API 文档

启动服务后访问：
- Swagger UI: http://localhost:8000/swagger
- API 文档: http://localhost:8000/docs

## 🔧 开发指南

### 创建新服务

```bash
# 使用 Kratos CLI 创建新服务
kratos new app/new-srv
cd app/new-srv
kratos proto add api/new-srv/v1/new.proto
make api
```

### 代码生成

```bash
# 生成 Proto 代码
make api

# 生成 Wire 依赖注入
make generate

# 生成 Mock 测试
make mock
```

### 代码规范

- 遵循 [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- 使用 `golangci-lint` 进行代码检查
- 提交前运行 `make lint` 和 `make test`

## 🧪 测试

```bash
# 单元测试
make test

# 集成测试
make integration-test

# 覆盖率报告
make coverage
```

## 📊 监控

- **链路追踪**: Jaeger (http://localhost:16686)
- **指标监控**: Prometheus + Grafana (http://localhost:3000)
- **日志聚合**: ELK Stack

## 🛠️ Makefile 命令

```bash
make api           # 生成 Proto 代码
make build         # 编译所有服务
make test          # 运行测试
make lint          # 代码检查
make docker-build  # 构建 Docker 镜像
make deploy        # 部署到 K8s
```

## 📝 License

MIT License

## 👥 贡献

欢迎提交 Issue 和 Pull Request！

## 📧 联系方式

- 项目地址: https://github.com/campuslink/campuslink
- 问题反馈: https://github.com/campuslink/campuslink/issues


