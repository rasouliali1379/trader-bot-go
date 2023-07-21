package okx

import "encoding/json"

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
