package influxdb

import (
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
)

type repository struct {
	db influxdb2.Client
}

func New(db influxdb2.Client) port.InfluxRepository {
	return &repository{db: db}
}

func (r repository) AddPoint(ctx context.Context, m domain.Price) error {
	writeAPI := r.db.WriteAPI(org, bucket)
	for i := range m.List {
		p := influxdb2.NewPointWithMeasurement("price").
			AddTag("unit", m.Market.Take).
			AddField("open", m.List[i].Open).
			AddField("close", m.List[i].Close).
			AddField("high", m.List[i].High).
			AddField("low", m.List[i].Low).
			SetTime(m.List[i].Time)
		writeAPI.WritePoint(p)
	}

	return nil
}
