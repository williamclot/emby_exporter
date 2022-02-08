package emby

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	URL        string
	apiKey     string
	HTTPClient *http.Client
	ctx        context.Context
}

func New(ctx context.Context, baseURL, apiKey string) *Client {
	return &Client{
		URL:    baseURL,
		apiKey: apiKey,
		ctx:    ctx,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (e *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Emby-Token", e.apiKey)

	res, err := e.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		message, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
		}

		return fmt.Errorf("error %d %s", res.StatusCode, message)
	}

	if err = json.NewDecoder(res.Body).Decode(v); err != nil {
		return err
	}

	return nil
}
