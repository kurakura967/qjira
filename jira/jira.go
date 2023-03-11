package jira

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	token      string
}

func NewClient(client *http.Client, url, token string) *Client {
	return &Client{
		httpClient: client,
		baseURL:    url,
		token:      token,
	}
}

func (c *Client) NewRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	return req, nil
}
