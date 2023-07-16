package exchanges

import "hamgit.ir/novin-backend/trader-bot/internal/core/port"

type binance struct {
}

func NewBinance() port.Exchange {
	return &binance{}
}
