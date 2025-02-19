package api

import (
	"fmt"

	"github.com/AranGarcia/droop/characters/internal/core/entities"
	commonerrors "github.com/AranGarcia/droop/shared/common-errors"
)

// ListSortyKey is the key in which the ListCharacter results will be sorted with.
type ListSortKey string

const (
	CreatedAt ListSortKey = "created_at"
	UpdatedAt ListSortKey = "updated_at"
)

var (
	listSorts map[ListSortKey]struct{} = map[ListSortKey]struct{}{
		CreatedAt: {}, UpdatedAt: {},
	}
)

// IsValid returns true if the sort key is a one of the valid options.
// TODO: Use validations with core/service.
func (s ListSortKey) IsValid() bool {
	_, ok := listSorts[s]
	return ok
}

type ListCharactersRequest struct {
	// Sort is the key to sort the results.
	Sort ListSortKey
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
