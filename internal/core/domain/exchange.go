package domain

import (
	"errors"
	"hamgit.ir/novin-backend/trader-bot/config"
	"strings"
)

const (
	OKX     = "okx"
	Binance = "binance"
	Kucoin  = "kucoin"
)

var (
	ErrExchangeNotConnected = errors.New("exchange is either disconnected either isn't initialized")
)

type ExchangeList struct {
	List []Exchange
}

type Exchange struct {
	Name    string
	Markets MarketList
}

func (e ExchangeList) FromConfig(strategies []config.Strategy) {
	for _, s := range strategies {
		for _, market := range s.Markets {
			givetake := strings.Split(market.Market, "-")
			if len(givetake) == 2 {
				index := e.Find(market.Exchange)
				if index == -1 {
					e.List = append(e.List, Exchange{
						Name: market.Exchange,
						Markets: MarketList{List: []Market{
							{
								Give:       givetake[0],
								Take:       givetake[1],
								Strategies: []Strategy{Strategy(s.Strategy)},
							},
						}},
					})
				} else {
					e.List[index].Find()
				}
			}
		}
	}
}

func (e ExchangeList) Find(exchange string) int {
	for i := range e.List {
		if e.List[i].Name == exchange {
			return i
		}
	}

	return -1
}
