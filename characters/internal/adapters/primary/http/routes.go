package http

func (s *Server) setupRoutes() {
	s.mux.HandleFunc("POST /character", s.handler.postCharacter)
	s.mux.HandleFunc("GET /character/{id}", s.handler.getCharacter)
	s.mux.HandleFunc("GET /characters", s.handler.listCharacters)
	s.mux.HandleFunc("PATCH /character/{id}", s.handler.patchCharacter)
	s.mux.HandleFunc("DELETE /character/{id}", s.handler.deleteCharacter)
}
