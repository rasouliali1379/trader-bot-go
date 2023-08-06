package strategies

import (
	"go.uber.org/zap"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
	"hamgit.ir/novin-backend/trader-bot/internal/core/service/strategies/ema"
)

func New(s domain.Strategy, r port.ExchangeRepository, i port.InfluxRepository) port.StrategyService {
	switch s {
	case domain.Ema:
		return ema.New(r, i)
	default:
		zap.L().Fatal("unknown strategy", zap.String("strategy", string(s)))
	}

	return nil
}
