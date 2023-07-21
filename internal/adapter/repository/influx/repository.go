package influx

import (
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
)

type repository struct {
	db influxdb2.Client
}

func New(db influxdb2.Client) port.InfluxRepository {
	return &repository{db: db}
}

func (r repository) AddPoint(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
