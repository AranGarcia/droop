package grpc

import (
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
