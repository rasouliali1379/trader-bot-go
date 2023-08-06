package domain

import (
	"fmt"
	"hamgit.ir/novin-backend/trader-bot/config"
	"strings"
)

type MarketList struct {
	List []Market
}

type Market struct {
	Give       string
	Take       string
	Exchange   *Exchange
	Price      *Price
	Strategies []Strategy
}

func (m *Market) String() string {
	return fmt.Sprintf("%s-%s", m.Give, m.Take)
}

func (m *MarketList) FindIndex(exchange string, give string, take string) int {
	for i := range m.List {
		if m.List[i].Exchange.Name == exchange && m.List[i].Give == give && m.List[i].Take == take {
			return i
		}
	}
	return -1
}

func (m *MarketList) FromConfig(strategies []config.Strategy) {
	for _, s := range strategies {
		for _, market := range s.Markets {
			givetake := strings.Split(market.Market, "-")
			if len(givetake) == 2 {
				index := m.FindIndex(market.Exchange, givetake[0], givetake[1])
				if index == -1 {
					m.List = append(m.List, Market{
						Give: givetake[0], Take: givetake[1],
						Exchange:   &Exchange{Name: market.Exchange},
						Strategies: []Strategy{Strategy(s.Strategy)},
					})
				} else {
					m.List[index].Strategies = append(
						m.List[index].Strategies, Strategy(s.Strategy))
				}
			}
		}
	}
}
