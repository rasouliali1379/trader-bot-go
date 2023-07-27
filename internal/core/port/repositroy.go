package port

import (
	"context"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"time"
)

type ExchangeRepository interface {
	Subscribe(c context.Context, channel string, instrumentID string) error
	Unsubscribe(c context.Context, channel string, instrumentID string) error
	Read(c context.Context) (any, error)
}

type InfluxRepository interface {
	AddPrice(ctx context.Context, m *domain.Market)
	GetPrices(ctx context.Context, m *domain.Market, period time.Duration) (*domain.Market, error)
}
