package data_service

import "net/http"

type Client struct {
	httpClient *http.Client
}

func (c *Client) InitClient() *http.Client {
	c.httpClient = http.DefaultClient

	return c.httpClient
}
