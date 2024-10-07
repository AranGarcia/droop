package grpc

import (
	"github.com/AranGarcia/droop/characters/internal/core/entities"

	characterspb "github.com/AranGarcia/droop/proto/gen/characters"
)

func CharacterPBFromCore(character entities.Character) *characterspb.Character {
	return &characterspb.Character{
		Id:           character.ID,
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
