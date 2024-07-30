package service

import "github.com/go-resty/resty/v2"

type Option func(request *resty.Request)

func WithAuthToken(token string) Option {
	return func(request *resty.Request) {
		request.SetAuthToken(token)
	}
}

func WithDebug() Option {
	return func(request *resty.Request) {
		request.SetDebug(true)
	}
}

func ApplyOptions(request *resty.Request, options []Option) {
	for _, opt := range options {
		opt(request)
	}
}
