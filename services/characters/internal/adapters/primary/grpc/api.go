package grpc

import (
	"github.com/AranGarcia/droop/characters/internal/core/api"
	"github.com/AranGarcia/droop/characters/internal/core/entities"

	characterspb "github.com/AranGarcia/droop/proto/gen/characters"
)

func CreateRequestToAPI(request *characterspb.CreateRequest) api.CreateCharacterRequest {
	if request == nil {
		return api.CreateCharacterRequest{}
	}

	return api.CreateCharacterRequest{
		Class:        entities.ClassName(request.Class),
		Level:        entities.Level(request.Level),
		Name:         request.Name,
		MaxHealth:    int(request.MaxHealth),
		ArmorClass:   int(request.ArmorClass),
		Strength:     entities.AbilityScore(request.Strength),
		Dexterity:    entities.AbilityScore(request.Dexterity),
		Constitution: entities.AbilityScore(request.Constitution),
		Intelligence: entities.AbilityScore(request.Intelligence),
		Wisdom:       entities.AbilityScore(request.Wisdom),
		Charisma:     entities.AbilityScore(request.Charisma),
	}
}

func CreateResponseToProto(response *api.CreateCharacterResponse) *characterspb.CreateResponse {
	if response == nil {
		return nil
	}

	return &characterspb.CreateResponse{
		Character: &characterspb.Character{
			Id:           response.Character.ID,
			Class:        string(response.Character.Class),
			Level:        int32(response.Character.Level),
			Name:         response.Character.Name,
			ArmorClass:   int32(response.Character.ArmorClass),
			Strength:     int32(response.Character.Strength),
			Dexterity:    int32(response.Character.Dexterity),
			Constitution: int32(response.Character.Constitution),
			Intelligence: int32(response.Character.Intelligence),
			Wisdom:       int32(response.Character.Wisdom),
			Charisma:     int32(response.Character.Charisma),
		},
	}
}

func RetrieveRequestToAPI(request *characterspb.RetrieveRequest) api.RetrieveCharacterRequest {
	if request == nil {
		return api.RetrieveCharacterRequest{}
	}
	return api.RetrieveCharacterRequest{
		ID: request.Id,
	}
}

func RetrieveResponseToProto(response *api.RetrieveCharacterResponse) *characterspb.RetrieveResponse {
	if response == nil {
		return nil
	}

	return &characterspb.RetrieveResponse{
		Character: &characterspb.Character{
			Id:           response.Character.ID,
			Class:        string(response.Character.Class),
			Level:        int32(response.Character.Level),
			Name:         response.Character.Name,
			ArmorClass:   int32(response.Character.ArmorClass),
			Strength:     int32(response.Character.Strength),
			Dexterity:    int32(response.Character.Dexterity),
			Constitution: int32(response.Character.Constitution),
			Intelligence: int32(response.Character.Intelligence),
			Wisdom:       int32(response.Character.Wisdom),
			Charisma:     int32(response.Character.Charisma),
		},
	}
}

func CharacterPBFromCore(character entities.Character) *characterspb.Character {
	return &characterspb.Character{
		Id:           character.ID,
		Class:        string(character.Class),
		Level:        int32(character.Level),
		Name:         character.Name,
		ArmorClass:   int32(character.ArmorClass),
		Strength:     int32(character.Strength),
		Dexterity:    int32(character.Dexterity),
		Constitution: int32(character.Constitution),
		Intelligence: int32(character.Intelligence),
		Wisdom:       int32(character.Wisdom),
		Charisma:     int32(character.Charisma),
	}
}
