package grpc

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	return nil, status.Error(codes.Unimplemented, "not implemented yet")
}

func (t Server) StopTracking(ctx context.Context, request *initiativespb.StopTrackingRequest) (*initiativespb.StopTrackingResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented yet")
}
