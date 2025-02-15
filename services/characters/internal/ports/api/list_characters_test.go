package api

import "testing"

func TestListCharactersRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		r       ListCharactersRequest
		wantErr bool
	}{
		{
			name: "invalid sort",
			r: ListCharactersRequest{
				Sort: "something else",
			},
			wantErr: true,
		},
		{
			name: "invalid size",
			r: ListCharactersRequest{
				Sort: CreatedAt,
				Size: -1,
			},
			wantErr: true,
		},
		{
			name: "invalid offset",
			r: ListCharactersRequest{
				Sort:   CreatedAt,
				Size:   0,
				Offset: -1,
			},
			wantErr: true,
		},
		{
			name: "valid request",
			r: ListCharactersRequest{
				Sort:   UpdatedAt,
				Size:   10,
				Offset: 20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ListCharactersRequest.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
