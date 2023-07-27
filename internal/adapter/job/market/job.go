package market

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
	"log"
)

type job struct {
	marketService port.MarketService
	exchangeRepo  port.ExchangeRepository
}

func New(marketService port.MarketService, exchangeRepo port.ExchangeRepository) port.MarketJob {
	return &job{marketService: marketService, exchangeRepo: exchangeRepo}
}

func (j *job) Run(c context.Context, m *domain.Market) error {
	zap.L().Info("Market job is running")

	if err := j.marketService.SubscribeToMarket(c, m); err != nil {
		return err
	}

	zap.L().Info("successfully subscribed on channel",
		zap.String("give", m.Give),
		zap.String("take", m.Take))

	go j.watch(c, m)

	return nil
}

func (j *job) watch(c context.Context, m *domain.Market) {

	for {

		msg, err := j.exchangeRepo.Read(c)
		if err != nil {
			zap.L().Error(fmt.Sprintf("error while tracing %s-%s market", m.Give, m.Take), zap.Error(err))
			continue
		}

		log.Println(msg)

		switch msg.(type) {
		case *domain.Price:
			price := msg.(*domain.Price)
			price.Exchange.Market = m
			j.marketService.TrackMarket(c, price)
		}
	}
}
