package okx

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hamgit.ir/novin-backend/trader-bot/internal/adapter/repository/exchanges/okx/dto"
	"hamgit.ir/novin-backend/trader-bot/internal/core/domain"
	"strconv"
	"time"
)

func createSubscribeRequest(channel string, instrumentID string) ([]byte, error) {
	request := dto.SubscribeRequest{
		Op: "subscribe",
		Args: []struct {
			Channel string `json:"channel"`
			InstID  string `json:"instId"`
		}{
			{Channel: channel, InstID: instrumentID},
		},
	}

	marshal, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	return marshal, nil
}

func createUnsubscribeRequest(channel string, instrumentID string) ([]byte, error) {
	request := dto.SubscribeRequest{
		Op: "unsubscribe",
		Args: []struct {
			Channel string `json:"channel"`
			InstID  string `json:"instId"`
		}{
			{Channel: channel, InstID: instrumentID},
		},
	}

	marshal, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	return marshal, nil
}

func createOKXSignature(method string, path string, secret string, body []byte) string {
	h := hmac.New(sha256.New, []byte(secret))
	var s string
	if body == nil {
		s = fmt.Sprintf("%s%s%s", time.Now().UTC().Format("2006-01-02T15:04:05.000Z"), method, path)
	} else {
		s = fmt.Sprintf("%s%s%s%s", time.Now().UTC().Format("2006-01-02T15:04:05.000Z"), method, path, string(body))
	}
	h.Write([]byte(s))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func createOKXAuthHeader(method string, path string, apiKey string, secret string, passphrase string, body []byte) map[string]string {
	headers := make(map[string]string)
	headers["OK-ACCESS-KEY"] = apiKey
	headers["OK-ACCESS-SIGN"] = createOKXSignature(method, path, secret, body)
	headers["OK-ACCESS-TIMESTAMP"] = time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	headers["OK-ACCESS-PASSPHRASE"] = passphrase
	headers["x-simulated-trading"] = "1"
	return headers
}

func createPlaceOrderRequest(m *domain.Order) ([]byte, error) {
	request := dto.PlaceOrderRequest{
		TdMode:  "cash",
		OrdType: "limit",
		InstID:  m.InstrumentID,
		Side:    string(m.Side),
		Px:      strconv.FormatFloat(m.OrderPrice, 'g', 5, 64),
		Sz:      strconv.FormatFloat(m.Quantity, 'g', 5, 64),
	}
	return json.Marshal(request)
}
