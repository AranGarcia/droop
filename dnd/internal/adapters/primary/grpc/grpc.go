package grpc

import (
	"context"

	"github.com/AranGarcia/droop/dnd/internal/ports/core"

	dndpb "github.com/AranGarcia/droop/proto/gen/dnd"
)

func (s Server) RollInitiative(ctx context.Context, request *dndpb.RollInitiativeRequest) (*dndpb.RollInitiativeResponse, error) {
	apiRequest := core.RollInitiativeRequest{ID: request.Id}
	apiResponse, err := s.DNDCore.RollInitiative(ctx, apiRequest)
	if err != nil {
		return nil, adaptCoreError(err)
	}

	response := &dndpb.RollInitiativeResponse{
		Result: int32(apiResponse.Result),
	}
	return response, nil
}
