package auth

import (
	"fmt"
	"net/http"

	"github.com/abgeo/pensions/internal/dto"
	"github.com/abgeo/pensions/internal/service"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

var errUnableToAuthenticateUser = errors.New("unable to authenticate user")

type Auth struct {
	client *resty.Client
}

func New(client *resty.Client) *Auth {
	return &Auth{
		client: client,
	}
}

func (svc *Auth) Authenticate(
	username string,
	password string,
	options ...service.Option,
) (*dto.AuthResponse, *resty.Response, error) {
	var response *dto.AuthResponse

	body := dto.AuthBody{
		Username:     username,
		PasswordOne:  password,
		LanguageCode: "ka-GE",
	}

	req := svc.client.R()
	service.ApplyOptions(req, options)

	resp, err := svc.client.R().
		SetResult(&response).
		SetError(&response).
		SetBody(body).
		Post("/v1/auth/participant-auth")
	if err != nil {
		return nil, resp, fmt.Errorf("unable to Authenticate User: %w", err)
	}

	if resp.StatusCode() != http.StatusOK || response.AccessToken == "" {
		return nil, resp, errUnableToAuthenticateUser
	}

	return response, resp, nil
}
