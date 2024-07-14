package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/AranGarcia/droop/characters/internal/ports/api"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func (j ErrorResponse) ToBytes() []byte {
	data, _ := json.Marshal(j)
	return data
}

func handleAPIError(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}
	response := ErrorResponse{Message: err.Error()}
	w.WriteHeader(errorToStatusCode(err))
	w.Write(response.ToBytes())
}

func errorToStatusCode(err error) int {
	var jsonSyntaxError *json.SyntaxError
	if errors.As(err, &jsonSyntaxError) {
		return http.StatusBadRequest
	}

	switch err {
	case api.ErrNotFound:
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
