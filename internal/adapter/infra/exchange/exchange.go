package exchange

import (
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"hamgit.ir/novin-backend/trader-bot/config"
)

const (
	Binance = iota
	OKX
)

type Exchange struct {
	Conns map[int]*websocket.Conn
}

func Init() *Exchange {
	conns := make(map[int]*websocket.Conn)

	conns[OKX] = newClient(config.C().OKX.Url)

	return &Exchange{Conns: conns}
}

func newClient(url string) *websocket.Conn {
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		zap.L().Fatal("error while creating exchange client", zap.Error(err))
	}
	return c
}
