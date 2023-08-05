package domain

import "errors"

const (
	OKX     = "okx"
	Binance = "binance"
	Kucoin  = "kucoin"
)

var (
	ErrExchangeNotConnected = errors.New("exchange is either disconnected either isn't initialized")
)

type Exchange struct {
	Name string
}
