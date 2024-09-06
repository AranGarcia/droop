package api

import (
	"net/url"
	"reflect"
	"testing"
)

func TestListCharactersRequestFromURLQuery(t *testing.T) {
	tests := []struct {
		name    string
		q       url.Values
		want    *ListCharactersRequest
		wantErr bool
	}{
		{
			name: "success",
			q: url.Values{
				"size":   []string{"5"},
				"offset": []string{"10"},
			},
			want: &ListCharactersRequest{
				Size:   5,
				Offset: 10,
			},
		},
		{
			name: "size non-int",
			q: url.Values{
				"size": []string{"size string"},
			},
			wantErr: true,
		},
		{
			name: "negative size",
			q: url.Values{
				"size": []string{"-1"},
			},
			wantErr: true,
		},
		{
			name: "offset non-int",
			q: url.Values{
				"offset": []string{"offset string"},
			},
			wantErr: true,
		},
		{
			name: "negative offset",
			q: url.Values{
				"offset": []string{"-1"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListCharactersRequestFromURLQuery(tt.q)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListCharactersRequestFromURLQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListCharactersRequestFromURLQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
