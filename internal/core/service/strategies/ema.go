package strategies

import (
	"go.uber.org/fx"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/infra/exchange"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
)

type EmaStrategy struct {
}

type Params struct {
	fx.In
	Exchange *exchange.Exchange
}

func NewEmaStrategy(params Params) port.StrategyService {
	return &EmaStrategy{}
}
