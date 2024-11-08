package turns

import (
	"errors"
	"fmt"

	"github.com/AranGarcia/droop/initiatives/internal/ports/out/repositories"
)

var (
	ErrNotFound = errors.New("not found")
)

func handleRepositoryErrors(err error) error {
	switch err {
	case repositories.ErrNotFound:
		return ErrNotFound
	}

	return fmt.Errorf("internal error; %v", err)
}
