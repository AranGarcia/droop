package services

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.con/AranGarcia/droop/dnd/ports/api"
	"github.con/AranGarcia/droop/dnd/ports/external/characters"
)

var (
	maxRoll = big.NewInt(20)
)

type DND struct {
	// External ports
	characters characters.Port
}

func (d DND) RollInitiative(ctx context.Context, request api.RollInitiativeRequest) (*api.RollInitiativeResponse, error) {
	character, err := d.characters.Retrieve(ctx, request.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve character data; %v", err)
	}

	roll, err := rand.Int(rand.Reader, maxRoll)
	if err != nil {
		return nil, fmt.Errorf("failed to roll; %v", err)
	}

	response := &api.RollInitiativeResponse{
		Result: character.Dexterity + int(roll.Int64()),
	}
	return response, nil
}
