package okx

import (
	"context"
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
	"hamgit.ir/novin-backend/trader-bot/config"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/infra/exchange"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/repository/exchanges/okx/dto"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
	"log"
	"net/url"
	"strings"
)

const (
	balancePath = "/api/v5/account/balance"
)

type repository struct {
	conn *exchange.ConnectionManager
}

func New(conn *exchange.ConnectionManager) port.ExchangeRepository {
	return &repository{conn: conn}
}

func (r *repository) Name() domain.Exchange {
	return domain.OKX
}

func (r *repository) GetBalance(c context.Context) error {

	path, err := url.JoinPath(config.C().OKX.HttpUrl, balancePath)
	if err != nil {
		return err
	}

	headers := createOKXAuthHeader(
		"GET", balancePath, config.C().OKX.ApiKey, config.C().OKX.SecretKey, config.C().OKX.Passphrase)

	res, err := r.conn.Http().Get(c, path, headers)
	if err != nil {
		return err
	}

	var balance dto.Balance
	if err = jsoniter.Unmarshal(res, &balance); err != nil {
		return err
	}
	log.Println(balance)
	return nil
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

		var p domain.Price
		p.FromIndexCandlesDto(indexCandles)

		instID := strings.Split(dynamic.Arg.InstID, "-")

		if len(instID) < 2 {
			return nil, domain.ErrInvalidInstrumentID
		}

		return &domain.Market{
			Give:  instID[0],
			Take:  instID[1],
			Price: p.FromIndexCandlesDto(indexCandles)}, nil
	}

	return nil, domain.ErrUnknownType
}

func (r *repository) HasMarket(c context.Context, m *domain.Market) error {
	return nil
}
