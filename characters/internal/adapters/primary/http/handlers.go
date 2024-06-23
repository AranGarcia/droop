package http

import (
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h Handler) postCharacter(w http.ResponseWriter, r *http.Request) {
}

func (h Handler) getCharacter(w http.ResponseWriter, r *http.Request) {
}

func (h Handler) listCharacters(w http.ResponseWriter, r *http.Request) {
}

func (h Handler) patchCharacter(w http.ResponseWriter, r *http.Request) {
}

func (h Handler) deleteCharacter(w http.ResponseWriter, r *http.Request) {
}
