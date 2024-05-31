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

func (h Handler) postCharacter(w http.ResponseWriter, r *http.Request) {
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("failed to read request body bytes")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	character := Character{}
	if err = json.Unmarshal(payload, &character); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to unmarshal request payload; %v", err)))
		return
	}

	output, err := h.repo.CreateCharacter(r.Context(), character)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	responseBody, err := json.Marshal(output)
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
