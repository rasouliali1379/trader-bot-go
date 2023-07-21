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
}

func NewOkxMarketService(okxRepo port.ExchangeRepository, influxRepo port.InfluxRepository) port.MarketService {
	return &okxMarketService{okxRepo: okxRepo, influxRepo: influxRepo}
}

func (o okxMarketService) SubscribeToMarket(c context.Context, m *domain.Market) error {
	return o.okxRepo.Subscribe(c, "index-tickers", fmt.Sprintf("%s-%s", m.Give, m.Take))
}

func (o okxMarketService) TrackMarket(c context.Context, m *domain.Price) error {

	return nil
}
