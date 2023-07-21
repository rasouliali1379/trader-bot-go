package influxdb

import (
	influxdb2 "github.com/influxdata/influx-client-go/v2"
	"hamgit.ir/novin-backend/trader-bot/config"
)

func Init() influxdb2.Client {
	return influxdb2.NewClient(config.C().InfluxDB.Url, config.C().InfluxDB.Token)
}
