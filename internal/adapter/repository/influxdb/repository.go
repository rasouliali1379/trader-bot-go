package influxdb

import (
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	influxapi "github.com/influxdata/influxdb-client-go/v2/api"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
)

type repository struct {
	api influxapi.WriteAPI
}

func New(api influxapi.WriteAPI) port.InfluxRepository {
	return &repository{api: api}
}

func (r *repository) AddPoint(_ context.Context, m *domain.Price) {
	for i := range m.List {
		measure := m.Market.Give + m.Market.Take
		p := influxdb2.NewPointWithMeasurement(m.Market.Exchange).
			AddTag("market", measure).
			AddField("open", m.List[i].Open).
			AddField("close", m.List[i].Close).
			AddField("high", m.List[i].High).
			AddField("low", m.List[i].Low).
			SetTime(m.List[i].Time)
		r.api.WritePoint(p)
	}
	r.api.Flush()
}
