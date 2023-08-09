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

	var exchangeList domain.ExchangeList
	exchangeList.FromConfig(config.C().Strategies)

	strategyMap := make(map[string]port.StrategyService)
	jobMap := make(map[string]port.MarketJob)
	for _, m := range marketList.List {
		switch m.Exchange.Name {
		case domain.OKX:
			if err := okxRepo.HasMarket(c, &m); err != nil {
				zap.L().Error("market wasn't found", zap.Any("market", m))
				continue
			}

			var observers domain.Observer

			for _, s := range m.Strategies {
				var strategy port.StrategyService
				var ok bool

				if strategy, ok = strategyMap[string(s)]; !ok {
					strategy = strategies.New(s, influxRepo, okxRepo)
					strategyMap[string(s)+m.Exchange.Name] = strategy
				}

				observers.Register(strategy.Execute)
			}

			var job port.MarketJob
			var ok bool

			if job, ok = jobMap[domain.OKX]; !ok {
				job = marketjob.New(marketsrv.New(&m, okxRepo, influxRepo, observers), okxRepo)
				jobMap[domain.OKX] = job
			}
		case domain.Binance:
			//Same as OKX
		case domain.Kucoin:
			//Same as OKX
		default:
			zap.L().Fatal("unknown exchange")
		}
	}

	for _, job := range jobMap {
		if err := job.Run(context.Background()); err != nil {
			zap.L().Fatal("error while running job", zap.Error(err))
		}
	}

	runtime.Goexit()
}
