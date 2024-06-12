package internal

import (
	"context"
	"fmt"
	"log"
	"time"

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

func IntPtr(i int) *int { return &i }

func StrPtr(s string) *string { return &s }

// CharacterRepository provides access to the repository where character are stored.
type CharacterRepository struct {
	client     *mongo.Client // TODO: figure out if this is required or if it should be deleted.
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
	result, err := c.collection.InsertOne(ctx, character)
	if err != nil {
		return nil, fmt.Errorf("failed to insert character; %v", err)
	}
	_id := result.InsertedID.(primitive.ObjectID)
	character.ID = _id.Hex()
	return &character, nil
}

func (c CharacterRepository) RetrieveCharacter(ctx context.Context, id string) (*Character, error) {
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
	character := &Character{}
	if err = result.Decode(character); err != nil {
		return nil, fmt.Errorf("decoding error; %v", err)
	}
	return character, nil
}

type UpdateFields struct {
	Level        *int    `json:"level"`
	Name         *string `json:"name"`
	HealthPoints *int    `json:"health_points"`
	ArmorClass   *int    `json:"armor_class"`
	Strength     *int    `json:"strength"`
	Dexterity    *int    `json:"dexterity"`
	Constitution *int    `json:"constitution"`
	Intelligence *int    `json:"intelligence"`
	Wisdom       *int    `json:"wisdom"`
	Charisma     *int    `json:"charisma"`
}

func (u UpdateFields) ToBsonMap() bson.M {
	data := bson.M{}
	if u.Level != nil {
		data["level"] = *u.Level
	}
	if u.Name != nil {
		data["name"] = *u.Name
	}
	if u.HealthPoints != nil {
		data["health_points"] = *u.HealthPoints
	}
	if u.ArmorClass != nil {
		data["armor_class"] = *u.ArmorClass
	}
	if u.Strength != nil {
		data["strength"] = *u.Strength
	}
	if u.Dexterity != nil {
		data["dexterity"] = *u.Dexterity
	}
	if u.Constitution != nil {
		data["constitution"] = *u.Constitution
	}
	if u.Intelligence != nil {
		data["intelligence"] = *u.Intelligence
	}
	if u.Wisdom != nil {
		data["wisdom"] = *u.Wisdom
	}
	if u.Charisma != nil {
		data["charisma"] = *u.Charisma
	}
	return data
}

func (c CharacterRepository) Update(ctx context.Context, id string, fields *UpdateFields) (*Character, error) {
	return &Character{}, nil
}

func (c CharacterRepository) DeleteCharacter(ctx context.Context, id string) error {
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
