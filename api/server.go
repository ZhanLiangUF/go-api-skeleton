package api

import (
	"net/http"

	"github.com/ZhanLiangUF/go-flights/api/httputils"
)

// Config provides configuration for server
type Config struct {
	Logging    bool
	CorsHeader string
	Version    string
}

// Server contains instance details of the server
type Server struct {
	cfg     *Config
	routers []Router
}

// New returns a new instance of the server based on the specified configurations
func New(cfg *Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) makeHTTPHandler(handler httputils.APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// InitRouters initializes the list of routers for the server.
func (s *Server) InitRouters(routers ...Router) {
	s.routers = append(s.routers, routers...)
}

// func (s *Server) createMux() *mux.Router {

// }
