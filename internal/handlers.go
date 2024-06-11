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
		jsonErr := JSONErrorResponse{Message: fmt.Sprintf("invalid request data; %v", err)}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonErr.ToBytes())
		return
	}

	if err := character.Validate(); err != nil {
		jsonErr := JSONErrorResponse{Message: fmt.Sprintf("invalid request data; %v", err)}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonErr.ToBytes())
		return
	}

	output, err := h.repo.CreateCharacter(r.Context(), character)
	if err != nil {
		jsonErr := JSONErrorResponse{Message: fmt.Sprintf("failed to create character; %v", err)}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonErr.ToBytes())
		return
	}
	responseBody, err := json.Marshal(output)
	if err != nil {
		log.Println("failed to marshal data for character with ID", output.ID)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)

}

// getCharacter is a GET endpoint for the Character resource.
func getCharacter(w http.ResponseWriter, r *http.Request) {
	character := Character{}
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

func deleteCharacter(_ http.ResponseWriter, _ *http.Request) {
	log.Println("deleting character")
}
