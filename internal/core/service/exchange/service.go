package exchange

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
)

type service struct {
	exchangeRepo port.ExchangeRepository
	influxRepo   port.InfluxRepository
	markets      map[string]*domain.Market
}

func New(
	exchangeRepo port.ExchangeRepository,
	influxRepo port.InfluxRepository,
) port.ExchangeService {
	return &service{
		exchangeRepo: exchangeRepo,
		influxRepo:   influxRepo,
		markets:      make(map[string]*domain.Market),
	}
}

func (o service) TrackMarket(c context.Context, m *domain.Market) error {

	if err := o.exchangeRepo.Subscribe(c, "index-candle1m", fmt.Sprintf("%s-%s", m.Give, m.Take)); err != nil {
		return err
	}

	zap.L().Info("successfully subscribed on channel",
		zap.String("give", m.Give),
		zap.String("take", m.Take))

	o.markets[m.Give+m.Take] = m

	return nil
}

func (o service) WatchMarkets(c context.Context) {
	go func() {
		for {
			msg, err := o.exchangeRepo.Read(c)
			if err != nil {
				zap.L().Error(fmt.Sprintf("error while tracing"), zap.Error(err))
				continue
			}

			switch msg.(type) {
			case *domain.Market:
				m := msg.(*domain.Market)

				if _, ok := o.markets[m.Give+m.Take]; !ok {
					zap.L().Error("unregistered market data", zap.Any("market", m))
				}

				o.influxRepo.AddPrice(c, o.exchangeRepo.Name(), m)
				o.markets[m.Give+m.Take].Strategy.NotifyAll(c, o.exchangeRepo.Name(), m)
			}
		}
	}()
}
