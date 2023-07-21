package okx

import (
	"context"
	"encoding/json"
	"golang.org/x/net/websocket"
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

func (r repository) Subscribe(c context.Context, channel string, instrumentID string) error {
	request, err := createSubscribeRequest(channel, instrumentID)
	if err != nil {
		return err
	}

	_, err = r.exchangeClient.Write(request)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) Unsubscribe(c context.Context, channel string, instrumentID string) error {
	request, err := createUnsubscribeRequest(channel, instrumentID)
	if err != nil {
		return err
	}

	_, err = r.exchangeClient.Write(request)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) Read(c context.Context) (any, error) {
	var msg []byte
	_, err := r.exchangeClient.Read(msg)
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
		if err := json.Unmarshal(msg, &indexTickers); err != nil {
			return nil, err
		}

		var m domain.Price
		return m.FromIndexTickersDto(indexTickers), nil
	}

	return nil, domain.ErrUnknownType
}
