package grpc

import (
	"context"

	pb "github.com/AranGarcia/droop/protoapis/proto/initiatives"
)

func (s Server) RegisterTurn(ctx context.Context, request *pb.RegisterTurnRequest) (*pb.RegisterTurnResponse, error) {
	apiRequest := RegisterRequestToAPI(request)
	_, err := s.api.Register(ctx, apiRequest)
	if err != nil {
		return nil, handleAPIError(err)
	}
	response := &pb.RegisterTurnResponse{}
	return response, nil
}

// ListTurns is an adapter for the core method.
func (s Server) ListTurns(ctx context.Context, request *pb.ListTurnsRequest) (*pb.ListTurnsResponse, error) {
	apiRequest := ListTurnsRequestToAPI(request)
	apiResponse, err := s.api.ListTurns(ctx, apiRequest)
	return ListTurnsResponseToAdapter(apiResponse), handleAPIError(err)
}

func (s Server) ClearAll(ctx context.Context, request *pb.ClearAllRequest) (*pb.ClearAllResponse, error) {
	apiRequest := ClearAllRequestToAPI(request)
	_, err := s.api.ClearAll(ctx, apiRequest)
	if err != nil {
		return nil, handleAPIError(err)
	}
	response := &pb.ClearAllResponse{}
	return response, nil
}
