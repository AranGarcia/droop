package http

import (
	"encoding/json"
	"net/http"

	"github.com/AranGarcia/droop/characters/internal/ports/api"
)

type ErrorResponse struct {
	Message string `json:"Message"`
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
	switch err {
	case api.ErrNotFound:
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
