package grpc

import (
	"net"

	"google.golang.org/grpc"

	"github.com/AranGarcia/droop/characters/internal/ports/api"

	characterspb "github.com/AranGarcia/droop/proto/gen/characters"
)

// Server is a gRPC adapter for the character service.
type Server struct {
	// Unimplemented gRPC server.
	characterspb.UnimplementedAPIServer

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
	characterspb.RegisterAPIServer(grpcServer, s)
	return grpcServer.Serve(lis)
}
