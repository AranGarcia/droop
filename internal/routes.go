package internal

func (s *Server) setupRoutes() {
	/*
		`POST /character`
		`GET /character/:id`
		`GET /characters`
		`PATCH /character/:id`
		`DELETE /character/:id`
	*/
	s.mux.HandleFunc("POST /character", postCharacter)
	s.mux.HandleFunc("GET /character/{id}", getCharacter)
}
