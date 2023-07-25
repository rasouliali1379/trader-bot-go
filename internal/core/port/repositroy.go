package port

import (
	"context"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
)

type ExchangeRepository interface {
	Subscribe(c context.Context, channel string, instrumentID string) error
	Unsubscribe(c context.Context, channel string, instrumentID string) error
	Read(c context.Context) (any, error)
}

type InfluxRepository interface {
	AddPoint(ctx context.Context, m *domain.Price)
}
