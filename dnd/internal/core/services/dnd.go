package services

import (
	"context"

	"github.com/AranGarcia/droop/dnd/internal/core/entities"
	"github.com/AranGarcia/droop/dnd/internal/ports/core"
	"github.com/AranGarcia/droop/dnd/internal/ports/external/characters"
)

type DND struct {
	d20 entities.Die

	// External ports
	characters characters.Port
}

func (d DND) RollInitiative(ctx context.Context, request core.RollInitiativeRequest) (*core.RollInitiativeResponse, error) {
	character, err := d.characters.Retrieve(ctx, request.ID)
	if err != nil {
		return nil, core.NewExternalServiceError("characters", err)
	}

	roll := d.d20.Roll()
	if roll < 1 {
		return nil, &core.InternalError{
			Message: "failed to generate random roll",
			Err:     err,
		}
	}

	dexMod := character.Dexterity.CalculateModifier()
	response := &core.RollInitiativeResponse{
		Result: dexMod + roll,
	}
	return response, nil
}
