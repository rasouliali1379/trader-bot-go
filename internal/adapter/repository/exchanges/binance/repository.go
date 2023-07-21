package binance

import (
	"context"
	"golang.org/x/net/websocket"
	"hamgit.ir/novin-backend/trader-bot/internal/core/port"
)

type repository struct {
	exchangeClient *websocket.Conn
}

func New(exchangeClient *websocket.Conn) port.ExchangeRepository {
	return &repository{exchangeClient: exchangeClient}
}

func (r repository) Subscribe(c context.Context, channel string, instrumentID string) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) Unsubscribe(c context.Context, channel string, instrumentID string) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) Read(c context.Context) (any, error) {
	//TODO implement me
	panic("implement me")
}
