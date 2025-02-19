package grpc

import (
	"errors"
	"reflect"
	"testing"

	"github.com/AranGarcia/droop/characters/internal/core/api"
	"google.golang.org/grpc/codes"
)

func Test_apiErrToCode(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want codes.Code
	}{
		{
			name: "invalid request",
			err:  api.InvalidRequestError{},
			want: codes.InvalidArgument,
		},
		{
			name: "not found",
			err:  api.ErrNotFound,
			want: codes.NotFound,
		},
		{
			name: "internal",
			err:  errors.New(""),
			want: codes.Internal,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := apiErrToCode(tt.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("apiErrToCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
