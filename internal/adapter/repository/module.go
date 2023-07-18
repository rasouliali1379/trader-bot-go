package repository

import (
	"go.uber.org/fx"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/repository/exchanges/okx"
)

var Module = fx.Options(
	fx.Provide(okx.New),
)
