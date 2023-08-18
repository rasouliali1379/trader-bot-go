package domain

type OrderSide string

const (
	OrderSideBuy  OrderSide = "buy"
	OrderSideSell           = "sell"
)

type Order struct {
	OrderPrice   float64
	Quantity     float64
	InstrumentID string
	Side         OrderSide
}
