package http

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/AranGarcia/droop/characters/internal/core/entities"
	"github.com/AranGarcia/droop/characters/internal/ports/api/mock"
	"github.com/google/go-cmp/cmp"
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

func TestHandler_getCharacter(t *testing.T) {
	now := time.Now()
	character := entities.Character{
		Base: entities.Base{
			ID:        now.Format(time.RFC3339),
			CreatedAt: now,
			UpdatedAt: now,
		},
		Name:  "Get Character - Name",
		Level: 7,
	}
	tests := []struct {
		name           string
		id             string
		body           string
		wantStatusCode int
		wantBody       map[string]any
	}{
		{
			name:           "get: not found",
			id:             "retrieve: character not found",
			wantStatusCode: http.StatusNotFound,
			wantBody: map[string]any{
				"message": "not found",
			},
		},
		{
			name:           "retrieve success",
			id:             character.ID,
			wantStatusCode: http.StatusOK,
			wantBody: map[string]any{
				"id":            character.ID,
				"created_at":    character.CreatedAt,
				"updated_at":    character.UpdatedAt,
				"name":          character.Name,
				"level":         float64(7),
				"health_points": float64(0),
				"armor_class":   float64(0),
				"strength":      float64(0),
				"dexterity":     float64(0),
				"constitution":  float64(0),
				"intelligence":  float64(0),
				"wisdom":        float64(0),
				"charisma":      float64(0),
			},
		},
	}
	h := Handler{
		characterService: &mock.Characters{
			Entities: map[string]entities.Character{character.ID: character},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.body))
			r.SetPathValue("id", tt.id)
			h.getCharacter(w, r)

			if w.Code != tt.wantStatusCode {
				t.Fatalf("HTTP status code; got %d, want %d", w.Code, tt.wantStatusCode)
			}
			payload, err := io.ReadAll(w.Body)
			if err != nil {
				t.Fatalf("failed to read response body; %v", err)
			}
			var gotBody map[string]any
			if err = json.Unmarshal(payload, &gotBody); err != nil {
				t.Fatalf("failed to unmarshal response body; %v", err)
			}

			if diff := cmp.Diff(gotBody, tt.wantBody); diff != "" {
				t.Fatalf("Response body (+got, -want);\n%s", diff)
			}
		})
	}
}
