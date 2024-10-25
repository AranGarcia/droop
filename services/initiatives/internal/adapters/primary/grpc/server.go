package grpc

import (
	"context"
	"net"

	"google.golang.org/grpc"

	"github.com/AranGarcia/droop/initiatives/internal/ports/core/tables"

	initiativespb "github.com/AranGarcia/droop/proto/gen/initiatives"
)

type Server struct {
	initiativespb.UnimplementedAPIServer

	addr      string
	tablesAPI tables.API
}

func NewServer(addr string, tables tables.API) *Server {
	return &Server{
		addr:      addr,
		tablesAPI: tables,
	}
}

func (s Server) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	return grpcServer.Serve(lis)
}

func (t Server) StartTracking(ctx context.Context, request *initiativespb.StartTrackingRequest) (*initiativespb.StartTrackingResponse, error) {
	apiRequest := tables.StartTrackingRequest{CampaignID: request.CampaignId}
	_, err := t.tablesAPI.StartTracking(ctx, apiRequest)
	if err != nil {
		// TODO: handle core errors
		return nil, err
	}
	response := &initiativespb.StartTrackingResponse{}
	return response, nil
}

func (t Server) StopTracking(ctx context.Context, request *initiativespb.StopTrackingRequest) (*initiativespb.StopTrackingResponse, error) {
	apiRequest := tables.StopTrackingRequest{CampaignID: request.CampaignId}
	_, err := t.tablesAPI.StopTracking(ctx, apiRequest)
	if err != nil {
		// TODO: handle core errors
		return nil, err
	}
	response := &initiativespb.StopTrackingResponse{}
	return response, nil
}
