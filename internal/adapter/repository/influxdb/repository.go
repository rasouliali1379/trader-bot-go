package influxdb

import (
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	influxapi "github.com/influxdata/influxdb-client-go/v2/api"
	"hamgit.ir/novin-backend/trader-bot/config"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
	"time"
)

type repository struct {
	write influxapi.WriteAPI
	read  influxapi.QueryAPI
}

func New(write influxapi.WriteAPI, read influxapi.QueryAPI) port.InfluxRepository {
	return &repository{write: write, read: read}
}

func (r *repository) AddPrice(_ context.Context, exchange domain.Exchange, m *domain.Market) {
	for i := range m.Price.Candles {
		measure := m.Give + m.Take
		p := influxdb2.NewPointWithMeasurement(string(exchange)).
			AddTag("exchange", measure).
			AddField("open", m.Price.Candles[i].Open).
			AddField("close", m.Price.Candles[i].Close).
			AddField("high", m.Price.Candles[i].High).
			AddField("low", m.Price.Candles[i].Low).
			SetTime(m.Price.Candles[i].Time)
		r.write.WritePoint(p)
	}
	r.write.Flush()
}

func (r *repository) GetPrices(c context.Context, exchange domain.Exchange, m *domain.Market, period time.Duration) (*domain.Market, error) {
	query := createOHLCFluxQuery(config.C().InfluxDB.Bucket, exchange, m, period)
	result, err := r.read.QueryRaw(c, query, influxdb2.DefaultDialect())
	if err != nil {
		return nil, err
	}

	m.Price.ParseFromInfluxDto(result)
	return m, nil
}
