package mongo

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/AranGarcia/droop/characters/internal/ports/repositories"
)

func characterFieldsToBSONMap(fields repositories.CharacterFields) bson.M {
	m := bson.M{}
	if fields.Class != nil {
		m["class"] = *fields.Class
	}
	if fields.Level != nil {
		m["level"] = *fields.Level
	}
	if fields.Name != nil {
		m["name"] = *fields.Name
	}
	if fields.MaxHealth != nil {
		m["max_health"] = *fields.MaxHealth
	}
	if fields.CurrentHealth != nil {
		m["current_health"] = *fields.CurrentHealth
	}
	if fields.TempHealth != nil {
		m["temp_health"] = *fields.TempHealth
	}
	if fields.HealthPoints != nil { // TODO: remove
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
