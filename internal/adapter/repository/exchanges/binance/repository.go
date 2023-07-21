package binance

import (
	"context"
	"go.uber.org/fx"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/infra/exchange"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
)

type Repository struct {
	exchange *exchange.Exchange
}

type Params struct {
	fx.In
	Exchange *exchange.Exchange
}

var _ port.ExchangeRepository = &Repository{}

func New(params Params) Repository {
	return Repository{exchange: params.Exchange}
}

func (r Repository) GetCandles(ctx context.Context, market *domain.Market) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) GetMarketSummary(ctx context.Context, market *domain.Market) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) GetOrderBook(ctx context.Context, market *domain.Market) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) BuyLimit(ctx context.Context, market *domain.Market, amount float64, limit float64) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) SellLimit(ctx context.Context, market *domain.Market, amount float64, limit float64) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) BuyMarket(ctx context.Context, market *domain.Market, amount float64) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) SellMarket(ctx context.Context, market *domain.Market, amount float64) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) CalculateTradingFees(ctx context.Context, market *domain.Market, amount float64, limit float64, orderType string) float64 {
	//TODO implement me
	panic("implement me")
}

func (r Repository) CalculateWithdrawFees(ctx context.Context, market *domain.Market, amount float64) float64 {
	//TODO implement me
	panic("implement me")
}

func (r Repository) GetBalance(ctx context.Context, symbol string) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) GetDepositAddress(ctx context.Context, coinTicker string) (string, bool) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) FeedConnect(ctx context.Context, markets []*domain.Market) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Withdraw(ctx context.Context, destinationAddress string, coinTicker string, amount float64) error {
	//TODO implement me
	panic("implement me")
}
