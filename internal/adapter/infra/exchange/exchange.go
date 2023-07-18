package exchange

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"golang.org/x/net/websocket"
	"hamgit.ir/novin-backend/trader-bot/config"
)

const (
	Binance = iota
	OKX
)

type Exchange struct {
	conns map[int]*websocket.Conn
}

func Init(lc fx.Lifecycle) *Exchange {
	conns := make(map[int]*websocket.Conn)
	lc.Append(fx.Hook{
		OnStart: func(c context.Context) error {
			conns[OKX] = newClient(config.C().OKX.Url)
			return nil
		},
		OnStop: func(c context.Context) error {
			for _, conn := range conns {
				err := conn.Close()
				if err != nil {
					zap.L().Error("error while closing the connection", zap.Error(err))
				}
			}
			return nil
		},
	})

	return &Exchange{}
}

func newClient(url string) *websocket.Conn {
	dial, err := websocket.Dial(url, "", "")
	if err != nil {
		zap.L().Fatal("error while creating exchange client", zap.Error(err))
	}
	return dial
}
