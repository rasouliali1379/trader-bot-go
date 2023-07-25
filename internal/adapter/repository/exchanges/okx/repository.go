package okx

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/repository/exchanges/okx/dto"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
)

type repository struct {
	exchangeClient *websocket.Conn
}

func New(exchangeClient *websocket.Conn) port.ExchangeRepository {
	return &repository{exchangeClient: exchangeClient}
}

func (r *repository) Subscribe(c context.Context, channel string, instrumentID string) error {
	request, err := createSubscribeRequest(channel, instrumentID)
	if err != nil {
		return err
	}

	if err = r.exchangeClient.WriteMessage(websocket.TextMessage, request); err != nil {
		return err
	}

	return nil
}

func (r *repository) Unsubscribe(c context.Context, channel string, instrumentID string) error {
	request, err := createUnsubscribeRequest(channel, instrumentID)
	if err != nil {
		return err
	}

	if err = r.exchangeClient.WriteMessage(websocket.BinaryMessage, request); err != nil {
		return err
	}

	return nil
}

func (r *repository) Read(c context.Context) (any, error) {
	_, msg, err := r.exchangeClient.ReadMessage()
	if err != nil {
		return nil, err
	}

	var dynamic dto.DynamicResponse
	if err := json.Unmarshal(msg, &dynamic); err != nil {
		return nil, err
	}

	switch dynamic.Arg.Channel {
	case "index-tickers":
		var indexTickers []dto.IndexTickers
		if err := json.Unmarshal(dynamic.Data, &indexTickers); err != nil {
			return nil, err
		}

		var m domain.Price
		return m.FromIndexTickersDto(indexTickers), nil
	case "index-candle1m":
		var indexCandles [][]string
		if err := json.Unmarshal(dynamic.Data, &indexCandles); err != nil {
			return nil, err
		}

		var m domain.Price
		return m.FromIndexCandlesDto(indexCandles), nil
	}

	return nil, domain.ErrUnknownType
}
