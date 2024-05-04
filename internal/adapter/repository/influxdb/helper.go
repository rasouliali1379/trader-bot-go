package influxdb

import (
	"fmt"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"strings"
	"time"
)

func createOHLCFluxQuery(bucket string, exchange domain.Exchange, m *domain.Market, duration time.Duration) string {
	return fmt.Sprintf(`from(bucket: "%s")
  |> range(start: -%s)
  |> filter(fn: (r) => r["_measurement"] == "%s")
  |> filter(fn: (r) => r["_field"] == "close" or r["_field"] == "high" or r["_field"] == "low" or r["_field"] == "open")
  |> filter(fn: (r) => r["exchange"] == "%s%s")
  |> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
  |> yield(name: "mean")`,
		bucket,
		shortDuration(duration),
		exchange,
		m.Give,
		m.Take,
	)
}

func shortDuration(d time.Duration) string {
	s := d.String()
	if strings.HasSuffix(s, "m0s") {
		s = s[:len(s)-2]
	}
	if strings.HasSuffix(s, "h0m") {
		s = s[:len(s)-2]
	}
	return s
}
