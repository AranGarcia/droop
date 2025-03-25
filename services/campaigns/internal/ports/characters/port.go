package characters

import (
	"context"

	"github.com/AranGarcia/droop/campaigns/internal/core/entities"
)

type Port interface {
	List(context.Context, int, int) ([]entities.Character, error)
}
