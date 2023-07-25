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

	ctx := context.Background()

	config.Init()
	log.Init()

	ex := exchange.Init()
	influxApi := influx.Init()

	influxRepo := influxdb.New(influxApi)
	okxRepo := okx.New(ex.Conns[exchange.OKX])

	emaStrategy := strategies.NewEmaStrategy(okxRepo)

	var okxMarketObservers domain.Observer
	okxMarketObservers.Register(emaStrategy.Execute)

	okxMarketService := markets.NewOkxMarketService(okxRepo, influxRepo, okxMarketObservers)

	marketJob := market.New(okxMarketService, okxRepo)

	for _, s := range config.C().Strategies {
		for _, _ = range s.Markets {
			if err := marketJob.Run(ctx, &domain.Market{Exchange: "okx", Give: "BTC", Take: "USDT"}); err != nil {
				zap.L().Fatal("error while running job", zap.Error(err))
			}
		}
	}

	runtime.Goexit()
	//http.HandleFunc("/health",
	//	func(writer http.ResponseWriter, request *http.Request) {
	//		_, _ = io.WriteString(writer, "nebula is up")
	//	})
	//
	//zap.L().Fatal("", zap.Error(http.ListenAndServe(":"+config.C().Server.Port, nil)))
}
