//go:build wireinject
// +build wireinject

package main

import (
	"github.com/campuslink/campuslink/app/product-srv/internal/biz"
	"github.com/campuslink/campuslink/app/product-srv/internal/conf"
	"github.com/campuslink/campuslink/app/product-srv/internal/data"
	"github.com/campuslink/campuslink/app/product-srv/internal/server"
	"github.com/campuslink/campuslink/app/product-srv/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp 使用 Wire 进行依赖注入
func wireApp(*conf.Server, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		newApp,
	))
}
