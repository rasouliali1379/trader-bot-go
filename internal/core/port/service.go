package port

import (
	"context"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
)

type StrategyService interface {
	Execute(c context.Context) error
}

type MarketService interface {
	SubscribeToMarket(c context.Context, m *domain.Market) error
	TrackMarket(c context.Context, m *domain.Price)
}
