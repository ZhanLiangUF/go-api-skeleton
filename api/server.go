package api // import "github.com/ZhanLiangUF/go-flights/api"

import (
	"net"
	"net/http"

	"github.com/ZhanLiangUF/go-flights/api/httputils"
	router "github.com/ZhanLiangUF/go-flights/api/router"
	"github.com/gorilla/mux"
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
	routers []router.Router
	srv     *http.Server
	l       net.Listener
}

// New returns a new instance of the server based on the specified configurations
func New(cfg *Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

// Serve starts listening for inbound requests
func (s *Server) Serve() error {
	return s.srv.Serve(s.l)
}

func (s *Server) makeHTTPHandler(handler httputils.APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// InitRouters initializes the list of routers for the server.
func (s *Server) InitRouters(routers ...router.Router) {
	s.routers = append(s.routers, routers...)
}

// createMux initializes
func (s *Server) createMux() *mux.Router {
	m := mux.NewRouter()
	// logrus.Debug("Registering router")
	for _, apiRouter := range s.routers {
		for _, r := range apiRouter.Routes() {
			f := s.makeHTTPHandler(r.Handler())
		}
	}
}
