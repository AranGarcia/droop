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

func (h Handler) getCharacter(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	character, err := h.repo.Retrieve(r.Context(), id)
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

func (h Handler) listCharacters(w http.ResponseWriter, r *http.Request) {
	type ListResponse struct {
		Length     int          `json:"length"`
		Characters []*Character `json:"characters"`
	}

	var err error
	response := ListResponse{}
	response.Characters, err = h.repo.List(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(JSONErrorResponse{Error: err.Error()}.ToBytes())
	}
	response.Length = len(response.Characters)

	var payload []byte
	if payload, err = json.Marshal(response); err != nil {
		w.Write(JSONErrorResponse{Error: fmt.Sprintf("failed to marshal response; %v", err)}.ToBytes())
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}

func (h Handler) patchCharacter(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write(JSONErrorResponse{Error: "failed to read JSON data"}.ToBytes())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updateData := &UpdateFields{}
	if err = json.Unmarshal(payload, updateData); err != nil {
		jsonErr := JSONErrorResponse{Error: fmt.Sprintf("invalid request data; %v", err)}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonErr.ToBytes())
		return
	}

	if err = h.repo.Update(r.Context(), id, updateData); err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(JSONErrorResponse{Error: fmt.Sprintf("repository error; %v", err)}.ToBytes())
		return
	}

	character, err := h.repo.Retrieve(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(JSONErrorResponse{Error: fmt.Sprintf("failed to retrieve updated character; %v", err)}.ToBytes())
		return
	}

	responseBody, err := json.Marshal(character)
	if err != nil {
		log.Println("failed to marshal data for character with ID", character.ID)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

func (h Handler) deleteCharacter(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.repo.DeleteCharacter(r.Context(), id); err != nil {
		w.Write(JSONErrorResponse{Error: err.Error()}.ToBytes())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
