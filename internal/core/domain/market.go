package domain

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidInstrumentID = errors.New("invalid instrument id")
)

type MarketList []Market

type Market struct {
	Give       string
	Take       string
	Price      *Price
	Strategies []Strategy
	Strategy   Strategies
}

func (m *Market) String() string {
	return fmt.Sprintf("%s-%s", m.Give, m.Take)
}
