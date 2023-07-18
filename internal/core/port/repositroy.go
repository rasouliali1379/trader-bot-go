package port

import (
	"context"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
)

type Exchange interface {
	GetCandles(ctx context.Context, market *domain.Market) (any, error)
	GetMarketSummary(ctx context.Context, market *domain.Market) (any, error)
	GetOrderBook(ctx context.Context, market *domain.Market) (any, error)
	BuyLimit(ctx context.Context, market *domain.Market, amount float64, limit float64) (string, error)
	SellLimit(ctx context.Context, market *domain.Market, amount float64, limit float64) (string, error)
	BuyMarket(ctx context.Context, market *domain.Market, amount float64) (string, error)
	SellMarket(ctx context.Context, market *domain.Market, amount float64) (string, error)
	CalculateTradingFees(ctx context.Context, market *domain.Market, amount float64, limit float64, orderType string) float64
	CalculateWithdrawFees(ctx context.Context, market *domain.Market, amount float64) float64
	GetBalance(ctx context.Context, symbol string) (any, error)
	GetDepositAddress(ctx context.Context, coinTicker string) (string, bool)
	FeedConnect(ctx context.Context, markets []*domain.Market) error
	Withdraw(ctx context.Context, destinationAddress string, coinTicker string, amount float64) error
}
