package internal

func (s *Server) setupRoutes() {
	s.mux.HandleFunc("POST /character", s.handler.postCharacter)
	s.mux.HandleFunc("GET /character/{id}", s.handler.getCharacter)
	s.mux.HandleFunc("GET /characters", listCharacters)
	s.mux.HandleFunc("PATCH /character/{id}", patchCharacter)
	s.mux.HandleFunc("DELETE /character/{id}", deleteCharacter)
}
