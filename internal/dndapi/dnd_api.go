package dndapi

import (
	"net/http"
	"time"
)

const (
	baseUrl = "https://www.dnd5eapi.co/api"
)

type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
