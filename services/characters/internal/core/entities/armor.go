package entities

const (
	// Weights
	LightArmorCategory  = "light"
	MediumArmorCategory = "medium"
	HeavyArmorCategory  = "heavy"

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
	WeightCategory      string
	ArmorClass          int
	StealthDisadvantage bool
}

func NewArmor(name string) *Armor {
	armor := &Armor{Name: name}
	switch name {
	case PaddedArmor:
		armor.WeightCategory = LightArmorCategory
		armor.ArmorClass = 11
		armor.StealthDisadvantage = true
	case LeatherArmor:
		armor.WeightCategory = LightArmorCategory
		armor.ArmorClass = 11
	case StudedLeatherArmor:
		armor.WeightCategory = LightArmorCategory
		armor.ArmorClass = 12

	// Medium armors
	case HideArmor:
		armor.WeightCategory = MediumArmorCategory
		armor.ArmorClass = 12
	case ChainShirt:
		armor.WeightCategory = MediumArmorCategory
		armor.ArmorClass = 13
	case ScaleMailArmor:
		armor.WeightCategory = MediumArmorCategory
		armor.ArmorClass = 14
		armor.WeightCategory = MediumArmorCategory
		armor.StealthDisadvantage = true
	case BreastPlateArmor:
		armor.WeightCategory = MediumArmorCategory
		armor.ArmorClass = 14
	case HalfPlateArmor:
		armor.WeightCategory = MediumArmorCategory
		armor.ArmorClass = 15
		armor.StealthDisadvantage = true

	// Heavy armors
	case RingMailArmor:
		armor.WeightCategory = HeavyArmorCategory
		armor.ArmorClass = 14
		armor.StealthDisadvantage = true
	case ChainMailArmor:
		armor.WeightCategory = HeavyArmorCategory
		armor.ArmorClass = 16
		armor.StealthDisadvantage = true
	case SplintArmor:
		armor.WeightCategory = HeavyArmorCategory
		armor.ArmorClass = 17
		armor.StealthDisadvantage = true
	case PlateArmor:
		armor.WeightCategory = HeavyArmorCategory
		armor.ArmorClass = 18
		armor.StealthDisadvantage = true
	default:
		armor.Name = "unknown"
	}
	return armor
}

func (a Armor) CalculateArmorBonus(c Character) int {
	var ac int
	dexMod := c.Dexterity.Modifier()
	switch a.WeightCategory {
	case LightArmorCategory:
		ac = a.ArmorClass + dexMod
	case MediumArmorCategory:
		if dexMod > 2 {
			ac = a.ArmorClass + 2
		} else {
			ac = a.ArmorClass + dexMod
		}
	case HeavyArmorCategory:
		ac = a.ArmorClass
	}
	return ac
}
