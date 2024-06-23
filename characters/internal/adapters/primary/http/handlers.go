package http

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/AranGarcia/droop/characters/internal/core/entities"
	"github.com/AranGarcia/droop/characters/internal/ports/api"
)

type Handler struct {
	characterService api.Characters
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h Handler) postCharacter(w http.ResponseWriter, r *http.Request) {
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("failed to read request body bytes")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	character := entities.Character{}
	if err = json.Unmarshal(payload, &character); err != nil {
		jsonErr := JSONErrorResponse{Error: fmt.Sprintf("invalid request data; %v", err)}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonErr.ToBytes())
		return
	}

	apiRequest := api.CreateCharacterRequest{Character: character}
	_, err = h.characterService.Create(r.Context(), apiRequest)
	if err != nil {
		jsonErr := JSONErrorResponse{Error: fmt.Sprintf("failed to create character; %v", err)}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonErr.ToBytes())
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{}`))
}

func (h Handler) getCharacter(w http.ResponseWriter, r *http.Request) {
}

func (h Handler) listCharacters(w http.ResponseWriter, r *http.Request) {
}

func (h Handler) patchCharacter(w http.ResponseWriter, r *http.Request) {
}

func (h Handler) deleteCharacter(w http.ResponseWriter, r *http.Request) {
}
