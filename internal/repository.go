package internal

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// database within the MongoDB.
	database = "droop"
	// collection where characters will be stored.
	collection = "characters"
)

// CharacterRepository provides access to the repository where character are stored.
type CharacterRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

// NewCharacterRepository initializes a new repository with a MongoDB client.
func NewCharacterRepository(config MongoConfig) (*CharacterRepository, error) {
	client, err := config.InitializeClient()
	if err != nil {
		return nil, fmt.Errorf("failed to connect; %v", err)
	}

	repo := &CharacterRepository{
		client:     client,
		collection: client.Database(database).Collection(collection),
	}
	return repo, nil
}

// CreateCharacter inserts the input character into the database. The same character is returned,
// but with its ID set.
func (c CharacterRepository) CreateCharacter(ctx context.Context, character Character) (*Character, error) {
	// Deprecated: use CharacterRepository.collection instead
	result, err := c.collection.InsertOne(ctx, character)
	if err != nil {
		return nil, fmt.Errorf("failed to insert character; %v", err)
	}
	character.ID = fmt.Sprint(result.InsertedID)
	return &character, nil
}

func (c CharacterRepository) RetrieveCharacter(ctx context.Context, id string) (*Character, error) {
	result := c.collection.FindOne(ctx, bson.M{"_id": id})
	if result.Err() != nil {
		return nil, fmt.Errorf("query error; %v", result.Err())
	}
	character := &Character{}
	if err := result.Decode(character); err != nil {
		return nil, fmt.Errorf("decoding error; %v", err)
	}
	return character, nil
}
