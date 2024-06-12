package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Handler struct {
	repo *CharacterRepository
}

func NewHandler(repository *CharacterRepository) *Handler {
	return &Handler{
		repo: repository,
	}
}

func (h Handler) postCharacter(w http.ResponseWriter, r *http.Request) {
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("failed to read request body bytes")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	character := Character{}
	if err = json.Unmarshal(payload, &character); err != nil {
		jsonErr := JSONErrorResponse{Error: fmt.Sprintf("invalid request data; %v", err)}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonErr.ToBytes())
		return
	}

	if err := character.Validate(); err != nil {
		jsonErr := JSONErrorResponse{Error: fmt.Sprintf("invalid request data; %v", err)}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonErr.ToBytes())
		return
	}

	output, err := h.repo.CreateCharacter(r.Context(), character)
	if err != nil {
		jsonErr := JSONErrorResponse{Error: fmt.Sprintf("failed to create character; %v", err)}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonErr.ToBytes())
		return
	}
	responseBody, err := json.Marshal(output)
	if err != nil {
		log.Println("failed to marshal data for character with ID", output.ID)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(responseBody)

}

// getCharacter is a GET endpoint for the Character resource.
func (h Handler) getCharacter(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	character, err := h.repo.RetrieveCharacter(r.Context(), id)
	if err != nil {
		jsonErr := JSONErrorResponse{Error: fmt.Sprintf("failed to retrieve character; %v", err)}
		w.WriteHeader(http.StatusNotFound)
		w.Write((jsonErr.ToBytes()))
		return
	}
	responseBody, err := json.Marshal(character)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

func listCharacters(_ http.ResponseWriter, _ *http.Request) {
	log.Println("Listing characters")
}

func patchCharacter(_ http.ResponseWriter, _ *http.Request) {
	log.Println("Updating character")
}

func (h Handler) deleteCharacter(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.repo.deleteCharacter(r.Context(), id); err != nil {
		w.Write(JSONErrorResponse{Error: err.Error()}.ToBytes())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
