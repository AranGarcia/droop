package characters

import (
	"errors"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_translateGRPCError(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		wantErr error
	}{
		{
			name: "nil input",
		},
		{
			name:    "not a status",
			err:     errors.New(""),
			wantErr: ErrUnexpectedError,
		},
		{
			name:    "not found",
			err:     status.Error(codes.NotFound, ""),
			wantErr: ErrNotFound,
		},
		{
			name:    "default status",
			err:     status.Error(codes.Aborted, ""),
			wantErr: ErrUnexpectedError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := translateGRPCError(tt.err)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("translateGRPCError() error = %T, wantErr %T", err, tt.wantErr)
			}
		})
	}
}
