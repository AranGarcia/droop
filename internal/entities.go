package internal

type Character struct {
	Level        int    `json:"level"`
	Name         string `json:"name"`
	HealthPoints int    `json:"health_points"`
	ArmorClass   int    `json:"armor_class"`

	Strength     int `json:"strength"`
	Dexterity    int `json:"dexterity"`
	Constitution int `json:"constitution"`
	Wisdom       int `json:"wisdom"`
	Charisma     int `json:"charisma"`
}
