package dto

import "encoding/json"

type DynamicResponse struct {
	Arg struct {
		Channel string `json:"channel"`
		InstID  string `json:"instId"`
	} `json:"arg"`
	Data json.RawMessage `json:"data"`
}

type IndexTickers struct {
	InstID  string `json:"instId"`
	IdxPx   string `json:"idxPx"`
	High24H string `json:"high24h"`
	Low24H  string `json:"low24h"`
	Open24H string `json:"open24h"`
	SodUtc0 string `json:"sodUtc0"`
	SodUtc8 string `json:"sodUtc8"`
	Ts      string `json:"ts"`
}
