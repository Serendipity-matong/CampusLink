# CampusLink 项目总结

## 📊 项目概览

**CampusLink (校园互联平台)** 是一个基于 Go 微服务架构的校园生活服务平台，对标 Java 生态的 mall4cloud 项目。本项目使用 Go-Kratos 框架，采用现代化的云原生技术栈，为校园场景提供二手交易、跑腿服务、物品租借等功能。

### 核心特性

✅ **Go 微服务架构**: 基于 Go-Kratos 框架的完整微服务解决方案  
✅ **gRPC 通信**: 高性能的服务间通信  
✅ **云原生**: 支持 Docker 和 Kubernetes 部署  
✅ **分布式事务**: 使用 DTM 保证数据一致性  
✅ **服务治理**: Consul 注册发现 + APISIX 网关  
✅ **完整监控**: Prometheus + Grafana + Jaeger 全链路追踪  
✅ **数据搜索**: Elasticsearch 全文搜索  
✅ **消息队列**: NATS JetStream 异步处理  
✅ **对象存储**: MinIO 文件管理  

---

## 🏗️ 技术架构

### 架构对比

| 层次 | mall4cloud (Java) | CampusLink (Go) |
|------|-------------------|-----------------|
| **核心框架** | Spring Cloud 2020 | **Go-Kratos** |
| **ORM** | MyBatis | **GORM** |
| **服务注册** | Nacos 2 | **Consul** |
| **API网关** | Spring Cloud Gateway 2 | **Apache APISIX** |
| **负载均衡** | Spring Cloud Loadbalancer | **gRPC 内置 LB** |
| **服务调用** | Feign | **gRPC Client** |
| **分布式事务** | Seata | **DTM** |
| **缓存** | Redis | **Redis (go-redis)** |
| **消息队列** | RocketMQ | **NATS JetStream** |
| **搜索引擎** | Elasticsearch | **Elasticsearch** |
| **定时任务** | XXL-JOB | **gocron** |
| **文件存储** | MinIO | **MinIO** |
| **数据同步** | Canal | **Debezium / go-mysql** |

### 项目结构

```
CampusLink/
├── api/                    # Proto 定义 (9个服务)
│   ├── auth/v1/            # 认证接口
│   ├── user/v1/            # 用户接口
│   ├── product/v1/         # 商品接口
│   ├── task/v1/            # 任务接口
│   ├── order/v1/           # 订单接口
│   ├── payment/v1/         # 支付接口
│   ├── search/v1/          # 搜索接口
│   ├── admin/v1/           # 管理接口
│   └── biz/v1/             # 通用业务接口
│
├── app/                    # 微服务实现
│   ├── auth-srv/           # 认证服务 (已完成)
│   │   ├── cmd/server/     # 入口 + Wire
│   │   └── internal/       # 内部实现
│   │       ├── conf/       # 配置
│   │       ├── data/       # 数据层 (PO)
│   │       ├── biz/        # 业务层 (BO)
│   │       ├── service/    # 服务层 (DTO/VO)
│   │       └── server/     # gRPC/HTTP Server
│   ├── user-srv/           # 用户服务 (待创建)
│   ├── product-srv/        # 商品服务 (待创建)
│   ├── task-srv/           # 任务服务 (待创建)
│   ├── order-srv/          # 订单服务 (待创建)
│   ├── payment-srv/        # 支付服务 (待创建)
│   ├── search-srv/         # 搜索服务 (待创建)
│   ├── admin-srv/          # 管理服务 (待创建)
│   └── biz-srv/            # 通用业务 (待创建)
│
├── common/                 # 公共库 (已完成)
│   ├── common-core/        # 核心工具
│   │   ├── errors/         # 错误定义
│   │   ├── response/       # 响应封装
│   │   ├── logger/         # 日志封装
│   │   └── utils/          # 工具函数
│   ├── common-db/          # 数据库封装
│   ├── common-cache/       # Redis 封装
│   ├── common-auth/        # JWT 认证
│   ├── common-mq/          # NATS 封装
│   └── common-dtm/         # DTM 封装
│
├── configs/                # 配置文件
│   ├── auth-srv.yaml
│   └── ...
│
├── deployments/            # 部署配置
│   ├── docker/             # Docker 配置
│   │   ├── mysql/init.sql  # 数据库初始化
│   │   └── prometheus/     # 监控配置
│   ├── k8s/                # Kubernetes 配置
│   │   ├── namespace.yaml
│   │   └── auth-srv-deployment.yaml
│   └── apisix/             # 网关配置
│       └── config.yaml
│
├── scripts/                # 脚本工具
│   ├── build.sh            # 构建脚本
│   ├── run-all.sh          # 启动脚本
│   └── stop-all.sh         # 停止脚本
│
├── docker-compose.yaml     # Docker Compose 配置
├── Makefile                # Make 命令
├── go.mod                  # Go 依赖
├── README.md               # 项目说明
├── QUICKSTART.md           # 快速开始
├── ARCHITECTURE.md         # 架构文档
├── API_EXAMPLES.md         # API 示例
└── PROJECT_SUMMARY.md      # 项目总结 (本文档)
```

---

## 📦 已完成的组件

### ✅ 公共库 (common/)

| 模块 | 功能 | 状态 |
|------|------|------|
| `common-core/errors` | 统一错误码定义 | ✅ 完成 |
| `common-core/response` | 统一响应格式 | ✅ 完成 |
| `common-core/logger` | 日志封装 | ✅ 完成 |
| `common-core/utils` | 工具函数 (哈希/时间/字符串) | ✅ 完成 |
| `common-db` | GORM 数据库封装 | ✅ 完成 |
| `common-cache` | Redis 缓存封装 | ✅ 完成 |
| `common-auth` | JWT 认证封装 | ✅ 完成 |
| `common-mq` | NATS 消息队列封装 | ✅ 完成 |
| `common-dtm` | DTM 分布式事务封装 | ✅ 完成 |

### ✅ API 定义 (api/)

所有 9 个服务的 Proto 文件已完成：

| 服务 | Proto 文件 | 接口数量 | 状态 |
|------|-----------|----------|------|
| auth-srv | `api/auth/v1/auth.proto` | 6 | ✅ 完成 |
| user-srv | `api/user/v1/user.proto` | 7 | ✅ 完成 |
| product-srv | `api/product/v1/product.proto` | 9 | ✅ 完成 |
| task-srv | `api/task/v1/task.proto` | 9 | ✅ 完成 |
| order-srv | `api/order/v1/order.proto` | 8 | ✅ 完成 |
| payment-srv | `api/payment/v1/payment.proto` | 8 | ✅ 完成 |
| search-srv | `api/search/v1/search.proto` | 6 | ✅ 完成 |
| admin-srv | `api/admin/v1/admin.proto` | 10 | ✅ 完成 |
| biz-srv | `api/biz/v1/biz.proto` | 7 | ✅ 完成 |

### ✅ 微服务实现

| 服务 | 端口 | 状态 | 说明 |
|------|------|------|------|
| auth-srv | 9001/10001 | ✅ 完成 | 完整的 Kratos 服务实现 (含 Service/Biz/Data) |
| user-srv | 9002/10002 | ⏳ 待创建 | 可参考 auth-srv 结构创建 |
| product-srv | 9003/10003 | ⏳ 待创建 | 可参考 auth-srv 结构创建 |
| task-srv | 9004/10004 | ⏳ 待创建 | 可参考 auth-srv 结构创建 |
| order-srv | 9005/10005 | ⏳ 待创建 | 可参考 auth-srv 结构创建 |
| payment-srv | 9006/10006 | ⏳ 待创建 | 可参考 auth-srv 结构创建 |
| search-srv | 9007/10007 | ⏳ 待创建 | 可参考 auth-srv 结构创建 |
| admin-srv | 9008/10008 | ⏳ 待创建 | 可参考 auth-srv 结构创建 |
| biz-srv | 9009/10009 | ⏳ 待创建 | 可参考 auth-srv 结构创建 |

### ✅ 基础设施配置

| 组件 | 配置文件 | 状态 |
|------|----------|------|
| Docker Compose | `docker-compose.yaml` | ✅ 完成 |
| MySQL 初始化 | `deployments/docker/mysql/init.sql` | ✅ 完成 |
| APISIX 网关 | `deployments/apisix/config.yaml` | ✅ 完成 |
| Prometheus | `deployments/docker/prometheus/prometheus.yml` | ✅ 完成 |
| Kubernetes | `deployments/k8s/*.yaml` | ✅ 完成 |
| 构建脚本 | `scripts/build.sh` | ✅ 完成 |
| 启动脚本 | `scripts/run-all.sh` | ✅ 完成 |
| 停止脚本 | `scripts/stop-all.sh` | ✅ 完成 |

### ✅ 文档

| 文档 | 文件名 | 状态 |
|------|--------|------|
| 项目说明 | `README.md` | ✅ 完成 |
| 快速开始 | `QUICKSTART.md` | ✅ 完成 |
| 架构设计 | `ARCHITECTURE.md` | ✅ 完成 |
| API 示例 | `API_EXAMPLES.md` | ✅ 完成 |
| 项目总结 | `PROJECT_SUMMARY.md` | ✅ 完成 |

---

## 🔄 下一步工作

### 1. 创建其他微服务 (参考 auth-srv)

所有微服务都遵循相同的 Kratos 分层结构：

```bash
# 使用 Kratos CLI 创建 (推荐)
cd app
kratos new user-srv
cd user-srv
kratos proto add api/user/v1/user.proto
kratos proto client api/user/v1/user.proto
cd cmd/server && wire

# 或手动复制 auth-srv 结构并修改
cp -r app/auth-srv app/user-srv
# 然后修改业务逻辑
```

### 2. 生成 Proto 代码

```bash
# 为每个服务生成 gRPC 代码
make api

# 或手动生成
cd api/user/v1
protoc --go_out=. --go-grpc_out=. --go-http_out=. user.proto
```

### 3. 服务间通信

在需要调用其他服务时，注入 gRPC Client：

```go
// 在 data 层注入其他服务的客户端
type productRepo struct {
    data      *Data
    userClient userv1.UserClient  // 注入 user-srv 的客户端
}

// 调用
userInfo, err := r.userClient.GetUser(ctx, &userv1.GetUserRequest{
    UserId: productBO.SellerID,
})
```

### 4. 分布式事务 (DTM)

对于需要跨服务事务的场景：

```go
// 创建 Saga 事务
saga := r.dtm.NewSaga()
saga.Add(
    "product-srv/deduct-stock",     // 正向操作
    "product-srv/add-stock",        // 补偿操作
    &DeductStockRequest{...},
)
saga.Add(
    "order-srv/create-order",
    "order-srv/cancel-order",
    &CreateOrderRequest{...},
)
err := saga.Submit()
```

### 5. 搜索服务集成

```go
// 商品更新时同步到 ES
func (r *productRepo) UpdateProduct(ctx context.Context, product *biz.Product) error {
    // 1. 更新数据库
    err := r.data.db.Save(product).Error
    
    // 2. 发送消息到 NATS (异步)
    r.nats.Publish("product.updated", productJSON)
    
    return err
}

// search-srv 消费消息并更新 ES
r.nats.Subscribe("product.updated", func(msg *nats.Msg) {
    // 更新 Elasticsearch
    r.es.Index(ctx, "products", productID, productDoc)
})
```

---

## 🚀 快速启动

### 1. 启动基础服务

```bash
docker-compose up -d
```

### 2. 运行认证服务

```bash
cd app/auth-srv
go mod tidy
go run cmd/server/main.go -conf ../../configs/auth-srv.yaml
```

### 3. 测试API

```bash
# 注册
curl -X POST http://localhost:10001/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"123456","phone":"13800138888"}'

# 登录
curl -X POST http://localhost:10001/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"123456","login_type":"username"}'
```

---

## 📈 性能指标

### 预期性能

- **QPS**: 单服务 5000+ (gRPC)
- **延迟**: P99 < 100ms
- **并发**: 支持 10000+ 并发连接
- **可用性**: 99.9% (通过服务熔断、降级保证)

### 扩展性

- **水平扩展**: 无状态设计，可轻松扩展到数十个实例
- **垂直扩展**: 支持容器资源动态调整
- **数据库**: 支持读写分离、分库分表

---

## 🛡️ 安全性

- ✅ **密码加密**: bcrypt 哈希
- ✅ **Token 认证**: JWT + 黑名单机制
- ✅ **HTTPS**: 生产环境强制 HTTPS
- ✅ **SQL 注入防护**: GORM 参数化查询
- ✅ **XSS 防护**: 输入校验和转义
- ✅ **限流**: APISIX 网关限流
- ✅ **RBAC**: 基于角色的访问控制

---

## 📝 业务场景

### 1. 二手书交易

用户 A 发布二手书 → 用户 B 搜索购买 → 创建订单 → 支付 → 确认收货 → 评价

### 2. 物品租借

用户 A 发布自行车出租 → 用户 B 租借 → 创建租借订单 → 支付押金+租金 → 归还 → 退押金

### 3. 跑腿服务

用户 A 发布跑腿任务 → 用户 B 接单 → 完成任务 → 确认完成 → 支付报酬 → 评价

### 4. 代课服务

用户 A 发布代课任务 → 用户 B 接单 → 完成代签 → 上传凭证 → 支付报酬 → 评价

---

## 🎯 核心亮点

1. **完整的微服务架构**: 从认证、用户、商品、任务到订单、支付、搜索的完整链路
2. **Go 生态最佳实践**: Kratos 框架 + Wire 依赖注入 + gRPC 通信
3. **云原生设计**: Docker + Kubernetes + Consul + APISIX
4. **分布式事务**: DTM 保证跨服务数据一致性
5. **可观测性**: Prometheus + Grafana + Jaeger 全链路监控
6. **高性能**: gRPC + Redis 缓存 + Elasticsearch 搜索
7. **详细文档**: 从快速开始到API示例的完整文档

---

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

### 开发规范

1. 遵循 [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
2. 使用 `golangci-lint` 进行代码检查
3. 提交前运行 `make lint` 和 `make test`
4. 提交信息格式: `[模块] 简短描述`

---

## 📄 License

MIT License

---

## 🙏 致谢

本项目参考了以下优秀开源项目：

- [mall4cloud](https://github.com/gz-yami/mall4cloud) - Java 微服务商城
- [go-kratos/kratos](https://github.com/go-kratos/kratos) - Go 微服务框架
- [dtm-labs/dtm](https://github.com/dtm-labs/dtm) - 分布式事务管理器

---

**Happy Coding! 🎉**


