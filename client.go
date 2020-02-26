package redmine

import "net/http"

type Client struct {
	endpoint string
	apikey   string
	*http.Client
}

func NewClient(endpoint, apikey string) *Client {
	return &Client{endpoint, apikey, http.DefaultClient}
}