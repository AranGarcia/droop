package internal

import (
	"net/http"
)

type Server struct {
	addr string
	mux  http.ServeMux
}

func NewServer(addr string) *Server {
	s := &Server{
		addr: addr,
	}
	s.setupRoutes()
	return s
}

func (s *Server) Run() error {
	return http.ListenAndServe(s.addr, &s.mux)
}
