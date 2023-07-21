package port

import (
	"context"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
)

type MarketJob interface {
	Run(ctx context.Context, m *domain.Market) error
}
