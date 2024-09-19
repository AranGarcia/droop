package services

import (
	"context"
	"crypto/rand"
	"math/big"

	"github.con/AranGarcia/droop/dnd/ports/core"
	"github.con/AranGarcia/droop/dnd/ports/external/characters"
)

var (
	maxRoll = big.NewInt(20)
)

type DND struct {
	// External ports
	characters characters.Port
}

func (d DND) RollInitiative(ctx context.Context, request core.RollInitiativeRequest) (*core.RollInitiativeResponse, error) {
	character, err := d.characters.Retrieve(ctx, request.ID)
	if err != nil {
		return nil, core.NewExternalServiceError("characters", err)
	}

	roll, err := rand.Int(rand.Reader, maxRoll)
	if err != nil {
		return nil, &core.InternalError{
			Message: "failed to generate random roll",
			Err:     err,
		}
	}

	dexMod := character.Dexterity.CalculateModifier()
	response := &core.RollInitiativeResponse{
		Result: dexMod + int(roll.Int64()),
	}
	return response, nil
}
