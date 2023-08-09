package strategies

import (
	"go.uber.org/zap"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
	"hamgit.ir/novin-backend/trader-bot/internal/core/service/strategies/ema"
)

func New(s domain.Strategy, i port.InfluxRepository, r port.ExchangeRepository) port.StrategyService {
	switch s {
	case domain.Ema:
		return ema.New(i, r)
	default:
		zap.L().Fatal("unknown strategy", zap.String("strategy", string(s)))
	}

	return nil
}
