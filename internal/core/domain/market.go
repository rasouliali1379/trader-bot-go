package domain

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type Market struct {
	Give     string
	Take     string
	Exchange *Exchange
	Price    *Price
}

func (m *Market) String() string {
	return fmt.Sprintf("%s-%s", m.Give, m.Take)
}

func (m *Market) ParseFromInfluxDto(result string) *Market {
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

	candles := make([]Candle, len(dataMap))
	for key, value := range dataMap {
		parsedTime, err := time.Parse(key, "2006-01-02T15:04:05Z")
		if err != nil {
			log.Println(err)
		}
		value.Time = parsedTime
		candles = append(candles, value)
	}

	m.Price.Candles = candles
	return m
}
