package domain

import (
	"github.com/markcheno/go-talib"
	"log"
	"strings"

	"hamgit.ir/novin-backend/trader-bot/internal/adapter/repository/exchanges/okx/dto"
	"strconv"
	"time"
)

type Price struct {
	Exchange *ExchangeItem
	Candles  []Candle
	Prices   []PriceItem
}

type Candle struct {
	Open  float64
	Close float64
	High  float64
	Low   float64
	Time  time.Time
}

type PriceItem struct {
	Price float64
	Time  time.Time
}

func (p *Price) FromIndexTickersDto(prices []dto.IndexTickers) *Price {
	priceList := make([]Candle, len(prices))

	for i := range prices {
		var price Candle
		priceList[i] = price.FromIndexTickersDto(prices[i])
	}

	return &Price{Candles: priceList}
}

func (p *Price) FromIndexCandlesDto(prices [][]string) *Price {
	priceList := make([]Candle, len(prices))

	for i := range prices {
		ts, _ := strconv.ParseInt(prices[i][0], 10, 64)
		open, _ := strconv.ParseFloat(prices[i][1], 10)
		high, _ := strconv.ParseFloat(prices[i][2], 10)
		low, _ := strconv.ParseFloat(prices[i][3], 10)
		cls, _ := strconv.ParseFloat(prices[i][4], 10)

		priceList[i] = Candle{
			Open:  open,
			Close: cls,
			High:  high,
			Low:   low,
			Time:  time.Unix(ts/1000, 0),
		}
	}

	return &Price{Candles: priceList}
}

func (p *Price) Ema(timePeriod int) []float64 {
	closePrices := p.close()
	return talib.Ema(closePrices, timePeriod)
}

func (p *Price) close() []float64 {
	closePrices := make([]float64, len(p.Candles))

	for i := range p.Candles {
		closePrices[i] = p.Candles[i].Close
	}

	return closePrices
}

func (i Candle) FromIndexTickersDto(price dto.IndexTickers) Candle {

	ts, _ := strconv.ParseInt(price.Ts, 10, 64)
	open, _ := strconv.ParseFloat(price.Open24H, 10)
	high, _ := strconv.ParseFloat(price.High24H, 10)
	low, _ := strconv.ParseFloat(price.Low24H, 10)

	return Candle{
		Time: time.Unix(ts/1000, 0),
		Open: open,
		High: high,
		Low:  low,
	}
}

func (p *Price) ParseFromInfluxDto(result string) {
	rows := strings.Split(result, "\n")

	rawData := rows[4:]

	dataMap := make(map[string]Candle)

	for i := range rawData {
		columns := strings.Split(rawData[i], ",")
		if len(columns) == 10 {
			value, ok := dataMap[columns[5]]
			if !ok {
				value = Candle{}
			}
			price, _ := strconv.ParseFloat(columns[6], 10)

			switch columns[7] {
			case "open":
				value.Open = price
			case "close":
				value.Close = price
			case "high":
				value.High = price
			case "low":
				value.Low = price
			}

			dataMap[columns[5]] = value

		}
	}

	var candles []Candle
	for key, value := range dataMap {
		parsedTime, err := time.Parse("2006-01-02T15:04:05Z", key)
		if err != nil {
			log.Println(err)
		}
		value.Time = parsedTime
		candles = append(candles, value)
	}

	p.Candles = candles
}
