//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/menta2l/go-hwc/internal/biz"
	"github.com/menta2l/go-hwc/internal/conf"
	"github.com/menta2l/go-hwc/internal/data"
	"github.com/menta2l/go-hwc/internal/worker"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func wireWorker(confData *conf.Data, logger log.Logger) (*worker.Worker, func(), error) {
	panic(wire.Build(biz.WorkerProviderSet, data.ProviderSet, worker.ProviderSet))
}
