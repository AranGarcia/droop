package characters

import (
	"context"

	"github.com/AranGarcia/droop/dnd/internal/core/entities"

	characterspb "github.com/AranGarcia/droop/protoapis/characters/v1"
)

// Retrieve a Character from the service.
func (c Client) Retrieve(ctx context.Context, id string) (*entities.Character, error) {
	request := characterspb.RetrieveRequest{Id: id}
	response, err := c.grpcClient.Retrieve(ctx, &request)
	if err != nil {
		return nil, translateGRPCError(err)
	}

	character, err := CharacterCoreFromExternal(response.Character)
	if err != nil {
		return nil, err
	}
	return character, nil
}
