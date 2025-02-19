package api

import (
	"fmt"

	"github.com/AranGarcia/droop/characters/internal/core/entities"
	"github.com/AranGarcia/droop/shared/common-errors"
)

type ListSortType string

const (
	CreatedAt ListSortType = "created_at"
	UpdatedAt ListSortType = "updated_at"
)

var (
	listSorts map[ListSortType]struct{} = map[ListSortType]struct{}{
		CreatedAt: {}, UpdatedAt: {},
	}
)

func (s ListSortType) IsValid() bool {
	_, ok := listSorts[s]
	return ok
}

type ListCharactersRequest struct {
	// Sort is the key to sort the results.
	Sort ListSortType
	// Size of the page
	Size int `json:"size"`
	// Offset of the result set
	Offset int `json:"offset"`
}

// Validate returns an error if the request has any invalid value.
func (r ListCharactersRequest) Validate() error {
	if !r.Sort.IsValid() {
		return commonerrors.InvalidInputError{
			Field:  "sort",
			Reason: fmt.Sprintf("invalid sort key '%s'", r.Sort),
		}
	}
	if r.Size < 0 {
		return commonerrors.InvalidInputError{
			Field:  "size",
			Reason: "must be greater than 0",
		}
	}
	if r.Offset < 0 {
		return commonerrors.InvalidInputError{
			Field:  "offset",
			Reason: "must be greater than 0",
		}
	}
	return nil
}

type ListCharactersResponse struct {
	Characters []entities.Character `json:"characters"`
}
