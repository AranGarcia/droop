package redistools

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Addr string
}

func (c Config) InitializeClient() (*redis.Client, error) {
	opts := &redis.Options{}
	client := redis.NewClient(opts)

	ctx := context.TODO()
	status := client.Ping(ctx)
	if status.Err() != nil { // TODO: confirm that this is how to ping successfully
		return nil, status.Err()
	}

	return client, nil
}
