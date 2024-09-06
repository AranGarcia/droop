package api

import (
	"net/url"
	"strconv"

	"github.com/AranGarcia/droop/characters/internal/core/entities"
)

type ListCharactersRequest struct {
	// Size of the page
	Size int `json:"size"`
	// Offset of the result set
	Offset int `json:"offset"`
}

// ListCharactersRequestFromURLQuery extracts the request's fields from the URL query parameters.
func ListCharactersRequestFromURLQuery(q url.Values) (*ListCharactersRequest, error) {
	r := &ListCharactersRequest{}

	var err error
	r.Size, err = strconv.Atoi(q.Get("size"))
	if err != nil || r.Size < 0 {
		return nil, InvalidRequestError{Fields: map[string]string{"size": "invalid value"}}
	}

	r.Offset, err = strconv.Atoi(q.Get("offset"))
	if err != nil || r.Offset < 0 {
		return nil, InvalidRequestError{Fields: map[string]string{"offset": "invalid value"}}
	}

	return r, nil
}

type ListCharactersResponse struct {
	Characters []entities.Character `json:"characters"`
}
