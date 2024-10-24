package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config contains the configuration parameters to connect to MongoDB.
type Config struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
	Timeout  time.Duration
}

// InitializeClient returns an initialized Mongo Client. It also performs a ping to validate it could
// connect properly.
func (m Config) InitializeClient() (*mongo.Client, error) {
	connectionString := m.ConnectionString()
	clientOpt := options.Client().ApplyURI(connectionString)
	clientOpt.Timeout = &m.Timeout

	ctx := context.TODO()
	client, err := mongo.Connect(ctx, clientOpt)
	if err != nil {
		return nil, fmt.Errorf("failed to connect; %v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed ping to database; %v", err)
	}

	return client, nil
}

// ConnectionString built from the configuration parameters.
func (m Config) ConnectionString() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", m.User, m.Password, m.Host, m.Port, m.Database)
}
