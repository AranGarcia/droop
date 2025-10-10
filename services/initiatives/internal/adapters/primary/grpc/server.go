package grpc

import (
	"net"

	"google.golang.org/grpc"

	"github.com/AranGarcia/droop/initiatives/internal/ports/core/turns"

	pb "github.com/AranGarcia/droop/protoapis/initiatives/v1"
)

type Server struct {
	pb.UnimplementedServiceServer

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
	pb.RegisterServiceServer(grpcServer, s)
	return grpcServer.Serve(lis)
}
