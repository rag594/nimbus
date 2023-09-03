package httpClient

import (
	"context"
	"net/http"
	"time"
)

// Client defines Http Client
type Client struct {
	httpClient *http.Client
}

// Init initializes HTTP client
func Init() *Client {
	to, _ := time.ParseDuration("60s")
	t := http.DefaultTransport.(*http.Transport).Clone()
	client := &http.Client{
		Transport: t,
		Timeout:   to,
	}

	return &Client{httpClient: client}
}

// Get make an HTTP GET request
func (c *Client) Get(ctx context.Context, uri string, headers map[string][]string) (*http.Response, error) {
	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	if headers == nil {
		headers = map[string][]string{}
	}
	headers["Content-Type"] = []string{"application/json"}

	return c.doRequest(ctx, req, headers)
}

func (c *Client) doRequest(ctx context.Context, req *http.Request, headers map[string][]string) (*http.Response, error) {
	if ctx == nil {
		res, err := c.httpClient.Do(req)
		return res, err
	}

	res, err := c.httpClient.Do(req)
	return res, err
}
