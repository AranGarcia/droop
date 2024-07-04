package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/AranGarcia/droop/characters/internal/ports/api"
)

type Handler struct {
	characterService api.Characters
}

func NewHandler(api api.Characters) *Handler {
	return &Handler{
		characterService: api,
	}
}

func (h Handler) postCharacter(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { handleAPIError(w, err) }()

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	apiRequest := api.CreateCharacterRequest{}
	if err = json.Unmarshal(payload, &apiRequest); err != nil {
		return
	}

	apiResponse, err := h.characterService.Create(r.Context(), apiRequest)
	if err != nil {
		return
	}

	body, err := json.Marshal(apiResponse)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

func (h Handler) getCharacter(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { handleAPIError(w, err) }()

	id := r.PathValue("id")
	apiRequest := api.RetrieveCharacterRequest{ID: id}
	apiResponse, err := h.characterService.Retrieve(r.Context(), apiRequest)
	if err != nil {
		return
	}
	body, err := json.Marshal(apiResponse)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

func (h Handler) listCharacters(w http.ResponseWriter, r *http.Request) {
}

func (h Handler) patchCharacter(w http.ResponseWriter, r *http.Request) {
}

func (h Handler) deleteCharacter(w http.ResponseWriter, r *http.Request) {
}
