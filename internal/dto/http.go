package dto

type HTTPResponse[T any] struct {
	Errors      []string `json:"errors,omitempty"`
	StatusCode  int      `json:"statusCode"`
	MessageType int      `json:"messageType"`
	Message     string   `json:"message"`
	Result      T        `json:"result,omitempty"`
}
