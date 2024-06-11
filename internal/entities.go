package internal

type Character struct {
	ID           string `json:"id"`
	Level        int    `json:"level" bson:"level" validate:"gte=1,lte=20"`
	Name         string `json:"name" bson:"name" validate:"required"`
	HealthPoints int    `json:"health_points" bson:"health_points"`
	ArmorClass   int    `json:"armor_class" bson:"armor_class"`

	Strength     int `json:"strength" bson:"strength" validate:"gte=1,lte=20"`
	Dexterity    int `json:"dexterity" bson:"dexterity" validate:"gte=1,lte=20"`
	Constitution int `json:"constitution" bson:"constitution" validate:"gte=1,lte=20"`
	Intelligence int `json:"intelligence" bson:"intelligence" validate:"gte=1,lte=20"`
	Wisdom       int `json:"wisdom" bson:"wisdom,lte=20"`
	Charisma     int `json:"charisma" bson:"charisma" validate:"gte=1,lte=20"`
}

// Validate all of the fields and return an error if an value violates the fields' constraints.
func (c Character) Validate() error { return validate.Struct(c) }
