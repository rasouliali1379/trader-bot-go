package main

import (
	"context"
	"go.uber.org/zap"
	influx "hamgit.ir/novin-backend/trader-bot/internal/adapter/infra/influxdb"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/repository/exchanges/okx"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/repository/influxdb"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
	market_srv "hamgit.ir/novin-backend/trader-bot/internal/core/service/market"
	"hamgit.ir/novin-backend/trader-bot/internal/core/service/strategies"
	"hamgit.ir/novin-backend/trader-bot/internal/core/service/strategies/ema"

	"runtime"
	"strings"

	"hamgit.ir/novin-backend/trader-bot/config"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/infra/exchange"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/infra/log"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/job/market"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
)

func main() {
	//infra
	config.Init("/dev/config/trader/")
	log.Init()

	influxWrite, influxRead := influx.Init()
	connectionManager := exchange.Init()

	//Repo
	influxRepo := influxdb.New(influxWrite, influxRead)
	okxRepo := okx.New(connectionManager)

	markets := make(map[string]port.StrategyService)
	for _, s := range config.C().Strategies {
		for _, m := range s.Markets {
			switch m.Exchange {
			case domain.OKX:
				if value, ok := markets[m.Market]; !ok {
					var mrkt *domain.Market
					givetake := strings.Split(m.Market, "-")
					if len(givetake) == 2 {
						mrkt = &domain.Market{
							Give: givetake[0], Take: givetake[1],
							Exchange: &domain.Exchange{Name: domain.OKX},
						}
					}

					if mrkt == nil {
						markets[m.Market] = nil
						continue
					}

					if err := okxRepo.HasMarket(context.Background(), mrkt); err != nil {
						markets[m.Market] = nil
						continue
					}

					strategy := strategies.New(domain.Strategy(s.Strategy), okxRepo, influxRepo)

					markets[m.Market] = value
				}

			case domain.Binance:
				//Same as OKX
			case domain.Kucoin:
				//Same as OKX
			default:
				zap.L().Fatal("unknown exchange")
			}
		}
	}

	//service
	emaStrategy := ema.New(okxRepo, influxRepo)
	var okxMarketObservers domain.Observer
	okxMarketObservers.Register(emaStrategy.Execute)
	okxMarketService := market_srv.New(okxRepo, influxRepo, okxMarketObservers)

	//jobs
	marketJob := market.New(okxMarketService, okxRepo)

	if err := marketJob.Run(context.Background(), testMarket); err != nil {
		zap.L().Fatal("error while running job", zap.Error(err))
	}

	runtime.Goexit()
}
