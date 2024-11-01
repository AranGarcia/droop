package redistools

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Addr     string
	Password string
	DB       int
}

func (c Config) InitializeClient() (*redis.Client, error) {
	opts := &redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	}
	client := redis.NewClient(opts)

	ctx := context.TODO()
	status := client.Ping(ctx)
	if status.Err() != nil {
		return nil, fmt.Errorf("redis client error; %v", status.Err())
	}

	return client, nil
}
