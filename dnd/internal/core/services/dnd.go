package services

import (
	"context"

	"github.com/AranGarcia/droop/dnd/internal/core/entities"
	"github.com/AranGarcia/droop/dnd/internal/ports/core"
	"github.com/AranGarcia/droop/dnd/internal/ports/core/dnd"
	"github.com/AranGarcia/droop/dnd/internal/ports/events"
	"github.com/AranGarcia/droop/dnd/internal/ports/external/characters"
)

type DND struct {
	d20 entities.Die

	// External ports
	characters characters.Port
	events     events.Producer
}

func NewDNDService(characters characters.Port) DND {
	d20 := entities.D20{}

	return DND{
		d20:        d20,
		characters: characters,
	}
}

func (d DND) RollInitiative(ctx context.Context, request dnd.RollInitiativeRequest) (*dnd.RollInitiativeResponse, error) {
	if request.ID == "" {
		return nil, core.ErrNoIDProvided
	}

	character, err := d.characters.Retrieve(ctx, request.ID)
	if err != nil {
		return nil, core.NewExternalServiceError("characters", err)
	}

	roll, err := d.d20.Roll()
	if err != nil {
		return nil, err
	}

	result := character.Dexterity.CalculateModifier() + roll
	response := &dnd.RollInitiativeResponse{
		Result: result,
	}

	// Domain events
	success := events.RollInitiativeSuccessMessage{
		CharacterID: request.ID,
		Result:      result,
	}
	d.events.RollInitiativeSuccess(ctx, success)
	return response, nil
}
