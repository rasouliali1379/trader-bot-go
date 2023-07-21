package port

import (
	"context"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
)

type MarketJob interface {
	Run(c context.Context, lm *domain.Market) error
}
