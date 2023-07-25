package domain

import (
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/repository/exchanges/okx/dto"
	"strconv"
	"time"
)

type Price struct {
	Market *Market
	List   []PriceItem
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

func (p Price) FromIndexCandlesDto(prices [][]string) *Price {
	priceList := make([]PriceItem, len(prices))

	for i := range prices {
		ts, _ := strconv.ParseInt(prices[i][0], 10, 64)
		open, _ := strconv.ParseFloat(prices[i][1], 10)
		high, _ := strconv.ParseFloat(prices[i][2], 10)
		low, _ := strconv.ParseFloat(prices[i][3], 10)
		cls, _ := strconv.ParseFloat(prices[i][4], 10)

		priceList[i] = PriceItem{
			Open:  open,
			Close: cls,
			High:  high,
			Low:   low,
			Time:  time.Unix(ts/1000, 0),
		}
	}

	return &Price{List: priceList}
}

func (i PriceItem) FromIndexTickersDto(price dto.IndexTickers) PriceItem {

	ts, _ := strconv.ParseInt(price.Ts, 10, 64)
	open, _ := strconv.ParseFloat(price.Open24H, 10)
	high, _ := strconv.ParseFloat(price.High24H, 10)
	low, _ := strconv.ParseFloat(price.Low24H, 10)

	return PriceItem{
		Time: time.Unix(ts/1000, 0),
		Open: open,
		High: high,
		Low:  low,
	}
}
