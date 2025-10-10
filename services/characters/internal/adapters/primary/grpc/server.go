package grpc

import (
	"net"

	"google.golang.org/grpc"

	"github.com/AranGarcia/droop/characters/internal/core/api"

	characterspb "github.com/AranGarcia/droop/protoapis/characters/v1"
)

// Server is a gRPC adapter for the character service.
type Server struct {
	// Unimplemented gRPC server.
	characterspb.UnimplementedServiceServer

	addr             string
	characterService api.Characters
}

func NewServer(addr string, characterService api.Characters) *Server {
	s := &Server{
		addr:             addr,
		characterService: characterService,
	}
	return s
}

// Run the server and start listening for new connections.
func (s Server) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	characterspb.RegisterServiceServer(grpcServer, s)
	return grpcServer.Serve(lis)
}
