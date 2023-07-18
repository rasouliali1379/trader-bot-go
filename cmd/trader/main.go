package main

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"hamgit.ir/novin-backend/trader-bot/config"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/infra/exchange"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/infra/log"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/repository"
	"hamgit.ir/novin-backend/trader-bot/internal/core/service"
)

func main() {

	config.Init()

	app := fx.New(
		fx.Provide(exchange.Init),

		repository.Module,
		service.Module,

		fx.Invoke(log.Init),
	)

	zap.L().Fatal("error while starting the app", zap.Error(app.Start(context.Background())))
}
