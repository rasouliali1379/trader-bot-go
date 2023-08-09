package domain

import (
	"fmt"
)

type MarketList struct {
	List []Market
}

type Market struct {
	Give       string
	Take       string
	Price      *Price
	Strategies []Strategy
}

func (m *Market) String() string {
	return fmt.Sprintf("%s-%s", m.Give, m.Take)
}
