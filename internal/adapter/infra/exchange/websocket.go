package exchange

import (
	"errors"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"sync"
	"time"
)

const (
	pingPeriod = time.Second * 5
)

var (
	ErrUnexpectedConnectionClose = errors.New("connection was closed unexpectedly")
)

type WebSocketWrapper interface {
	Read() ([]byte, error)
	Write(msg []byte) error
	IsConnected() bool
}

type webSocketWrapper struct {
	m           sync.Mutex
	conn        *websocket.Conn
	isConnected bool
}

func newWebSocketWrapper(url string) WebSocketWrapper {
	wrapper := &webSocketWrapper{conn: newConnection(url)}
	wrapper.keepAlive()
	return wrapper
}

func (w *webSocketWrapper) Read() ([]byte, error) {
	_, msg, err := w.conn.ReadMessage()
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			return nil, ErrUnexpectedConnectionClose
		}
	}
	return msg, nil
}

func (w *webSocketWrapper) Write(msg []byte) error {
	w.m.Lock()
	defer w.m.Unlock()
	return w.conn.WriteMessage(websocket.TextMessage, msg)
}

func (w *webSocketWrapper) IsConnected() bool {
	return w.isConnected
}

func newConnection(url string) *websocket.Conn {
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		zap.L().Fatal("error while creating exchange client", zap.Error(err))
	}
	return c
}

func (w *webSocketWrapper) keepAlive() {
	go func() {
		ticker := time.NewTicker(pingPeriod)
		for ; true; <-ticker.C {
			w.m.Lock()
			if err := w.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				zap.L().Error("error while pinging", zap.Error(err))
			}
			w.m.Unlock()
		}
	}()
}
