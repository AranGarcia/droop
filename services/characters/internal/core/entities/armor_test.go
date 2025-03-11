package entities

import "testing"

func TestArmor_CalculateArmorBonus(t *testing.T) {
	c := Character{Dexterity: AbilityScore(15)}
	tests := []struct {
		name      string
		armor     Armor
		character Character
		want      int
	}{
		{
			name:      "light armor",
			armor:     *NewArmor(LeatherArmor),
			character: c,
			want:      11 + c.Dexterity.Modifier(),
		},
		{
			name:      "medium armor with dex mod <= 2",
			armor:     *NewArmor(ScaleMailArmor),
			character: c,
			want:      14 + c.Dexterity.Modifier(),
		},
		{
			name:      "medium armor with dex mod > 2",
			armor:     *NewArmor(HalfPlateArmor),
			character: Character{Dexterity: AbilityScore(20)}, // Dex mod of +5
			want:      17,                                     // Half Plate AC + 2
		},
		{
			name:      "heavy armor",
			armor:     *NewArmor(PlateArmor),
			character: c,
			want:      18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.armor.CalculateArmorBonus(tt.character); got != tt.want {
				t.Errorf("Armor.CalculateArmorBonus() = %v, want %v", got, tt.want)
			}
		})
	}
}
