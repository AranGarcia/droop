package characters

import (
	"github.com/AranGarcia/droop/dnd/internal/core/entities"

	characterspb "github.com/AranGarcia/droop/proto/gen/characters"
)

func CharacterCoreFromExternal(character *characterspb.Character) (*entities.Character, error) {
	dex, err := entities.NewAbilityScore(int(character.Dexterity))
	if err != nil {
		return nil, ErrInvalidAttributeType
	}

	coreCharacter := &entities.Character{
		ID:        character.Id,
		Name:      character.Name,
		Dexterity: *dex,
	}
	return coreCharacter, nil
}
