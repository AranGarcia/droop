package internal

type Character struct {
	ID           string `json:"id"`
	Level        int    `json:"level" bson:"level"`
	Name         string `json:"name" bson:"name"`
	HealthPoints int    `json:"health_points" bson:"health_points"`
	ArmorClass   int    `json:"armor_class" bson:"armor_class"`
	Dexterity    int    `json:"dexterity" bson:"dexterity"`
	Constitution int    `json:"constitution" bson:"constitution"`
	Wisdom       int    `json:"wisdom" bson:"wisdom"`
	Charisma     int    `json:"charisma" bson:"charisma"`
}

// Validate all of the fields and return an error if an value violates the fields' constraints.
func (c Character) Validate() error { return validate.Struct(c) }
