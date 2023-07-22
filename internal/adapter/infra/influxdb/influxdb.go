package influxdb

import (
	"context"
	"errors"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/domain"
	"hamgit.ir/novin-backend/trader-bot/config"
)

func Init() influxdb2.Client {
	return influxdb2.NewClientWithOptions(
		config.C().InfluxDB.Url,
		config.C().InfluxDB.Token,
		influxdb2.DefaultOptions().SetBatchSize(20))
}

func HealthCheck(client influxdb2.Client) error {
	health, err := client.Health(context.Background())
	if err != nil {
		return err
	}

	if health.Status == domain.HealthCheckStatusFail {
		return errors.New("influxdb connection is not healthy")
	}

	return nil
}
