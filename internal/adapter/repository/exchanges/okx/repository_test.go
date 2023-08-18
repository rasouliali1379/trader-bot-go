package okx

import (
	"context"
	"hamgit.ir/novin-backend/trader-bot/config"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/infra/exchange"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/infra/log"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"testing"
)

func Test_repository_GetBalance(t *testing.T) {

	config.Init("../../../../../dev/config/trader/")
	log.Init()

	connectionManager := exchange.Init()

	okxRepo := New(connectionManager)

	if err := okxRepo.GetBalance(context.Background()); err != nil {
		t.Error(err)
	}

}

func Test_repository_PlaceOrder(t *testing.T) {

	config.Init("../../../../../dev/config/trader/")
	log.Init()

	connectionManager := exchange.Init()

	okxRepo := New(connectionManager)

	if err := okxRepo.PlaceOrder(context.Background(), &domain.Order{
		OrderPrice:   29000,
		Quantity:     1,
		InstrumentID: "BTC-USDT",
		Side:         "sell",
	}); err != nil {
		t.Error(err)
	}

}
