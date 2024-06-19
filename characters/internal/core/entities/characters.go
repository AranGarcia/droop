package entities

import "time"

// Base attributes shared among entities.
type Base struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type Character struct {
	Base
	ID           string `json:"id"`
	Level        int    `json:"level"`
	Name         string `json:"name"`
	HealthPoints int    `json:"health_points"`
	ArmorClass   int    `json:"armor_class"`
	Strength     int    `json:"strength"`
	Dexterity    int    `json:"dexterity"`
	Constitution int    `json:"constitution"`
	Intelligence int    `json:"intelligence"`
	Wisdom       int    `json:"wisdom"`
	Charisma     int    `json:"charisma"`
}
