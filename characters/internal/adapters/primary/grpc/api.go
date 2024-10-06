package grpc

import (
	"github.com/AranGarcia/droop/characters/internal/ports/api"

	characterspb "github.com/AranGarcia/droop/proto/gen/characters"
)

func CreateRequestToAPI(request *characterspb.CreateRequest) api.CreateCharacterRequest {
	if request == nil {
		return api.CreateCharacterRequest{}
	}

	return api.CreateCharacterRequest{
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
