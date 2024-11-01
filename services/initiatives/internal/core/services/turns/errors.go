package turns

import (
	"fmt"
)

func handleRepositoryErrors(err error) error {
	// TODO: translate to core errors.
	return fmt.Errorf("internal error; %v", err)
}
