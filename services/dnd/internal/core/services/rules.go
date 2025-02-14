package services

import (
	"context"
	"log"

	"github.com/AranGarcia/droop/dnd/internal/core/entities"
	"github.com/AranGarcia/droop/dnd/internal/ports/core"
	"github.com/AranGarcia/droop/dnd/internal/ports/core/rules"
	"github.com/AranGarcia/droop/dnd/internal/ports/events"
	"github.com/AranGarcia/droop/dnd/internal/ports/external/characters"
)

type DND struct {
	d20 entities.Die

	// External ports
	characters characters.Port
	events     events.Port
}

func NewDNDService(characters characters.Port, events events.Port) DND {
	d20 := entities.D20{}

	return DND{
		d20:        d20,
		characters: characters,
		events:     events,
	}
}

func (d DND) RollInitiative(ctx context.Context, request rules.RollInitiativeRequest) (*rules.RollInitiativeResponse, error) {
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
	response := &rules.RollInitiativeResponse{
		Result: result,
	}

	// Domain events
	success := events.RollInitiativeSuccessMessage{
		CharacterID: request.ID,
		Result:      result,
	}
	err = d.events.RollInitiativeSuccess(ctx, success)
	if err != nil {
		log.Printf("failed to produce: %v", err)
	}
	return response, nil
}
