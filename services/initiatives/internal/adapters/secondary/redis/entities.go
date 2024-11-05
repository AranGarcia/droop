package redis

import (
	"fmt"
	"strings"

	"github.com/redis/go-redis/v9"

	"github.com/AranGarcia/droop/initiatives/internal/core/entities"
)

func ZMemberToTurn(z redis.Z) (entities.Turn, error) {
	memberStr, ok := z.Member.(string)
	if !ok {
		return entities.Turn{}, fmt.Errorf("corrupt data: member not a string")
	}
	characterData := strings.SplitN(memberStr, ":", 2)
	if len(characterData) < 2 {
		return entities.Turn{}, fmt.Errorf("corrupt data: should have id and name")
	}
	turn := entities.Turn{
		CharacterID:   characterData[0],
		CharacterName: characterData[1],
	}
	return turn, nil
}
