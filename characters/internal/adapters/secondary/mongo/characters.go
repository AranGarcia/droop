package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/AranGarcia/droop/characters/internal/core/entities"
	"github.com/AranGarcia/droop/characters/internal/ports/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (c CharacterRepository) Create(ctx context.Context, entity entities.Character) (*entities.Character, error) {
	character := NewCharacterFromEntity(entity)
	character.CreatedAtNow()
	character.UpdatedAtNow()

	result, err := c.collection.InsertOne(ctx, character)
	if err != nil {
		return nil, fmt.Errorf("failed to insert character; %v", err)
	}
	created := entity.Copy()
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		created.ID = oid.Hex()
	}
	created.CreatedAt = character.CreatedAt
	created.UpdatedAt = character.UpdatedAt
	return &created, nil
}

func (c CharacterRepository) Retrieve(ctx context.Context, id string) (*entities.Character, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidID
	}
	filter := bson.M{
		"_id":        _id,
		"deleted_at": nil,
	}
	result := c.collection.FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		return nil, handleMongoError(err)
	}
	character := &Character{}
	if err = result.Decode(character); err != nil {
		return nil, fmt.Errorf("decoding error; %v", err)
	}
	entity := character.ToEntity()
	return &entity, nil
}

func (c CharacterRepository) Update(ctx context.Context, id string, fields repositories.CharacterFields) (*entities.Character, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id; %v", err)
	}
	filter := bson.M{
		"_id":        _id,
		"deleted_at": nil,
	}
	update := bson.M{"$set": characterFieldsToBSONMap(fields)}
	_, err = c.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, handleMongoError(err)
	}
	charcter, err := c.Retrieve(ctx, id)
	if err != nil {
		return nil, handleMongoError(err)
	}
	return charcter, nil
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

	_, err = c.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to delete; %v", err)
	}

	return nil
}

func (c CharacterRepository) List(ctx context.Context, offset, limit int) ([]entities.Character, error) {
	opts := options.Find().
		SetSkip(int64(offset)).
		SetLimit(int64(limit))
	cursor, err := c.collection.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, handleMongoError(err)
	}
	characters := make([]entities.Character, cursor.RemainingBatchLength())
	dbCharacter := Character{}
	i := 0
	for cursor.Next(ctx) {
		err = cursor.Decode(&dbCharacter)
		if err != nil {
			return nil, handleMongoError(err)
		}
		characters[i] = dbCharacter.ToEntity()
		i++
	}

	return characters, nil
}
