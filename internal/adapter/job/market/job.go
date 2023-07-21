package market

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"hamgit.ir/novin-backend/trader-bot/config"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
	"time"
)

type Job struct {
}

type Params struct {
	fx.In
	Exchange *port.ExchangeRepository
}

func New(params Params) port.MarketJob {
	return &Job{}
}

func (j Job) Run(ctx context.Context, m *domain.Market) error {
	zap.L().Info("Market job is running")

	ticker := time.NewTicker(config.C().JobDuration.Market)
	for ; true; <-ticker.C {
		if err := j.process(ctx, m); err != nil {
			zap.L().Error("error while running job", zap.Error(err))
		}
	}

	return nil
}

func (j Job) process(ctx context.Context, m *domain.Market) error {
	return nil
}
