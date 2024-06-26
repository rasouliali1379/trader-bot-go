package ema

import (
	"context"
	"go.uber.org/zap"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
	"log"
	"time"
)

type service struct {
	influxRepo   port.InfluxRepository
	exchangeRepo port.ExchangeRepository
}

func New(
	influxRepo port.InfluxRepository,
	exchangeRepo port.ExchangeRepository,
) port.StrategyService {
	return &service{influxRepo: influxRepo, exchangeRepo: exchangeRepo}
}

func (e service) Execute(c context.Context, exchange domain.Exchange, m *domain.Market) error {
	zap.L().Info("Executing EMA strategy")
	defer zap.L().Info("EMA strategy executed successfully")

	market, err := e.influxRepo.GetPrices(c, exchange, m, time.Minute*120)
	if err != nil {
		return err
	}

	log.Println(market.Price.Prices)

	if market.Price.Prices.Nrow() > 21 {
		log.Println(market.Price.Ema(21))
		log.Println(market.Price.Ema(8))
	}

	return nil
}
