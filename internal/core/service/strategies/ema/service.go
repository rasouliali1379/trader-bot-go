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
	influxRepo port.InfluxRepository
}

func New(
	influxRepo port.InfluxRepository,
) port.StrategyService {
	return &service{influxRepo: influxRepo}
}

func (e service) Execute(c context.Context, m *domain.Market, r port.ExchangeRepository) error {
	zap.L().Info("Executing EMA strategy")
	defer zap.L().Info("EMA strategy executed successfully")

	market, err := e.influxRepo.GetPrices(c, m, time.Minute*60)
	if err != nil {
		return err
	}

	if len(market.Price.Candles) > 21 {
		log.Println(market.Price.Ema(21))
		log.Println(market.Price.Ema(8))
	}

	return nil
}
