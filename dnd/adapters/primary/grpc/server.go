package grpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	dndpb "github.com/AranGarcia/droop/proto/gen/dnd"
)

type Server struct {
	dndpb.UnimplementedAPIServer

	// Addr is the connection string in the format "HOST:PORT"
	addr string
}

type Config struct {
	// Addr is the connection string in the format "HOST:PORT"
	Addr string
}

func NewServer(c Config) Server {
	s := Server{
		addr: c.Addr,
	}
	return s
}

func (s Server) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to bind addres %q; %s", s.addr, err)
	}
	grpcServer := grpc.NewServer()
	dndpb.RegisterAPIServer(grpcServer, s)
	return grpcServer.Serve(lis)
}
