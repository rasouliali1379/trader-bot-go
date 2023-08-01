package strategies

import (
	"context"
	"go.uber.org/zap"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
	"log"
	"time"
)

type ema struct {
	market       *domain.Market
	exchangeRepo port.ExchangeRepository
	influxRepo   port.InfluxRepository
}

func NewEmaStrategy(
	market *domain.Market,
	exchangeRepo port.ExchangeRepository,
	influxRepo port.InfluxRepository,
) port.StrategyService {
	return &ema{market: market, exchangeRepo: exchangeRepo, influxRepo: influxRepo}
}

func (e ema) Execute(c context.Context) error {
	zap.L().Info("Executing EMA strategy")
	defer zap.L().Info("EMA strategy executed successfully")

	candles, err := e.influxRepo.GetPrices(c, e.market, time.Minute*15)
	if err != nil {
		return err
	}

	log.Println(candles.Price.Candles)

	return nil
}
