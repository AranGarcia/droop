package internal

func (s *Server) setupRoutes() {
	s.mux.HandleFunc("POST /character", postCharacter)
	s.mux.HandleFunc("GET /character/{id}", getCharacter)
	s.mux.HandleFunc("GET /characters", listCharacters)
	s.mux.HandleFunc("PATCH /character/{id}", patchCharacter)
	s.mux.HandleFunc("DELETE /character/{id}", deleteCharacter)
}
