package client

import (
	"net/http"
	"time"
)

const TIMEOUT = 10

type Client struct {
	Timeout time.Duration
}

type Connection interface {
	Request(url string) (resp *http.Response, err error)
}

func InitClient() Connection {
	return Client{
		Timeout: time.Duration(time.Second * TIMEOUT),
	}
}

func (c Client) Request(url string) (resp *http.Response, err error) {
	client := &http.Client{Timeout: c.Timeout}
	return client.Get(url)
}
