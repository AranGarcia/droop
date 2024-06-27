package http

import "encoding/json"

// Deprecated: use ErrorResponse instead.
type JSONErrorResponse struct {
	Error string `json:"error,omitempty"`
}

func (j JSONErrorResponse) ToBytes() []byte {
	data, _ := json.Marshal(j)
	return data
}
