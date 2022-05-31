// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/menta2l/go-hwc/internal/biz"
	"github.com/menta2l/go-hwc/internal/conf"
	"github.com/menta2l/go-hwc/internal/data"
	"github.com/menta2l/go-hwc/internal/server"
	"github.com/menta2l/go-hwc/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
