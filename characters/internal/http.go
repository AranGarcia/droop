package internal

import (
	"log"
	"net/http"
)

// Server is an HTTP server.
type Server struct {
	addr    string
	mux     *http.ServeMux
	handler *Handler
}

// NewServer initializes a server with the URI routes added with their handlers.s
func NewServer(addr string, handler *Handler) *Server {
	s := &Server{
		addr:    addr,
		mux:     http.NewServeMux(),
		handler: handler,
	}
	s.setupRoutes()
	return s
}

// Run the server.
func (s *Server) Run() error {
	handler := loggingMiddleware(jsonMiddleware(s.mux))
	return http.ListenAndServe(s.addr, handler)
}

// loggingMiddleware logs all server requests.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL)

		next.ServeHTTP(w, r)
	})
}

// jsonMiddleware manages all requests with the header Content-Type set as application/json.
func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
