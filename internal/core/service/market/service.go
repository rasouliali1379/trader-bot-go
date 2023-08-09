package market

import (
	"context"
	"fmt"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
)

type service struct {
	market     *domain.Market
	exchange   port.ExchangeRepository
	influxRepo port.InfluxRepository
	observers  domain.Observer
}

func New(
	market *domain.Market,
	exchange port.ExchangeRepository,
	influxRepo port.InfluxRepository,
	observers domain.Observer,
) port.MarketService {
	return &service{
		market:     market,
		exchange:   exchange,
		influxRepo: influxRepo,
		observers:  observers,
	}
}

func (o service) SubscribeToMarket(c context.Context, m *domain.Market) error {
	return o.exchange.Subscribe(c, "index-candle1m", fmt.Sprintf("%s-%s", m.Give, m.Take))
}

func (o service) TrackMarket(c context.Context, p *domain.Price) {
	o.market.Price = p
	o.influxRepo.AddPrice(c, o.market)
	o.observers.NotifyAll(c, o.market)
}
