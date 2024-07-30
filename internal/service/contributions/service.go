package contributions

import (
	"fmt"
	"net/http"

	"github.com/abgeo/pensions/internal/dto"
	"github.com/abgeo/pensions/internal/errors"
	"github.com/abgeo/pensions/internal/service"
	"github.com/go-resty/resty/v2"
)

type Contributions struct {
	client *resty.Client
}

func New(client *resty.Client) *Contributions {
	return &Contributions{
		client: client,
	}
}

func (svc *Contributions) Get(options ...service.Option) (*dto.Contributions, *resty.Response, error) {
	var response *dto.Contributions

	req := svc.client.R()
	service.ApplyOptions(req, options)

	resp, err := req.
		SetError(&errors.V2HTTPError{}).
		SetResult(&response).
		Get("/v2/contributions/participant/get")
	if err != nil {
		return nil, resp, fmt.Errorf("unable to get Contributions: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, resp, errors.NewV2HTTPError(resp)
	}

	return response, resp, nil
}
