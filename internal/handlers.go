package internal

import (
	"encoding/json"
	"log"
	"net/http"
)

func postCharacter(_ http.ResponseWriter, _ *http.Request) {
	log.Println("Creating character")
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
