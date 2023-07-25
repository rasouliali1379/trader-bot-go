package markets

import (
	"context"
	"fmt"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
)

type okxMarketService struct {
	okxRepo    port.ExchangeRepository
	influxRepo port.InfluxRepository
	observers  domain.Observer
}

func NewOkxMarketService(
	okxRepo port.ExchangeRepository,
	influxRepo port.InfluxRepository,
	observers domain.Observer,
) port.MarketService {
	return &okxMarketService{
		okxRepo:    okxRepo,
		influxRepo: influxRepo,
		observers:  observers,
	}
}

func (o okxMarketService) SubscribeToMarket(c context.Context, m *domain.Market) error {
	return o.okxRepo.Subscribe(c, "index-candle1m", fmt.Sprintf("%s-%s", m.Give, m.Take))
}

func (o okxMarketService) TrackMarket(c context.Context, m *domain.Price) {
	o.influxRepo.AddPoint(c, m)
	o.observers.NotifyAll(c)
}
