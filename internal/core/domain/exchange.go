package domain

import (
	"errors"
	"go.uber.org/zap"
	"hamgit.ir/novin-backend/trader-bot/config"
	"strings"
)

type Exchange string

const (
	OKX     Exchange = "okx"
	Binance          = "binance"
	Kucoin           = "kucoin"
)

var (
	ErrExchangeNotConnected = errors.New("exchange is either disconnected either isn't initialized")
)

type ExchangeList []ExchangeItem

type ExchangeItem struct {
	Name    Exchange
	Markets MarketList
}

func (e ExchangeList) FromConfig(strategies []config.Strategy) ExchangeList {

	exchangeList := make(map[string]map[string]map[string]bool)

	for _, s := range strategies {
		for i := range s.Markets {
			exchange := s.Markets[i].Exchange
			market := s.Markets[i].Market

			if exchangeList[exchange] == nil {
				exchangeList[exchange] = make(map[string]map[string]bool)
			}
			if exchangeList[exchange][market] == nil {
				exchangeList[exchange][market] = make(map[string]bool)
			}
			exchangeList[exchange][market][s.Strategy] = true
		}
	}

	var exchanges ExchangeList

	for exchange, marketMap := range exchangeList {
		var markets []Market
		for market, strategyMap := range marketMap {
			var strategyList []Strategy
			for s := range strategyMap {
				strategyList = append(strategyList, Strategy(s))
			}

			givetake := strings.Split(market, "-")
			if len(givetake) != 2 {
				zap.L().Fatal("invalid exchange name")
			}
			markets = append(markets, Market{
				Give:       givetake[0],
				Take:       givetake[1],
				Strategies: strategyList,
			})
		}
		exchanges = append(exchanges, ExchangeItem{
			Name:    Exchange(exchange),
			Markets: markets,
		})
	}

	return exchanges
}

func (e ExchangeList) Find(exchange Exchange) int {
	for i := range e {
		if e[i].Name == exchange {
			return i
		}
	}

	return -1
}
