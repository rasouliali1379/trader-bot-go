package exchange

import (
	"context"
	"io"
	"net/http"
)

type HttpWrapper interface {
	Get(c context.Context, url string) ([]byte, error)
	Post(c context.Context, url string, body io.Reader) ([]byte, error)
}

type httpWrapper struct {
	client *http.Client
}

func newHttpWrapper() HttpWrapper {
	return &httpWrapper{client: &http.Client{}}
}

func (h *httpWrapper) Get(c context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (h *httpWrapper) Post(c context.Context, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequestWithContext(c, "POST", url, body)
	if err != nil {
		return nil, err
	}

	res, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}
