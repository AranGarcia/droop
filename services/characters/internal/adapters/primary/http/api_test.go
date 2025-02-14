package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/AranGarcia/droop/characters/internal/ports/api"
)

func Test_getStringParam(t *testing.T) {
	tests := []struct {
		name         string
		q            url.Values
		paramName    string
		defaultValue string
		want         string
	}{
		{
			name:         "empty string",
			q:            url.Values{},
			defaultValue: "non empty string",
			want:         "non empty string",
		},
		{
			name:      "param value",
			paramName: "stringParam",
			q:         url.Values{"stringParam": []string{"stringValue"}},
			want:      "stringValue",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStringParam(tt.q, tt.paramName, tt.defaultValue); got != tt.want {
				t.Errorf("getStringParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getIntParam(t *testing.T) {
	tests := []struct {
		name         string
		q            url.Values
		paramName    string
		defaultValue int
		want         int
		wantErr      error
	}{
		{
			name:         "default int value",
			defaultValue: 4,
			want:         4,
		},
		{
			name:      "parse int error",
			q:         url.Values{"nonInteger": []string{"stringValue"}},
			paramName: "nonInteger",
			wantErr:   InvalidInputError{},
		},
		{
			name:      "int value",
			q:         url.Values{"integerParam": []string{"11"}},
			paramName: "integerParam",
			want:      11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getIntParam(tt.q, tt.paramName, tt.defaultValue)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("getIntParam() error = %T, wantErr %T", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getIntParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listCharactersRequestFromHTTP(t *testing.T) {
	tests := []struct {
		name    string
		r       *http.Request
		want    *api.ListCharactersRequest
		wantErr error
	}{
		{
			name: "nil request",
			want: &api.ListCharactersRequest{},
		},
		{
			name:    "invalid sort",
			r:       httptest.NewRequest("", "http://example.com?size=a", nil),
			wantErr: InvalidInputError{},
		},
		{
			name:    "invalid offset",
			r:       httptest.NewRequest("", "http://example.com?offset=a", nil),
			wantErr: InvalidInputError{},
		},
		{
			name: "list characters request",
			r:    httptest.NewRequest("", "http://example.com?size=10&offset=5", nil),
			want: &api.ListCharactersRequest{
				Sort:   api.CreatedAt,
				Size:   10,
				Offset: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := listCharactersRequestFromHTTP(tt.r)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("listCharactersRequestFromHTTP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listCharactersRequestFromHTTP() = %v, want %v", got, tt.want)
			}
		})
	}
}
