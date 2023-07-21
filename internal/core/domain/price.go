package domain

import (
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/repository/exchanges/okx/dto"
	"strconv"
	"time"
)

type Price struct {
	List []PriceItem
}

type PriceItem struct {
	Open  float64
	Close float64
	High  float64
	Low   float64
	Time  time.Time
}

func (p Price) FromIndexTickersDto(prices []dto.IndexTickers) *Price {
	priceList := make([]PriceItem, len(prices))

	for i := range prices {
		var price PriceItem
		priceList[i] = price.FromIndexTickersDto(prices[i])
	}

	return &Price{List: priceList}
}

func (i PriceItem) FromIndexTickersDto(price dto.IndexTickers) PriceItem {

	ts, _ := strconv.ParseInt(price.Ts, 10, 64)
	open, _ := strconv.ParseFloat(price.Open24H, 10)
	high, _ := strconv.ParseFloat(price.High24H, 10)
	low, _ := strconv.ParseFloat(price.Low24H, 10)

	return PriceItem{
		Time: time.Unix(0, ts),
		Open: open,
		High: high,
		Low:  low,
	}
}
