package http

import (
	"log"
	"net/http"

	"github.com/AranGarcia/droop/characters/internal/core/api"
)

// Server is an HTTP server.
type Server struct {
	// addr is the address where the server will be running on.
	addr string
	// accessControlAllowOrigin is used with CORS policies.
	accessControlAllowOrigin string

	mux     *http.ServeMux
	handler *Handler
}

// NewServer initializes a server with the URI routes added with their handlers.s
func NewServer(addr, accessControlAllowOrigin string, api api.Characters) *Server {
	s := &Server{
		addr:                     addr,
		accessControlAllowOrigin: accessControlAllowOrigin,
		mux:                      http.NewServeMux(),
		handler:                  NewHandler(api),
	}
	s.setupRoutes()
	return s
}

// Run the server.
func (s *Server) Run() error {
	handler := loggingMiddleware(jsonMiddleware(s.corsMiddleware(s.mux)))
	return http.ListenAndServe(s.addr, handler)
}

func (s Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", s.accessControlAllowOrigin)
		next.ServeHTTP(w, r)
	})
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
