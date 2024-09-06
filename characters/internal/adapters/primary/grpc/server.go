package grpc

import (
	"net"

	"google.golang.org/grpc"
)

// Server is a gRPC adapter for the character service.
type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	s := &Server{
		addr: addr,
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
	return grpcServer.Serve(lis)
}
