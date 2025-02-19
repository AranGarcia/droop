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
		Level:        int(request.Level),
		Name:         request.Name,
		HealthPoints: int(request.HealthPoints),
		ArmorClass:   int(request.ArmorClass),
		Strength:     int(request.Strength),
		Dexterity:    int(request.Dexterity),
		Constitution: int(request.Constitution),
		Intelligence: int(request.Intelligence),
		Wisdom:       int(request.Wisdom),
		Charisma:     int(request.Charisma),
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
			HealthPoints: int32(response.Character.HealthPoints),
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
			HealthPoints: int32(response.Character.HealthPoints),
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
		HealthPoints: int32(character.HealthPoints),
		ArmorClass:   int32(character.ArmorClass),
		Strength:     int32(character.Strength),
		Dexterity:    int32(character.Dexterity),
		Constitution: int32(character.Constitution),
		Intelligence: int32(character.Intelligence),
		Wisdom:       int32(character.Wisdom),
		Charisma:     int32(character.Charisma),
	}
}
