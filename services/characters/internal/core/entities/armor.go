package entities

const (
	// Light Armor

	PaddedArmor        = "padded"
	LeatherArmor       = "leather"
	StudedLeatherArmor = "studded_leather"

	// Medium armor

	HideArmor        = "hide"
	ChainShirt       = "chain_shirt"
	ScaleMailArmor   = "scale_mail"
	BreastPlateArmor = "breast_plate"
	HalfPlateArmor   = "half_plate"

	// Heavy Armor

	RingMailArmor  = "ring_mail"
	ChainMailArmor = "chain_mail"
	SplintArmor    = "splint"
	PlateArmor     = "plate"
)

type Armor struct {
	Name                string
	ArmorClass          int
	StealthDisadvantage bool
}

func NewArmor(name string) *Armor {
	armor := &Armor{Name: name}
	switch name {
	case PaddedArmor:
		armor.ArmorClass = 11
		armor.StealthDisadvantage = true
	case LeatherArmor:
		armor.ArmorClass = 11
	case StudedLeatherArmor:
		armor.ArmorClass = 12

	// Medium armors
	case HideArmor:
		armor.ArmorClass = 12
	case ChainShirt:
		armor.ArmorClass = 13
	case ScaleMailArmor:
		armor.ArmorClass = 14
		armor.StealthDisadvantage = true
	case BreastPlateArmor:
		armor.ArmorClass = 14
	case HalfPlateArmor:
		armor.ArmorClass = 15
		armor.StealthDisadvantage = true

	// Heavy armors
	case RingMailArmor:
		armor.ArmorClass = 14
		armor.StealthDisadvantage = true
	case ChainMailArmor:
		armor.ArmorClass = 16
		armor.StealthDisadvantage = true
	case SplintArmor:
		armor.ArmorClass = 17
		armor.StealthDisadvantage = true
	case PlateArmor:
		armor.ArmorClass = 18
		armor.StealthDisadvantage = true
	default:
		armor.Name = "unknown"
	}
	return armor
}
