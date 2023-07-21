package port

import (
	"context"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
)

type StrategyService interface {
}

type MarketService interface {
	SubscribeToMarket(c context.Context, m *domain.Market) error
	TrackMarket(c context.Context, m *domain.Price) error
}
