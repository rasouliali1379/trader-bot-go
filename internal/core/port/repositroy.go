package port

import (
	"context"
)

type ExchangeRepository interface {
	Subscribe(c context.Context, channel string, instrumentID string) error
	Unsubscribe(c context.Context, channel string, instrumentID string) error
	Read(c context.Context) (any, error)
}

type InfluxRepository interface {
	AddPoint(ctx context.Context) error
}
