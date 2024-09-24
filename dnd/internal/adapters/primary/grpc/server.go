package grpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/AranGarcia/droop/dnd/internal/ports/core"

	dndpb "github.com/AranGarcia/droop/proto/gen/dnd"
)

// Config for the server.
type Config struct {
	// Addr is the connection string in the format "HOST:PORT"
	Addr string

	// DND is a port to the core service.
	DND core.DND
}

// Server is a primary port implemented with GRPC.
type Server struct {
	dndpb.UnimplementedAPIServer

	// Addr is the connection string in the format "HOST:PORT"
	addr string

	// DND is a port to the core service.
	DNDCore core.DND
}

// NewServer constructs the server using the configuration parameters.
func NewServer(c Config) Server {
	s := Server{
		addr:    c.Addr,
		DNDCore: c.DND,
	}
	return s
}

// Run the server.
func (s Server) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to bind addres %q; %s", s.addr, err)
	}
	grpcServer := grpc.NewServer()
	dndpb.RegisterAPIServer(grpcServer, s)
	return grpcServer.Serve(lis)
}
