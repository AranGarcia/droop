package grpc

import (
	"context"

	"github.com/AranGarcia/droop/initiatives/internal/ports/core/turns"
	pb "github.com/AranGarcia/droop/proto/gen/initiatives"
)

func (s Server) RegisterTurn(ctx context.Context, request *pb.RegisterTurnRequest) (*pb.RegisterTurnResponse, error) {
	apiRequest := turns.RegisterRequest{}
	_, err := s.api.Register(ctx, apiRequest)
	if err != nil {
		return nil, handleAPIError(err)
	}
	response := &pb.RegisterTurnResponse{}
	return response, nil
}
