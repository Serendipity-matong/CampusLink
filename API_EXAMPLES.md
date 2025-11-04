# CampusLink API 使用示例

本文档提供 CampusLink 平台各个服务的 API 使用示例。

## 目录

- [认证服务 (auth-srv)](#认证服务-auth-srv)
- [用户服务 (user-srv)](#用户服务-user-srv)
- [商品服务 (product-srv)](#商品服务-product-srv)
- [任务服务 (task-srv)](#任务服务-task-srv)
- [订单服务 (order-srv)](#订单服务-order-srv)
- [支付服务 (payment-srv)](#支付服务-payment-srv)
- [搜索服务 (search-srv)](#搜索服务-search-srv)

## 通用说明

### 请求头

所有需要认证的API都需要在请求头中携带Token：

```http
Authorization: Bearer <access_token>
Content-Type: application/json
```

### 响应格式

所有API响应都遵循统一格式：

```json
{
  "code": 0,           // 0表示成功，非0表示错误码
  "message": "success", // 响应消息
  "data": {}           // 响应数据
}
```

### 分页响应

```json
{
  "code": 0,
  "message": "success",
  "data": [...],       // 数据列表
  "total": 100,        // 总数
  "page": 1,           // 当前页
  "size": 10           // 每页大小
}
```

---

## 认证服务 (auth-srv)

**基础URL**: `http://localhost:10001` (开发环境)  
**生产URL**: `http://api.campuslink.com/auth` (通过APISIX网关)

### 1. 用户注册

**接口**: `POST /api/v1/auth/register`

**请求示例**:

```bash
curl -X POST http://localhost:10001/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "student01",
    "password": "Abc123456!",
    "phone": "13800138888",
    "student_id": "2024001",
    "real_name": "张三",
    "school": "某某大学",
    "verification_code": "123456"
  }'
```

**响应示例**:

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "user_id": 1,
    "message": "注册成功"
  }
}
```

### 2. 用户登录

**接口**: `POST /api/v1/auth/login`

**请求示例**:

```bash
curl -X POST http://localhost:10001/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "student01",
    "password": "Abc123456!",
    "login_type": "username"
  }'
```

**多种登录方式**:

```javascript
// 用户名登录
{
  "username": "student01",
  "password": "Abc123456!",
  "login_type": "username"
}

// 手机号登录
{
  "username": "13800138888",
  "password": "Abc123456!",
  "login_type": "phone"
}

// 学号登录
{
  "username": "2024001",
  "password": "Abc123456!",
  "login_type": "student_id"
}
```

**响应示例**:

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_in": 7200,
    "user_info": {
      "user_id": 1,
      "username": "student01",
      "phone": "13800138888",
      "student_id": "2024001",
      "real_name": "张三",
      "school": "某某大学",
      "avatar": "https://example.com/avatar.jpg",
      "role": "student"
    }
  }
}
```

### 3. 刷新Token

**接口**: `POST /api/v1/auth/refresh-token`

**请求示例**:

```bash
curl -X POST http://localhost:10001/api/v1/auth/refresh-token \
  -H "Content-Type: application/json" \
  -d '{
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }'
```

### 4. 修改密码

**接口**: `POST /api/v1/auth/change-password`

**请求示例**:

```bash
curl -X POST http://localhost:10001/api/v1/auth/change-password \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <access_token>" \
  -d '{
    "user_id": 1,
    "old_password": "Abc123456!",
    "new_password": "NewPass123!"
  }'
```

### 5. 登出

**接口**: `POST /api/v1/auth/logout`

```bash
curl -X POST http://localhost:10001/api/v1/auth/logout \
  -H "Authorization: Bearer <access_token>" \
  -d '{
    "token": "<access_token>"
  }'
```

---

## 用户服务 (user-srv)

**基础URL**: `http://localhost:10002`

### 1. 获取用户信息

**接口**: `GET /api/v1/user/{user_id}`

```bash
curl -X GET http://localhost:10002/api/v1/user/1 \
  -H "Authorization: Bearer <access_token>"
```

**响应示例**:

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "user_id": 1,
    "username": "student01",
    "nickname": "小明",
    "phone": "13800138888",
    "email": "student01@example.com",
    "student_id": "2024001",
    "real_name": "张三",
    "school": "某某大学",
    "avatar": "https://example.com/avatar.jpg",
    "bio": "热爱学习的大学生",
    "role": "student",
    "verified": true,
    "credit_score": 100,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### 2. 更新用户信息

**接口**: `PUT /api/v1/user/{user_id}`

```bash
curl -X PUT http://localhost:10002/api/v1/user/1 \
  -H "Authorization: Bearer <access_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "nickname": "新昵称",
    "avatar": "https://example.com/new-avatar.jpg",
    "bio": "这是我的个人简介",
    "email": "newemail@example.com"
  }'
```

### 3. 学生认证

**接口**: `POST /api/v1/user/verify-student`

```bash
curl -X POST http://localhost:10002/api/v1/user/verify-student \
  -H "Authorization: Bearer <access_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "student_id": "2024001",
    "real_name": "张三",
    "school": "某某大学",
    "id_card": "110101199001011234",
    "student_card_image": "https://example.com/student-card.jpg"
  }'
```

### 4. 获取用户统计

**接口**: `GET /api/v1/user/{user_id}/stats`

```bash
curl -X GET http://localhost:10002/api/v1/user/1/stats \
  -H "Authorization: Bearer <access_token>"
```

**响应示例**:

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "order_count": 15,
    "task_count": 8,
    "product_count": 5,
    "credit_score": 100,
    "balance": 158.50
  }
}
```

---

## 商品服务 (product-srv)

**基础URL**: `http://localhost:10003`

### 1. 发布商品

**接口**: `POST /api/v1/product`

```bash
curl -X POST http://localhost:10003/api/v1/product \
  -H "Authorization: Bearer <access_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "高等数学教材（第七版）",
    "description": "九成新，无笔记，适合大一学生",
    "category": "book",
    "product_type": "sale",
    "price": 35.00,
    "stock": 1,
    "images": [
      "https://example.com/book1.jpg",
      "https://example.com/book2.jpg"
    ],
    "seller_id": 1,
    "condition": "like_new",
    "location": "东区宿舍楼3栋",
    "extra": {
      "author": "同济大学",
      "isbn": "9787040396638"
    }
  }'
```

### 2. 获取商品列表

**接口**: `GET /api/v1/products`

```bash
curl -X GET "http://localhost:10003/api/v1/products?page=1&page_size=10&category=book&product_type=sale&price_min=0&price_max=100&keyword=数学" \
  -H "Authorization: Bearer <access_token>"
```

**查询参数**:

- `page`: 页码 (默认: 1)
- `page_size`: 每页数量 (默认: 10, 最大: 100)
- `category`: 分类筛选 (book, electronics, sports, daily)
- `product_type`: 类型筛选 (sale, rent)
- `seller_id`: 卖家筛选
- `keyword`: 关键词搜索
- `price_min`: 最低价格
- `price_max`: 最高价格
- `condition`: 成色筛选 (new, like_new, good, fair)
- `sort_by`: 排序方式 (price_asc, price_desc, created_desc)

**响应示例**:

```json
{
  "code": 0,
  "message": "success",
  "data": [
    {
      "product_id": 1,
      "name": "高等数学教材（第七版）",
      "description": "九成新，无笔记，适合大一学生",
      "category": "book",
      "product_type": "sale",
      "price": 35.00,
      "images": ["https://example.com/book1.jpg"],
      "seller_id": 1,
      "seller_name": "张三",
      "condition": "like_new",
      "location": "东区宿舍楼3栋",
      "status": "online",
      "view_count": 125,
      "favorite_count": 8,
      "created_at": "2024-01-01T00:00:00Z"
    }
  ],
  "total": 1,
  "page": 1,
  "size": 10
}
```

### 3. 发布租借商品

**接口**: `POST /api/v1/product`

```bash
curl -X POST http://localhost:10003/api/v1/product \
  -H "Authorization: Bearer <access_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "山地自行车",
    "description": "捷安特山地车，适合校园代步",
    "category": "sports",
    "product_type": "rent",
    "price": 500.00,
    "rent_price_daily": 10.00,
    "stock": 1,
    "images": ["https://example.com/bike.jpg"],
    "seller_id": 1,
    "condition": "good",
    "location": "西区操场旁",
    "extra": {
      "brand": "捷安特",
      "model": "ATX660"
    }
  }'
```

### 4. 上架/下架商品

**上架**:

```bash
curl -X POST http://localhost:10003/api/v1/product/1/publish \
  -H "Authorization: Bearer <access_token>" \
  -d '{
    "seller_id": 1
  }'
```

**下架**:

```bash
curl -X POST http://localhost:10003/api/v1/product/1/unpublish \
  -H "Authorization: Bearer <access_token>" \
  -d '{
    "seller_id": 1
  }'
```

---

## 任务服务 (task-srv)

**基础URL**: `http://localhost:10004`

### 1. 发布跑腿任务

**接口**: `POST /api/v1/task`

```bash
curl -X POST http://localhost:10004/api/v1/task \
  -H "Authorization: Bearer <access_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "帮忙取快递",
    "description": "菜鸟驿站有个快递，帮忙取到东区3栋",
    "task_type": "errand",
    "reward": 5.00,
    "pickup_location": "南门菜鸟驿站",
    "delivery_location": "东区宿舍3栋201",
    "deadline": "2024-01-15T18:00:00Z",
    "publisher_id": 1,
    "contact_phone": "13800138888",
    "images": ["https://example.com/qr-code.jpg"],
    "extra": {
      "pickup_code": "8888"
    }
  }'
```

### 2. 发布代课任务

```bash
curl -X POST http://localhost:10004/api/v1/task \
  -H "Authorization: Bearer <access_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "代签到 - 高等数学",
    "description": "明天上午第二节课，教学楼A301，帮忙代签到",
    "task_type": "substitute",
    "reward": 20.00,
    "pickup_location": "教学楼A301",
    "delivery_location": "教学楼A301",
    "deadline": "2024-01-15T08:00:00Z",
    "publisher_id": 1,
    "contact_phone": "13800138888",
    "extra": {
      "course_name": "高等数学",
      "teacher": "李老师"
    }
  }'
```

### 3. 接受任务

**接口**: `POST /api/v1/task/{task_id}/accept`

```bash
curl -X POST http://localhost:10004/api/v1/task/1/accept \
  -H "Authorization: Bearer <access_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "executor_id": 2,
    "message": "我马上就去，10分钟到"
  }'
```

### 4. 完成任务

**接口**: `POST /api/v1/task/{task_id}/complete`

```bash
curl -X POST http://localhost:10004/api/v1/task/1/complete \
  -H "Authorization: Bearer <access_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "executor_id": 2,
    "completion_note": "已送达",
    "proof_images": ["https://example.com/proof.jpg"]
  }'
```

### 5. 评价任务

**接口**: `POST /api/v1/task/{task_id}/rate`

```bash
curl -X POST http://localhost:10004/api/v1/task/1/rate \
  -H "Authorization: Bearer <access_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "rater_id": 1,
    "rating": 5,
    "comment": "送得很快，服务态度好！"
  }'
```

---

## 订单服务 (order-srv)

**基础URL**: `http://localhost:10005`

### 1. 创建订单

**接口**: `POST /api/v1/order`

```bash
curl -X POST http://localhost:10005/api/v1/order \
  -H "Authorization: Bearer <access_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "buyer_id": 2,
    "seller_id": 1,
    "order_type": "product_sale",
    "items": [
      {
        "product_id": 1,
        "product_name": "高等数学教材（第七版）",
        "quantity": 1,
        "price": 35.00
      }
    ],
    "total_amount": 35.00,
    "delivery_address": "东区宿舍3栋201",
    "contact_phone": "13800138888",
    "remark": "请在宿舍楼下交易"
  }'
```

**响应示例**:

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "order_id": 1,
    "order_no": "O202401150001",
    "message": "订单创建成功"
  }
}
```

### 2. 创建租借订单

```bash
curl -X POST http://localhost:10005/api/v1/order \
  -H "Authorization: Bearer <access_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "buyer_id": 2,
    "seller_id": 1,
    "order_type": "product_rent",
    "items": [
      {
        "product_id": 2,
        "product_name": "山地自行车",
        "quantity": 1,
        "price": 10.00,
        "rent_days": 7
      }
    ],
    "total_amount": 70.00,
    "delivery_address": "东区宿舍3栋",
    "contact_phone": "13800138888",
    "remark": "租用7天，周日归还"
  }'
```

### 3. 查询订单

**接口**: `GET /api/v1/order/{order_id}`

```bash
curl -X GET http://localhost:10005/api/v1/order/1 \
  -H "Authorization: Bearer <access_token>"
```

### 4. 取消订单

**接口**: `POST /api/v1/order/{order_id}/cancel`

```bash
curl -X POST http://localhost:10005/api/v1/order/1/cancel \
  -H "Authorization: Bearer <access_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 2,
    "reason": "不想要了"
  }'
```

---

## 支付服务 (payment-srv)

**基础URL**: `http://localhost:10006`

### 1. 创建支付

**接口**: `POST /api/v1/payment`

```bash
curl -X POST http://localhost:10006/api/v1/payment \
  -H "Authorization: Bearer <access_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "order_id": 1,
    "user_id": 2,
    "amount": 35.00,
    "payment_method": "wechat",
    "subject": "购买二手书",
    "body": "高等数学教材"
  }'
```

**响应示例**:

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "payment_id": 1,
    "payment_no": "P202401150001",
    "payment_url": "https://wxpay.com/qrcode/xxx",
    "qr_code": "data:image/png;base64,iVBORw0KGg...",
    "message": "请扫码支付"
  }
}
```

### 2. 余额支付

```bash
curl -X POST http://localhost:10006/api/v1/payment \
  -H "Authorization: Bearer <access_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "order_id": 1,
    "user_id": 2,
    "amount": 35.00,
    "payment_method": "balance",
    "subject": "购买二手书"
  }'
```

### 3. 查询余额

**接口**: `GET /api/v1/payment/balance/{user_id}`

```bash
curl -X GET http://localhost:10006/api/v1/payment/balance/2 \
  -H "Authorization: Bearer <access_token>"
```

**响应示例**:

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "balance": 158.50,
    "frozen_balance": 35.00,
    "total_income": 500.00,
    "total_expense": 306.50
  }
}
```

---

## 搜索服务 (search-srv)

**基础URL**: `http://localhost:10007`

### 1. 搜索商品

**接口**: `GET /api/v1/search/products`

```bash
curl -X GET "http://localhost:10007/api/v1/search/products?keyword=数学&page=1&page_size=10&sort_by=relevance" \
  -H "Authorization: Bearer <access_token>"
```

### 2. 综合搜索

**接口**: `GET /api/v1/search/global`

```bash
curl -X GET "http://localhost:10007/api/v1/search/global?keyword=数学&page=1&page_size=10" \
  -H "Authorization: Bearer <access_token>"
```

### 3. 获取热搜词

**接口**: `GET /api/v1/search/hot-keywords`

```bash
curl -X GET "http://localhost:10007/api/v1/search/hot-keywords?limit=10" \
  -H "Authorization: Bearer <access_token>"
```

**响应示例**:

```json
{
  "code": 0,
  "message": "success",
  "data": [
    {
      "keyword": "高等数学",
      "search_count": 1258,
      "rank": 1
    },
    {
      "keyword": "自行车",
      "search_count": 856,
      "rank": 2
    }
  ]
}
```

---

## 错误码说明

| 错误码 | 说明 |
|-------|------|
| 0 | 成功 |
| 1000 | 内部服务错误 |
| 1001 | 参数错误 |
| 1002 | 资源不存在 |
| 1003 | 资源已存在 |
| 1004 | 权限不足 |
| 1005 | 未授权 |
| 2001 | 用户不存在 |
| 2003 | 密码错误 |
| 2004 | Token无效 |
| 2005 | Token已过期 |
| 3001 | 商品不存在 |
| 3002 | 商品已下架 |
| 3003 | 库存不足 |
| 4001 | 任务不存在 |
| 5001 | 订单不存在 |
| 6001 | 支付失败 |

## 完整的业务流程示例

### 场景: 购买二手书

```bash
# 1. 用户登录
ACCESS_TOKEN=$(curl -X POST http://localhost:10001/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"buyer01","password":"Abc123456!","login_type":"username"}' \
  | jq -r '.data.access_token')

# 2. 搜索商品
curl -X GET "http://localhost:10007/api/v1/search/products?keyword=数学" \
  -H "Authorization: Bearer $ACCESS_TOKEN"

# 3. 查看商品详情
curl -X GET http://localhost:10003/api/v1/product/1 \
  -H "Authorization: Bearer $ACCESS_TOKEN"

# 4. 创建订单
ORDER_ID=$(curl -X POST http://localhost:10005/api/v1/order \
  -H "Authorization: Bearer $ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"buyer_id":2,"seller_id":1,"order_type":"product_sale","items":[{"product_id":1,"product_name":"高等数学","quantity":1,"price":35.00}],"total_amount":35.00}' \
  | jq -r '.data.order_id')

# 5. 创建支付
curl -X POST http://localhost:10006/api/v1/payment \
  -H "Authorization: Bearer $ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d "{\"order_id\":$ORDER_ID,\"user_id\":2,\"amount\":35.00,\"payment_method\":\"balance\"}"

# 6. 确认收货
curl -X POST http://localhost:10005/api/v1/order/$ORDER_ID/complete \
  -H "Authorization: Bearer $ACCESS_TOKEN" \
  -d '{"buyer_id":2}'

# 7. 评价订单
curl -X POST http://localhost:10005/api/v1/order/$ORDER_ID/rate \
  -H "Authorization: Bearer $ACCESS_TOKEN" \
  -d '{"rater_id":2,"rating":5,"comment":"质量很好！"}'
```

## 调试工具推荐

- **Postman**: 导入项目根目录的 `postman_collection.json`
- **Insomnia**: API调试工具
- **grpcurl**: gRPC命令行调试工具
- **BloomRPC**: gRPC图形化调试工具


