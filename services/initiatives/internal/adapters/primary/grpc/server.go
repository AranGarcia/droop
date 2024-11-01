package grpc

import (
	"net"

	"google.golang.org/grpc"

	"github.com/AranGarcia/droop/initiatives/internal/ports/core/turns"
)

type Server struct {
	addr string
	api  turns.API
}

func NewServer(addr string, turns turns.API) *Server {
	return &Server{
		addr: addr,
		api:  turns,
	}
}

func (s Server) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	// initiativespb.RegisterAPIServer(grpcServer, s)
	return grpcServer.Serve(lis)
}
