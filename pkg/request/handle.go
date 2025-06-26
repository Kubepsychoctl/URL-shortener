package request

import (
	"net/http"

	"app/url-shorter/pkg/response"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		response.Json(*w, err.Error(), 402)
		return nil, err
	}
	err = IsValid(body)
	if err != nil {
		response.Json(*w, err.Error(), 402)
		return nil, err
	}
	return &body, nil
}
