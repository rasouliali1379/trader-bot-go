package service

import (
	"go.uber.org/fx"
	"hamgit.ir/novin-backend/trader-bot/internal/core/service/strategies"
)

var Module = fx.Options(
	fx.Provide(strategies.NewEmaStrategy),
)
