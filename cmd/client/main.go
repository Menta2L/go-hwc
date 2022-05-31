package main

import (
	"flag"
	"os"

	"github.com/menta2l/go-hwc/internal/conf"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/northseadl/gopkg/clog"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Name     string
	Version  string
	id, _    = os.Hostname()
	flagConf string
)

func init() {
	flag.StringVar(&flagConf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()
	serverLogger := zerolog.New(&lumberjack.Logger{
		Filename:   "logs/worker.log",
		MaxSize:    4,
		MaxBackups: 7,
		MaxAge:     28,
		Compress:   false,
	})
	logger := log.With(clog.NewKratosLogger(&serverLogger, true),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name+"-worker",
		"service.version", Version,
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagConf),
		),
		config.WithLogger(logger),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	iWorker, cleanup, err := wireWorker(bc.Data, logger)
	defer cleanup()
	if err != nil {
		panic(err)
	}
	done := make(chan struct{}, 1)
	iWorker.Run(done)
	<-done
}
