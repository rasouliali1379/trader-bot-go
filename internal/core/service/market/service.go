package market

import (
	"context"
	"fmt"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
)

type service struct {
	exchange   port.ExchangeRepository
	influxRepo port.InfluxRepository
	observers  domain.Observer
}

func New(
	okxRepo port.ExchangeRepository,
	influxRepo port.InfluxRepository,
	observers domain.Observer,
) port.MarketService {
	return &service{
		exchange:   okxRepo,
		influxRepo: influxRepo,
		observers:  observers,
	}
}

func (o service) SubscribeToMarket(c context.Context, m *domain.Market) error {
	return o.exchange.Subscribe(c, "index-candle1m", fmt.Sprintf("%s-%s", m.Give, m.Take))
}

func (o service) TrackMarket(c context.Context, m *domain.Market) {
	o.influxRepo.AddPrice(c, m)
	o.observers.NotifyAll(c, m)
}
