package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"hamgit.ir/novin-backend/trader-bot/config"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/infra/exchange"
	influx "hamgit.ir/novin-backend/trader-bot/internal/adapter/infra/influxdb"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/infra/log"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/repository/exchanges/okx"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/repository/influxdb"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
	exchangesrv "hamgit.ir/novin-backend/trader-bot/internal/core/service/exchange"
	"hamgit.ir/novin-backend/trader-bot/internal/core/service/strategies"
	"runtime"
)

func main() {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()
	//infra
	config.Init("/dev/config/trader/")
	log.Init()

	influxWrite, influxRead := influx.Init()
	connectionManager := exchange.Init()

	influxRepo := influxdb.New(influxWrite, influxRead)
	okxRepo := okx.New(connectionManager)

	var exchangeList domain.ExchangeList
	exchangeList = exchangeList.FromConfig(config.C().Strategies)

	strategyMap := make(map[string]port.StrategyService)
	for i := range exchangeList {
		exchangeService := exchangesrv.New(okxRepo, influxRepo)
		switch exchangeList[i].Name {
		case domain.OKX:
			for _, market := range exchangeList[i].Markets {
				if err := okxRepo.HasMarket(c, &market); err != nil {
					zap.L().Error("exchange wasn't found", zap.Any("exchange", &market))
					continue
				}

				var observers domain.Strategies
				for _, s := range market.Strategies {
					var strategy port.StrategyService
					var ok bool

					if strategy, ok = strategyMap[string(s)]; !ok {
						strategy = strategies.New(s, influxRepo, okxRepo)
						strategyMap[fmt.Sprintf("%s-%s", s, exchangeList[i].Name)] = strategy
					}

					observers.Register(strategy.Execute)
				}

				market.Strategy = observers
				if err := exchangeService.TrackMarket(c, &market); err != nil {
					return
				}
			}

		case domain.Binance:
			//Same as OKX
		case domain.Kucoin:
			//Same as OKX
		default:
			zap.L().Fatal("unknown exchange")
		}
		if exchangeService != nil {
			exchangeService.WatchMarkets(context.Background())
		}
	}

	runtime.Goexit()
}
