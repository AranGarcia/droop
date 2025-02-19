package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/AranGarcia/droop/characters/internal/core/api"
)

// InvalidInputError is thrown whenever an input to the adapter cannot be used.
type InvalidInputError struct {
	// Field that caused the error.
	Field string
	// Reason of being invalid.
	Reason string
}

// Error implements the error interface.
func (i InvalidInputError) Error() string {
	return fmt.Sprintf("invalid value for %q: %s", i.Field, i.Reason)
}

// Is allows comparison through the errors.Is function.
func (i InvalidInputError) Is(target error) bool {
	_, ok := target.(InvalidInputError)
	return ok
}

type ErrorResponse struct {
	// Message is the error message.
	Message string `json:"error,omitempty"`
	// Fields from the request that may have caused the error.
	Fields map[string]string `json:"fields,omitempty"`
}

// NewErrorResponse returns the structure containing the data for the JSON response for an error.
func NewErrorResponse(err error) ErrorResponse {
	var apiInvalidRequest api.InvalidRequestError
	if !errors.As(err, &apiInvalidRequest) {
		return ErrorResponse{Message: err.Error()}
	}
	fields := make(map[string]string, len(apiInvalidRequest.Fields))
	for k, v := range apiInvalidRequest.Fields {
		fields[k] = v
	}
	return ErrorResponse{
		Message: "some fields contain invalid values",
		Fields:  fields,
	}
}

func (j ErrorResponse) ToBytes() []byte {
	data, _ := json.Marshal(j)
	return data
}

func handleAPIError(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}
	response := NewErrorResponse(err)
	w.WriteHeader(errorToStatusCode(err))
	w.Write(response.ToBytes())
}

func errorToStatusCode(err error) int {
	var jsonSyntaxError *json.SyntaxError
	if errors.As(err, &jsonSyntaxError) {
		return http.StatusBadRequest
	}

	var validationError api.InvalidRequestError
	if errors.As(err, &validationError) {
		return http.StatusBadRequest
	}

	switch err {
	case api.ErrNotFound:
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
