package okx

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type subscribeRequest struct {
	Op   string `json:"op"`
	Args []struct {
		Channel string `json:"channel"`
		InstID  string `json:"instId"`
	} `json:"args"`
}

func createSubscribeRequest(channel string, instrumentID string) ([]byte, error) {
	request := subscribeRequest{
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
	request := subscribeRequest{
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

func createOKXSignature(method string, path string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	s := fmt.Sprintf("%s%s%s", time.Now().UTC().Format("2006-01-02T15:04:05.000Z"), method, path)
	h.Write([]byte(s))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func createOKXAuthHeader(method string, path string, apiKey string, secret string, passphrase string) *http.Header {
	sign := createOKXSignature(method, path, secret)
	headers := http.Header{}
	headers.Set("OK-ACCESS-KEY", apiKey)
	headers.Set("OK-ACCESS-SIGN", sign)
	headers.Set("OK-ACCESS-TIMESTAMP", time.Now().UTC().Format("2006-01-02T15:04:05.000Z"))
	headers.Set("OK-ACCESS-PASSPHRASE", passphrase)
	return &headers
}
