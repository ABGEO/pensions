package errors

import (
	"fmt"
	"strings"

	"github.com/abgeo/pensions/internal/dto"
	"github.com/go-resty/resty/v2"
)

// @todo: combine to single error handler.

type V1HTTPError dto.HTTPResponse[string]

type V2HTTPError struct {
	ValidationStatus bool     `json:"validationStatus,omitempty"`
	Result           any      `json:"result,omitempty"`
	SuccessMessage   string   `json:"successMessage,omitempty"`
	ErrorMessages    []string `json:"errorMessages,omitempty"`
}

func NewV1HTTPError(response *resty.Response) *V1HTTPError {
	var errorMessage string

	if err, ok := response.Error().(*V1HTTPError); ok {
		return err
	}

	if response.Error() == nil {
		errorMessage = response.String()
	} else {
		errorMessage = fmt.Sprintf("%v", response.Error())
	}

	return &V1HTTPError{
		StatusCode: response.StatusCode(),
		Message:    errorMessage,
		Result:     errorMessage,
	}
}

func NewV2HTTPError(response *resty.Response) *V2HTTPError {
	var errorMessage string

	if err, ok := response.Error().(*V2HTTPError); ok {
		return err
	}

	if response.Error() == nil {
		errorMessage = response.String()
	} else {
		errorMessage = fmt.Sprintf("%v", response.Error())
	}

	return &V2HTTPError{
		ValidationStatus: false,
		ErrorMessages:    []string{errorMessage},
	}
}

func (err *V1HTTPError) Error() string {
	if len(err.Errors) > 0 {
		return strings.Join(err.Errors, "\r\n")
	}

	if err.Result != "" {
		return err.Result
	}

	return err.Message
}

func (err *V2HTTPError) Error() string {
	return strings.Join(err.ErrorMessages, "\r\n")
}
