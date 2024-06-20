package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AranGarcia/droop/characters/internal/core/entities"
	"github.com/AranGarcia/droop/characters/internal/ports/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// database within the MongoDB.
	database = "droop"
	// collection where characters will be stored.
	collection = "characters"
)

type CharacterRepository struct {
	collection *mongo.Collection
}

func NewCharacterRepository(config Config) (*CharacterRepository, error) {
	client, err := config.InitializeClient()
	if err != nil {
		return nil, fmt.Errorf("failed to connect; %v", err)
	}
	repo := &CharacterRepository{
		collection: client.Database(database).Collection(collection),
	}
	return repo, nil
}

func (c CharacterRepository) Create(ctx context.Context, character entities.Character) error {
	// TODO: decide whether caller sets ID or
	_, err := c.collection.InsertOne(ctx, character)
	if err != nil {
		return fmt.Errorf("failed to insert character; %v", err)
	}
	return nil
}

func (c CharacterRepository) Retrieve(ctx context.Context, id string) (*entities.Character, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id; %v", err)
	}
	filter := bson.M{
		"_id":        _id,
		"deleted_at": nil,
	}
	result := c.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		return nil, fmt.Errorf("query error; %v", result.Err())
	}
	character := &entities.Character{}
	if err = result.Decode(character); err != nil {
		return nil, fmt.Errorf("decoding error; %v", err)
	}
	return character, nil
}

func (c CharacterRepository) Update(_ context.Context, _ string, _ repositories.CharacterFields) (*entities.Character, error) {
	panic("not implemented") // TODO: Implement
}

func (c CharacterRepository) Delete(ctx context.Context, id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid ID; %v", err)
	}
	filter := bson.D{{Key: "_id", Value: _id}}
	update := bson.D{
		{
			Key: "$set", Value: bson.D{
				{
					Key:   "deleted_at",
					Value: time.Now(),
				},
			},
		},
	}

	result, err := c.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to delete; %v", err)
	}

	log.Println(result)
	return nil
}
