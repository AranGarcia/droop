package internal

import (
	"log"
	"net/http"
)

type Server struct {
	addr string
	mux  *http.ServeMux
}

func NewServer(addr string) *Server {
	s := &Server{
		addr: addr,
		mux:  http.NewServeMux(),
	}
	s.setupRoutes()
	return s
}

func (s *Server) Run() error {
	handler := loggingMiddleware(jsonMiddlewhare(s.mux))
	return http.ListenAndServe(s.addr, handler)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL)

		next.ServeHTTP(w, r)
	})
}

func jsonMiddlewhare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
