package data

import (
	"context"
	"database/sql"

	"entgo.io/ent/dialect"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	hwc "github.com/menta2l/go-hwc/api/hardware/v1"
	"github.com/menta2l/go-hwc/internal/conf"
	"github.com/menta2l/go-hwc/internal/data/ent"

	// init pgx driver
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewHardwareServiceClient, NewHardwareRepo)

// Data .
type Data struct {
	db  *ent.Client
	log *log.Helper
}

// NewData .

func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "server-service/data"))
	db, err := sql.Open("pgx", conf.Database.Source)
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))
	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		return nil, nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}
	d := &Data{
		db: client,
	}
	return d, func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}, nil

}

func NewHardwareServiceClient(ac *conf.Data) hwc.HardwareClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(ac.Server.Addr),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := hwc.NewHardwareClient(conn)
	return c
}
