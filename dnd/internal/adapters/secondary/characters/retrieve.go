package characters

import (
	"context"
	"fmt"

	"github.con/AranGarcia/droop/dnd/internal/core/entities"

	characterspb "github.com/AranGarcia/droop/proto/gen/characters"
)

// Retrieve a Character from the service.
func (c Client) Retrieve(ctx context.Context, id string) (*entities.Character, error) {
	request := characterspb.RetrieveRequest{Id: id}
	response, err := c.grpcClient.Retrieve(ctx, &request)
	if err != nil {
		return nil, fmt.Errorf("grpcClient retrieve failure; %v", err)
	}

	character := &entities.Character{
		ID: response.Character.Id,
	}
	return character, nil
}
