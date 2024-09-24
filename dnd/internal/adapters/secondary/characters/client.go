package characters

import (
	"fmt"

	"google.golang.org/grpc"

	characterspb "github.com/AranGarcia/droop/proto/gen/characters"
)

type Client struct {
	grpcClient characterspb.APIClient
}

type Config struct {
	Addr string
}

func NewClient(config Config) (*Client, error) {
	conn, err := grpc.NewClient(config.Addr)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize grpc client; %v", err)
	}

	grpcClient := characterspb.NewAPIClient(conn)

	c := &Client{
		grpcClient: grpcClient,
	}
	return c, nil
}
