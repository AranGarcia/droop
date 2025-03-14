package services

// func TestCharacters_GenerateSheet(t *testing.T) {
// 	tests := []struct {
// 		name       string
// 		repository repositories.Characters
// 		request    api.GenerateSheetRequest
// 		want       *api.GenerateSheetResponse
// 		wantErr    bool
// 	}{
// 		{
// 			name:       "not found",
// 			request:    api.GenerateSheetRequest{ID: "not-found"},
// 			repository: &mocks.Characters{},
// 			wantErr:    true,
// 		},
// 		{
// 			name: "calculated success",
// 			request: api.GenerateSheetRequest{
// 				ID: "character-id",
// 			},
// 			repository: &mocks.Characters{
// 				InMemory: map[string]entities.Character{
// 					"character-id": {
// 						Level:     entities.Level(6),
// 						Armor:     entities.NewArmor(entities.ScaleMailArmor), // 14 + DEX mod
// 						Dexterity: entities.AbilityScore(15),
// 					},
// 				},
// 			},
// 			want: &api.GenerateSheetResponse{
// 				Characer: entities.Character{
// 					Level:     entities.Level(6),
// 					Armor:     entities.NewArmor(entities.ScaleMailArmor),
// 					Dexterity: entities.AbilityScore(15),
// 				},
// 				ArmorClass:       16,
// 				ProficiencyBonus: 3,
// 				InitiativeBonus:  2,
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			c := Characters{
// 				repository: tt.repository,
// 			}
// 			ctx := context.Background()
// 			got, err := c.GenerateSheet(ctx, tt.request)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Characters.GenerateSheet() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if diff := cmp.Diff(tt.want, got); diff != "" {
// 				t.Fatalf("Characters.GenerateSheet() mismatch (-want,+got):\n%s", diff)
// 			}
// 		})
// 	}
// }
