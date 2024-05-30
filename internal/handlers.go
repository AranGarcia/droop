package internal

import (
	"log"
	"net/http"
)

func postCharacter(_ http.ResponseWriter, _ *http.Request) {
	log.Println("Creating character")
}

func getCharacter(_ http.ResponseWriter, _ *http.Request) {
	log.Println("Retrieving character")
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
