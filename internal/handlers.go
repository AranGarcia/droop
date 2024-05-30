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
