package grpc

import (
	"context"

	"github.com/AranGarcia/droop/dnd/internal/ports/core/rules"

	dndpb "github.com/AranGarcia/droop/protoapis/dnd/v1"
)

func (s Server) RollInitiative(ctx context.Context, request *dndpb.RollInitiativeRequest) (*dndpb.RollInitiativeResponse, error) {
	apiRequest := rules.RollInitiativeRequest{ID: request.Id}
	apiResponse, err := s.DNDCore.RollInitiative(ctx, apiRequest)
	if err != nil {
		return nil, adaptCoreError(err)
	}

	response := &dndpb.RollInitiativeResponse{
		Result: int32(apiResponse.Result),
	}
	return response, nil
}
