package grpc

import (
	"context"
	"fmt"

	"github.com/AranGarcia/droop/characters/internal/ports/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	characterspb "github.com/AranGarcia/droop/proto/gen/characters"
)

func (s Server) Create(ctx context.Context, request *characterspb.CreateRequest) (*characterspb.CreateResponse, error) {
	apiRequest := CreateRequestToAPI(request)
	apiResponse, err := s.characterService.Create(ctx, apiRequest)
	if err != nil {
		// TODO: convert API error
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to create; %v", err))
	}
	response := CreateResponseToProto(apiResponse)
	return response, nil
}

// Retrieve a Character.
func (s Server) Retrieve(ctx context.Context, request *characterspb.RetrieveRequest) (*characterspb.RetrieveResponse, error) {
	if request == nil || request.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "character ID is required")
	}

	apiRequest := api.RetrieveCharacterRequest{
		ID: request.Id,
	}
	apiResponse, err := s.characterService.Retrieve(ctx, apiRequest)
	if err != nil {
		// TODO: convert API error
		return nil, status.Error(codes.Internal, fmt.Sprintf("character service error; %v", err))
	}

	response := &characterspb.RetrieveResponse{
		Character: &characterspb.Character{
			Id: apiResponse.Character.ID,
		},
	}
	return response, nil
}
