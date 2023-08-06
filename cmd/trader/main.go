package main

import (
	"context"
	"go.uber.org/zap"
	"hamgit.ir/novin-backend/trader-bot/config"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/infra/exchange"
	influx "hamgit.ir/novin-backend/trader-bot/internal/adapter/infra/influxdb"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/infra/log"
	marketjob "hamgit.ir/novin-backend/trader-bot/internal/adapter/job/market"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/repository/exchanges/okx"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/repository/influxdb"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
	marketsrv "hamgit.ir/novin-backend/trader-bot/internal/core/service/market"
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

	var marketList domain.MarketList
	marketList.FromConfig(config.C().Strategies)

	strategyMap := make(map[string]port.StrategyService)
	for _, s := range config.C().Strategies {
		if _, ok := strategyMap[s.Strategy]; !ok {
			strategyMap[s.Strategy] = strategies.New(domain.Strategy(s.Strategy), influxRepo)
		}
	}

	for _, m := range marketList.List {
		var job port.MarketJob
		switch m.Exchange.Name {
		case domain.OKX:
			if err := okxRepo.HasMarket(c, &m); err != nil {
				zap.L().Error("market wasn't found", zap.Any("market", m))
			}

			var observers domain.Observer

			for _, s := range m.Strategies {
				if strategy, ok := strategyMap[string(s)]; ok {
					observers.Register(strategy.Execute)
				}
			}

			job = marketjob.New(marketsrv.New(okxRepo, influxRepo, observers), okxRepo)
		case domain.Binance:
			//Same as OKX
		case domain.Kucoin:
			//Same as OKX
		default:
			zap.L().Fatal("unknown exchange")
		}

		zap.L().Fatal("error while running job", zap.Error(job.Run(context.Background(), &m)))
	}

	runtime.Goexit()
}
