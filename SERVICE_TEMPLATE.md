# 微服务创建模板

本文档说明如何基于 `auth-srv` 模板创建新的微服务。

## 方法1: 使用 Kratos CLI (推荐)

### 1. 安装 Kratos CLI

```bash
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```

### 2. 创建新服务

```bash
cd app
kratos new user-srv
```

### 3. 创建 API 定义

```bash
cd user-srv
kratos proto add api/user/v1/user.proto
```

编辑 proto 文件后，生成代码：

```bash
kratos proto client api/user/v1/user.proto
```

### 4. 实现业务逻辑

参考 `auth-srv` 的结构实现 `internal/biz`、`internal/data`、`internal/service`

### 5. 生成 Wire 依赖注入

```bash
cd cmd/server
wire
```

## 方法2: 复制 auth-srv 模板 (快速)

### 1. 复制目录结构

```bash
# 复制整个 auth-srv 目录
cp -r app/auth-srv app/user-srv
```

### 2. 修改 go.mod

编辑 `app/user-srv/go.mod`:

```go
module github.com/campuslink/campuslink/app/user-srv

go 1.21

require (
	github.com/campuslink/campuslink v0.0.0
	github.com/go-kratos/kratos/v2 v2.7.0
	github.com/google/wire v0.5.0
	google.golang.org/grpc v1.59.0
	google.golang.org/protobuf v1.31.0
	gorm.io/gorm v1.25.5
)

replace github.com/campuslink/campuslink => ../..
```

### 3. 修改配置文件

创建 `configs/user-srv.yaml`:

```yaml
server:
  http:
    network: tcp
    addr: 0.0.0.0:10002  # ⚠️ 修改端口
    timeout: 10s
  grpc:
    network: tcp
    addr: 0.0.0.0:9002   # ⚠️ 修改端口
    timeout: 10s

data:
  database:
    driver: mysql
    source: root:password@tcp(127.0.0.1:3306)/campuslink?charset=utf8mb4&parseTime=True&loc=Local
    max_idle_conns: 10
    max_open_conns: 100
    conn_max_lifetime: 3600s
  redis:
    addr: 127.0.0.1:6379
    password: ""
    db: 0
    dial_timeout: 5s
    read_timeout: 3s
    write_timeout: 3s
```

### 4. 修改 main.go

编辑 `app/user-srv/cmd/server/main.go`:

```go
var (
	Name    = "user-srv"  // ⚠️ 修改服务名
	Version = "v1.0.0"

	flagconf string
	id, _    = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs/user-srv.yaml", "config path")  // ⚠️ 修改配置路径
}
```

### 5. 修改业务逻辑

#### 5.1 修改 Data 层 (internal/data/)

**data.go**:

```go
// 修改数据库模型
if err := db.AutoMigrate(&User{}); err != nil {  // ⚠️ 修改为实际模型
    panic(err)
}
```

**user.go** (或对应的数据文件):

```go
// 定义 PO (Persistent Object)
type User struct {
    ID        uint           `gorm:"primarykey"`
    Username  string         `gorm:"uniqueIndex;size:50;not null"`
    // ... 其他字段
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (User) TableName() string {
    return "user"  // ⚠️ 修改表名
}

// 定义 Repository 接口
type userRepo struct {
    data *Data
    log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
    return &userRepo{
        data: data,
        log:  log.NewHelper(logger),
    }
}

// 实现具体方法
func (r *userRepo) GetUser(ctx context.Context, id uint) (*biz.User, error) {
    var po User
    if err := r.data.db.WithContext(ctx).Where("id = ?", id).First(&po).Error; err != nil {
        return nil, err
    }
    return r.poToBiz(&po), nil
}
```

#### 5.2 修改 Biz 层 (internal/biz/)

**biz.go**:

```go
var ProviderSet = wire.NewSet(NewUserUsecase)  // ⚠️ 修改 Usecase
```

**user.go** (或对应的业务文件):

```go
// 定义 BO (Business Object)
type User struct {
    ID        uint
    Username  string
    // ... 业务字段
}

// 定义 Repository 接口
type UserRepo interface {
    GetUser(ctx context.Context, id uint) (*User, error)
    CreateUser(ctx context.Context, user *User) (uint64, error)
    UpdateUser(ctx context.Context, user *User) error
}

// 定义 Usecase
type UserUsecase struct {
    repo UserRepo
    log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
    return &UserUsecase{
        repo: repo,
        log:  log.NewHelper(logger),
    }
}

// 实现业务逻辑
func (uc *UserUsecase) GetUser(ctx context.Context, id uint) (*User, error) {
    user, err := uc.repo.GetUser(ctx, id)
    if err != nil {
        return nil, errors.ErrUserNotFound
    }
    return user, nil
}
```

#### 5.3 修改 Service 层 (internal/service/)

**service.go**:

```go
var ProviderSet = wire.NewSet(NewUserService)  // ⚠️ 修改 Service
```

**user.go** (或对应的服务文件):

```go
import (
    pb "github.com/campuslink/campuslink/api/user/v1"  // ⚠️ 修改导入
    "github.com/campuslink/campuslink/app/user-srv/internal/biz"
)

type UserService struct {
    pb.UnimplementedUserServer  // ⚠️ 修改为对应的 Server

    userUsecase *biz.UserUsecase
    log         *log.Helper
}

func NewUserService(userUsecase *biz.UserUsecase, logger log.Logger) *UserService {
    return &UserService{
        userUsecase: userUsecase,
        log:         log.NewHelper(logger),
    }
}

// 实现 gRPC 接口
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
    s.log.WithContext(ctx).Infof("GetUser: %v", req.UserId)
    
    user, err := s.userUsecase.GetUser(ctx, uint(req.UserId))
    if err != nil {
        return nil, err
    }

    return &pb.GetUserReply{
        User: &pb.UserInfo{
            UserId:   uint64(user.ID),
            Username: user.Username,
            // ... 其他字段
        },
    }, nil
}
```

#### 5.4 修改 Server 层 (internal/server/)

**grpc.go**:

```go
import (
    v1 "github.com/campuslink/campuslink/api/user/v1"  // ⚠️ 修改导入
)

func NewGRPCServer(c *conf.Server, userService *service.UserService, logger log.Logger) *grpc.Server {
    // ... 配置不变
    v1.RegisterUserServer(srv, userService)  // ⚠️ 注册对应的服务
    return srv
}
```

**http.go**:

```go
import (
    v1 "github.com/campuslink/campuslink/api/user/v1"  // ⚠️ 修改导入
)

func NewHTTPServer(c *conf.Server, userService *service.UserService, logger log.Logger) *http.Server {
    // ... 配置不变
    v1.RegisterUserHTTPServer(srv, userService)  // ⚠️ 注册对应的服务
    return srv
}
```

### 6. 重新生成 Wire 代码

```bash
cd app/user-srv/cmd/server
wire
```

### 7. 运行服务

```bash
go run cmd/server/main.go -conf ../../configs/user-srv.yaml
```

---

## 各服务端口分配

| 服务 | gRPC端口 | HTTP端口 |
|------|----------|----------|
| auth-srv | 9001 | 10001 |
| user-srv | 9002 | 10002 |
| product-srv | 9003 | 10003 |
| task-srv | 9004 | 10004 |
| order-srv | 9005 | 10005 |
| payment-srv | 9006 | 10006 |
| search-srv | 9007 | 10007 |
| admin-srv | 9008 | 10008 |
| biz-srv | 9009 | 10009 |

---

## 服务间调用示例

### 在 order-srv 中调用 user-srv

**1. 在 data.go 中注入 gRPC Client:**

```go
import (
    userv1 "github.com/campuslink/campuslink/api/user/v1"
)

type Data struct {
    db         *gorm.DB
    rdb        *redis.Client
    userClient userv1.UserClient  // 注入 user-srv 客户端
}

func NewData(db *gorm.DB, rdb *redis.Client, userClient userv1.UserClient) (*Data, func(), error) {
    // ...
}

// 创建 gRPC Client
func NewUserClient() userv1.UserClient {
    conn, err := grpc.Dial("localhost:9002", grpc.WithTransInsecure())
    if err != nil {
        panic(err)
    }
    return userv1.NewUserClient(conn)
}
```

**2. 在 wire.go 中添加 Provider:**

```go
func wireApp(...) (*kratos.App, func(), error) {
    panic(wire.Build(
        // ... 其他 Provider
        data.NewUserClient,  // 添加这一行
    ))
}
```

**3. 在业务逻辑中使用:**

```go
type orderRepo struct {
    data *Data
}

func (r *orderRepo) CreateOrder(ctx context.Context, order *biz.Order) error {
    // 调用 user-srv 获取用户信息
    userResp, err := r.data.userClient.GetUser(ctx, &userv1.GetUserRequest{
        UserId: order.BuyerID,
    })
    if err != nil {
        return err
    }
    
    // 使用用户信息
    buyerName := userResp.User.Username
    // ...
}
```

---

## 常见问题

### Q1: Wire 生成失败？

检查 `wire.go` 文件中的 ProviderSet 是否完整：

```go
panic(wire.Build(
    server.ProviderSet,   // ✅
    data.ProviderSet,     // ✅
    biz.ProviderSet,      // ✅
    service.ProviderSet,  // ✅
    newApp,               // ✅
))
```

### Q2: gRPC 注册失败？

确保在 `server/grpc.go` 和 `server/http.go` 中正确注册了服务：

```go
v1.RegisterUserServer(srv, userService)      // gRPC
v1.RegisterUserHTTPServer(srv, userService)  // HTTP
```

### Q3: 数据库表未创建？

检查 `data/data.go` 中是否包含了 AutoMigrate：

```go
if err := db.AutoMigrate(&User{}, &UserStats{}); err != nil {
    panic(err)
}
```

---

## 检查清单

创建新服务时，请确认以下项目：

- [ ] ✅ 修改 `go.mod` 的模块名
- [ ] ✅ 创建对应的 `configs/xxx-srv.yaml` 配置文件
- [ ] ✅ 修改 `cmd/server/main.go` 中的服务名和配置路径
- [ ] ✅ 更新 `internal/data/data.go` 的数据库模型
- [ ] ✅ 实现 `internal/data/` 的 Repository
- [ ] ✅ 实现 `internal/biz/` 的 Usecase
- [ ] ✅ 实现 `internal/service/` 的 gRPC 接口
- [ ] ✅ 更新 `internal/server/grpc.go` 和 `http.go` 的注册代码
- [ ] ✅ 运行 `wire` 生成依赖注入代码
- [ ] ✅ 运行服务并测试

---

**祝您开发顺利！🚀**


