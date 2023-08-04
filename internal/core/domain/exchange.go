package domain

const (
	OKX     = "okx"
	Binance = "binance"
	Kucoin  = "kucoin"
)

type Exchange struct {
	Name string
}
