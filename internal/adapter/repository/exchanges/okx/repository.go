package okx

import (
	"context"
	"encoding/json"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/infra/exchange"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/repository/exchanges/okx/dto"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
)

type repository struct {
	exchange *domain.Exchange
	conn     *exchange.ConnectionManager
}

func New(exchange *domain.Exchange, conn *exchange.ConnectionManager) port.ExchangeRepository {
	return &repository{exchange: exchange, conn: conn}
}

func (r *repository) Subscribe(c context.Context, channel string, instrumentID string) error {
	request, err := createSubscribeRequest(channel, instrumentID)
	if err != nil {
		return err
	}

	if err = r.conn.OKX().Write(request); err != nil {
		return err
	}

	return nil
}

func (r *repository) Unsubscribe(c context.Context, channel string, instrumentID string) error {
	request, err := createUnsubscribeRequest(channel, instrumentID)
	if err != nil {
		return err
	}

	if err = r.conn.OKX().Write(request); err != nil {
		return err
	}

	return nil
}

func (r *repository) Read(_ context.Context) (any, error) {
	msg, err := r.conn.OKX().Read()
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
