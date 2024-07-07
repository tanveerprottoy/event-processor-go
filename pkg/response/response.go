package response

import (
	"encoding/json"
	"github.com/tanveerprottoy/event-processor-go/pkg/constant"
	"net/http"
)

type Response[T any] struct {
	Data T `json:"data"`
}

type Error struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	StatusCode int `json:"statusCode"`
	Errors     any `json:"errors"`
}

func BuildData[T any](payload T) Response[T] {
	return Response[T]{Data: payload}
}

func RespondError(w http.ResponseWriter, code int, payload any) (int, error) {
	w.WriteHeader(code)
	res, errs := json.Marshal(payload)
	if errs != nil {
		// log failed to marshal
		return w.Write([]byte(constant.InternalServerError))
	}
	return w.Write(res)
}

func Respond(w http.ResponseWriter, code int, payload any) (int, error) {
	res, err := json.Marshal(payload)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, ErrorResponse{StatusCode: http.StatusInternalServerError, Errors: []any{"an error occured"}})
		return -1, err
	}
	w.WriteHeader(code)
	return w.Write(res)
}
