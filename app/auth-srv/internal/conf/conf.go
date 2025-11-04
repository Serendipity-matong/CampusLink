package conf

import "time"

// Bootstrap 启动配置
type Bootstrap struct {
	Server *Server
	Data   *Data
}

// Server 服务配置
type Server struct {
	HTTP *HTTP
	GRPC *GRPC
}

// HTTP 配置
type HTTP struct {
	Network string
	Addr    string
	Timeout time.Duration
}

// GRPC 配置
type GRPC struct {
	Network string
	Addr    string
	Timeout time.Duration
}

// Data 数据配置
type Data struct {
	Database *Database
	Redis    *Redis
	JWT      *JWT
}

// Database 数据库配置
type Database struct {
	Driver          string
	Source          string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

// Redis 配置
type Redis struct {
	Addr         string
	Password     string
	DB           int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// JWT 配置
type JWT struct {
	Secret     string
	Issuer     string
	Expiration time.Duration
}


