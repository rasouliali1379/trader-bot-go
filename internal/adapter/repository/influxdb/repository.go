package influxdb

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	influxapi "github.com/influxdata/influxdb-client-go/v2/api"
	"hamgit.ir/novin-backend/trader-bot/config"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
	"log"
	"time"
)

type repository struct {
	write influxapi.WriteAPI
	read  influxapi.QueryAPI
}

func New(write influxapi.WriteAPI, read influxapi.QueryAPI) port.InfluxRepository {
	return &repository{write: write, read: read}
}

func (r *repository) AddPrice(_ context.Context, m *domain.Market) {
	for i := range m.Price.List {
		measure := m.Give + m.Take
		p := influxdb2.NewPointWithMeasurement(m.Exchange.Name).
			AddTag("market", measure).
			AddField("open", m.Price.List[i].Open).
			AddField("close", m.Price.List[i].Close).
			AddField("high", m.Price.List[i].High).
			AddField("low", m.Price.List[i].Low).
			SetTime(m.Price.List[i].Time)
		r.write.WritePoint(p)
	}
	r.write.Flush()
}

func (r *repository) GetPrices(_ context.Context, m *domain.Market, period time.Duration) (*domain.Market, error) {

	query := createOHLCFluxQuery(config.C().InfluxDB.Bucket, m, period)
	result, err := r.read.QueryRaw(context.Background(), query, influxdb2.DefaultDialect())
	if err != nil {
		return nil, err
	}

	log.Println(result)

	return nil, nil
}

func createOHLCFluxQuery(bucket string, m *domain.Market, duration time.Duration) string {
	return fmt.Sprintf(`from(bucket: "%s")
  |> range(start: -%s)
  |> filter(fn: (r) => r["_measurement"] == "%s")
  |> filter(fn: (r) => r["_field"] == "close" or r["_field"] == "high" or r["_field"] == "low" or r["_field"] == "open")
  |> filter(fn: (r) => r["market"] == "%s-%s")
  |> yield(name: "mean")`,
		bucket,
		duration.String(),
		m.Exchange.Name,
		m.Give,
		m.Take,
	)
}
