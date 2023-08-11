package okx

import (
	"context"
	"hamgit.ir/novin-backend/trader-bot/config"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/infra/exchange"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/infra/log"
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
