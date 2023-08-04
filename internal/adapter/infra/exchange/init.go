package exchange

import (
	"go.uber.org/zap"
	"hamgit.ir/novin-backend/trader-bot/config"
)

const (
	OKX = iota
	Binance
	Kucoin
)

type ConnectionManager struct {
	w map[int]WebSocketWrapper
	h HttpWrapper
}

func Init() *ConnectionManager {
	conns := make(map[int]WebSocketWrapper)

	conns[OKX] = newWebSocketWrapper(config.C().OKX.WebSocketUrl)

	return &ConnectionManager{
		w: conns,
		h: newHttpWrapper(),
	}
}

func (c *ConnectionManager) Http() HttpWrapper {
	return c.h
}

func (c *ConnectionManager) OKX() WebSocketWrapper {
	return c.getConnection(OKX)
}

func (c *ConnectionManager) Binance() WebSocketWrapper {
	return c.getConnection(Binance)
}

func (c *ConnectionManager) Kucoin() WebSocketWrapper {
	return c.getConnection(Kucoin)
}

func (c *ConnectionManager) getConnection(conn int) WebSocketWrapper {
	if conn, ok := c.w[conn]; ok {
		return conn
	}

	zap.L().Fatal("connection wasn't found")
	return nil
}
