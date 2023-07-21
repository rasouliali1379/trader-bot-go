package job

import (
	"go.uber.org/fx"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/job/market"
)

var Module = fx.Options(
	fx.Provide(market.New),
)
