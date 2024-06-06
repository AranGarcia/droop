package internal

import "encoding/json"

type JSONErrorResponse struct {
	Message string `json:"message,omitempty"`
}

func (j JSONErrorResponse) ToBytes() []byte {
	data, _ := json.Marshal(j)
	return data
}
