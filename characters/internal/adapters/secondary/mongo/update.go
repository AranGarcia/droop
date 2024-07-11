package mongo

import (
	"github.com/AranGarcia/droop/characters/internal/ports/repositories"

	"go.mongodb.org/mongo-driver/bson"
)

func characterFieldsToBSONMap(fields repositories.CharacterFields) bson.M {
	m := bson.M{}
	if fields.Level != nil {
		m["level"] = *fields.Level
	}
	if fields.Name != nil {
		m["name"] = *fields.Name
	}
	if fields.HealthPoints != nil {
		m["health_points"] = *fields.HealthPoints
	}
	if fields.ArmorClass != nil {
		m["armor_class"] = *fields.ArmorClass
	}
	if fields.Strength != nil {
		m["strength"] = *fields.Strength
	}
	if fields.Dexterity != nil {
		m["dexterity"] = *fields.Dexterity
	}
	if fields.Constitution != nil {
		m["constitution"] = *fields.Constitution
	}
	if fields.Intelligence != nil {
		m["intelligence"] = *fields.Intelligence
	}
	if fields.Wisdom != nil {
		m["wisdom"] = *fields.Wisdom
	}
	if fields.Charisma != nil {
		m["charisma"] = *fields.Charisma
	}
	return m
}
