package market

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"hamgit.ir/novin-backend/trader-bot/config"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
	"time"
)

type job struct {
	marketService port.MarketService
	exchangeRepo  port.ExchangeRepository
}

func New(marketService port.MarketService, exchangeRepo port.ExchangeRepository) port.MarketJob {
	return &job{marketService: marketService, exchangeRepo: exchangeRepo}
}

func (j job) Run(c context.Context, m *domain.Market) error {
	zap.L().Info("Market job is running")

	if err := j.marketService.SubscribeToMarket(c, m); err != nil {
		return err
	}

	go j.watch(c, m)

	return nil
}

func (j job) watch(c context.Context, m *domain.Market) {
	ticker := time.NewTicker(config.C().JobDuration.Market)

	for ; true; <-ticker.C {

		msg, err := j.exchangeRepo.Read(c)
		if err != nil {
			zap.L().Error(fmt.Sprintf("error while tracing %s-%s market", m.Give, m.Take), zap.Error(err))
		}

		switch msg.(type) {
		case *domain.Price:
			price := msg.(*domain.Price)
			price.Market = m
			if err := j.marketService.TrackMarket(c, price); err != nil {
				zap.L().Error("", zap.Error(err))
			}
		}
	}
}
