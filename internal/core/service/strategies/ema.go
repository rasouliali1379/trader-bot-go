package strategies

import (
	"context"
	"go.uber.org/zap"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
)

type ema struct {
	exchangeRepo port.ExchangeRepository
}

func NewEmaStrategy(exchangeRepo port.ExchangeRepository) port.StrategyService {
	return &ema{exchangeRepo: exchangeRepo}
}

func (e ema) Execute(_ context.Context) error {
	zap.L().Info("executing ema strategy")
	return nil
}
