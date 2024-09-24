package grpc

import (
	"context"

	"github.con/AranGarcia/droop/dnd/ports/core"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	dndpb "github.com/AranGarcia/droop/proto/gen/dnd"
)

func (s Server) RollInitiative(ctx context.Context, request *dndpb.RollInitiativeRequest) (*dndpb.RollInitiativeResponse, error) {
	if request.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "character ID is required")
	}

	apiRequest := core.RollInitiativeRequest{ID: request.Id}
	apiResponse, err := s.DNDCore.RollInitiative(ctx, apiRequest)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "failed to roll initiative; %v", err)
	}

	response := &dndpb.RollInitiativeResponse{
		Result: int32(apiResponse.Result),
	}
	return response, nil
}
