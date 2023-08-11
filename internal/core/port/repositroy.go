package port

import (
	"context"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"time"
)

type ExchangeRepository interface {
	Name() domain.Exchange
	GetBalance(c context.Context) error
	Subscribe(c context.Context, channel string, instrumentID string) error
	Unsubscribe(c context.Context, channel string, instrumentID string) error
	Read(c context.Context) (any, error)
	HasMarket(c context.Context, m *domain.Market) error
}

type InfluxRepository interface {
	AddPrice(ctx context.Context, exchange domain.Exchange, m *domain.Market)
	GetPrices(ctx context.Context, exchange domain.Exchange, m *domain.Market, period time.Duration) (*domain.Market, error)
}
