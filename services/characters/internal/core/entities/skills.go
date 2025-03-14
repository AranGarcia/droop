package entities

type Abilities map[string]Skill

// Skill is the character's capability in a certain area.
type Skill struct {
	// Name       string `json:"name" validate:"oneof=acrobatics animal_handling arcana athletics deception history insight intimidation investigation medicine nature perception performance persuasion religion sleight_of_hand stealth survival"`
	Proficient bool `json:"proficient"`
	Bonus      int  `json:"bonus"`
}

const (
	AcrobaticsSkill     string = "acrobatics"
	AnimalHandlingSkill string = "animal_handling"
	ArcanaSkill         string = "arcana"
	AthleticsSkill      string = "athletics"
	DeceptionSkill      string = "deception"
	HistorySkill        string = "history"
	InsightSkill        string = "insight"
	IntimidationSkill   string = "intimidation"
	InvestigationSkill  string = "investigation"
	MedicineSkill       string = "medicine"
	NatureSkill         string = "nature"
	PerceptionSkill     string = "perception"
	PerformanceSkill    string = "performance"
	PersuasionSkill     string = "persuasion"
	ReligionSkill       string = "religion"
	SleightOfHandSkill  string = "sleight_of_hand"
	StealthSkill        string = "stealth"
	SurvivalSkill       string = "survival"
)
