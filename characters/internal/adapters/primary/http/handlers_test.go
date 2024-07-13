package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/AranGarcia/droop/characters/internal/ports/api/mock"
)

func TestHandler_postCharacter(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		wantStatusCode int
	}{
		{
			name:           "invalid post body",
			body:           `}{`,
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:           "create success",
			body:           `{}`,
			wantStatusCode: http.StatusCreated,
		},
	}
	h := Handler{
		characterService: mock.NewMockCharacterService(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.body))
			h.postCharacter(w, r)
			if w.Code != tt.wantStatusCode {
				t.Fatalf("HTTP status code; got %d, want %d", w.Code, tt.wantStatusCode)
			}
		})
	}
}
