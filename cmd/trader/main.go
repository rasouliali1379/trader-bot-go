package main

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"hamgit.ir/novin-backend/trader-bot/config"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/infra/exchange"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/infra/log"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/job"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/repository"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
	"hamgit.ir/novin-backend/trader-bot/internal/core/service"
)

func main() {

	config.Init()

	app := fx.New(
		fx.Provide(exchange.Init),

		repository.Module,
		service.Module,
		job.Module,

		fx.Invoke(log.Init),
		fx.Invoke(runJobs),
	)

	zap.L().Fatal("error while starting the app", zap.Error(app.Start(context.Background())))
}

func runJobs(lc fx.Lifecycle, marketJob port.MarketJob) {
	lc.Append(fx.Hook{
		OnStart: func(c context.Context) error {

			for _, s := range config.C().Strategies {
				for _, _ = range s.Markets {
					if err := marketJob.Run(c, &domain.Market{}); err != nil {
						return err
					}
				}
			}

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
