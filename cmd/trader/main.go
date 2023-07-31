package main

import (
	"context"
	"go.uber.org/zap"
	influx "hamgit.ir/novin-backend/trader-bot/internal/adapter/infra/influxdb"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/repository/exchanges/okx"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/repository/influxdb"
	"hamgit.ir/novin-backend/trader-bot/internal/core/service/strategies"
	"runtime"

	"hamgit.ir/novin-backend/trader-bot/config"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/infra/exchange"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/infra/log"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/job/market"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"hamgit.ir/novin-backend/trader-bot/internal/core/service/markets"
)

func main() {
	//infra
	config.Init("/dev/config/trader/")
	log.Init()

	influxWrite, influxRead := influx.Init()
	connectionManager := exchange.Init()

	testMarket := &domain.Market{
		Give:     "BTC",
		Take:     "USDT",
		Exchange: &domain.Exchange{Name: "okx"},
	}

	//Repo
	influxRepo := influxdb.New(influxWrite, influxRead)
	okxRepo := okx.New(testMarket.Exchange, connectionManager)

	//service
	emaStrategy := strategies.NewEmaStrategy(testMarket, okxRepo, influxRepo)
	var okxMarketObservers domain.Observer
	okxMarketObservers.Register(emaStrategy.Execute)
	okxMarketService := markets.NewOkxMarketService(okxRepo, influxRepo, okxMarketObservers)

	//jobs
	marketJob := market.New(okxMarketService, okxRepo)

	if err := marketJob.Run(context.Background(), testMarket); err != nil {
		zap.L().Fatal("error while running job", zap.Error(err))
	}

	runtime.Goexit()
}
