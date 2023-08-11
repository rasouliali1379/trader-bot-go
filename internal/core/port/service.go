package port

import (
	"context"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
)

type StrategyService interface {
	Execute(c context.Context, exchange domain.Exchange, m *domain.Market) error
}

type ExchangeService interface {
	TrackMarket(c context.Context, m *domain.Market) error
	WatchMarkets(c context.Context)
}
