package client

import (
	"encoding/json"

	"github.com/abgeo/pensions/internal/config"
	"github.com/go-resty/resty/v2"
)

func New(conf config.PensionsConfig) *resty.Client {
	httpClient := resty.New()

	httpClient.
		SetBaseURL(conf.URL).
		SetTimeout(conf.ClientTimeout).
		SetRetryCount(conf.ClientRetryCount).
		SetDebug(conf.ClientDebug).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept-Type", "application/json")

	httpClient.JSONMarshal = json.Marshal
	httpClient.JSONUnmarshal = json.Unmarshal

	return httpClient
}
