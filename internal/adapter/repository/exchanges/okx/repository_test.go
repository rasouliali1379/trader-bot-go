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

	testMarket := &domain.Market{
		Give:     "BTC",
		Take:     "USDT",
		Exchange: &domain.Exchange{Name: "okx"},
	}

	okxRepo := New(testMarket.Exchange, connectionManager)

	if err := okxRepo.GetBalance(context.Background()); err != nil {
		t.Error(err)
	}

}
