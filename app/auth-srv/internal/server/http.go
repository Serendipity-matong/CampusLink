package server

import (
	v1 "github.com/campuslink/campuslink/api/auth/v1"
	"github.com/campuslink/campuslink/app/auth-srv/internal/conf"
	"github.com/campuslink/campuslink/app/auth-srv/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer 创建 HTTP Server.
func NewHTTPServer(c *conf.Server, authService *service.AuthService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.HTTP.Network != "" {
		opts = append(opts, http.Network(c.HTTP.Network))
	}
	if c.HTTP.Addr != "" {
		opts = append(opts, http.Address(c.HTTP.Addr))
	}
	if c.HTTP.Timeout != 0 {
		opts = append(opts, http.Timeout(c.HTTP.Timeout))
	}
	srv := http.NewServer(opts...)
	// TODO: 需要先生成 HTTP proto 代码
	// v1.RegisterAuthHTTPServer(srv, authService)
	_ = v1.UnimplementedAuthServer{}  // 避免未使用导入警告
	return srv
}

