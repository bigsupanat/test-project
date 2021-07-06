package service

import "net/http"

type Service struct {
	RequestFn func(url string) (resp *http.Response, err error)
}

func (s Service) Request(url string) (resp *http.Response, err error) {
	return s.RequestFn(url)
}
