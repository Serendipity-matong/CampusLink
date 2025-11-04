.PHONY: api build test lint docker-build deploy clean

# 生成所有 Proto 代码
api:
	@echo "Generating API code..."
	@cd api/auth/v1 && protoc --go_out=. --go-grpc_out=. --go-http_out=. auth.proto
	@cd api/user/v1 && protoc --go_out=. --go-grpc_out=. --go-http_out=. user.proto
	@cd api/product/v1 && protoc --go_out=. --go-grpc_out=. --go-http_out=. product.proto
	@cd api/task/v1 && protoc --go_out=. --go-grpc_out=. --go-http_out=. task.proto
	@cd api/order/v1 && protoc --go_out=. --go-grpc_out=. --go-http_out=. order.proto
	@cd api/payment/v1 && protoc --go_out=. --go-grpc_out=. --go-http_out=. payment.proto
	@cd api/search/v1 && protoc --go_out=. --go-grpc_out=. --go-http_out=. search.proto
	@cd api/admin/v1 && protoc --go_out=. --go-grpc_out=. --go-http_out=. admin.proto
	@cd api/biz/v1 && protoc --go_out=. --go-grpc_out=. --go-http_out=. biz.proto

# 生成 Wire 依赖注入
generate:
	@echo "Generating wire code..."
	@go generate ./...

# 编译所有服务
build:
	@echo "Building all services..."
	@cd app/auth-srv && go build -o ../../bin/auth-srv cmd/server/main.go
	@cd app/user-srv && go build -o ../../bin/user-srv cmd/server/main.go
	@cd app/product-srv && go build -o ../../bin/product-srv cmd/server/main.go
	@cd app/task-srv && go build -o ../../bin/task-srv cmd/server/main.go
	@cd app/order-srv && go build -o ../../bin/order-srv cmd/server/main.go
	@cd app/payment-srv && go build -o ../../bin/payment-srv cmd/server/main.go
	@cd app/search-srv && go build -o ../../bin/search-srv cmd/server/main.go
	@cd app/admin-srv && go build -o ../../bin/admin-srv cmd/server/main.go
	@cd app/biz-srv && go build -o ../../bin/biz-srv cmd/server/main.go

# 运行测试
test:
	@echo "Running tests..."
	@go test -v -cover ./...

# 代码检查
lint:
	@echo "Running linters..."
	@golangci-lint run ./...

# 构建 Docker 镜像
docker-build:
	@echo "Building Docker images..."
	@docker-compose build

# 部署到 Kubernetes
deploy:
	@echo "Deploying to Kubernetes..."
	@kubectl apply -f deployments/k8s/

# 清理
clean:
	@echo "Cleaning..."
	@rm -rf bin/
	@go clean -cache

# 启动所有服务 (本地开发)
run-all:
	@echo "Starting all services..."
	@docker-compose up -d

# 停止所有服务
stop-all:
	@echo "Stopping all services..."
	@docker-compose down


