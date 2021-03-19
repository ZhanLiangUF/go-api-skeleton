package api // import "github.com/ZhanLiangUF/go-flights/api"

import (
	"net"
	"net/http"

	"github.com/ZhanLiangUF/go-flights/api/httputils"
	"github.com/ZhanLiangUF/go-flights/api/middleware"
	router "github.com/ZhanLiangUF/go-flights/api/router"
	"github.com/gorilla/mux"
)

// can access version using mux.vars
const versionMatcher = "/v{version:[0-9.]+}"

// Config provides configuration for server
type Config struct {
	Logging    bool
	CorsHeader string
	Version    string
}

// Server contains instance details of the server
type Server struct {
	cfg        *Config
	routers    []router.Router
	srv        *HTTPServer
	middleware middleware.Middleware
}

// New returns a new instance of the server based on the specified configurations
func New(cfg *Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

// UseMiddleware sets a new middleware to the request chain.
// Needs to be called before API are configured
func (s *Server) UseMiddleware(m middleware.Middleware) {
	s.middleware = m
}

// Accept sets listener the server accepts connection into
func (s *Server) Accept(addr string, listener net.Listener) {
	s.srv = &HTTPServer{
		srv: &http.Server{
			Addr: addr,
		},
		l: listener,
	}
}

// Close closes server and stops receiving requests
func (s *Server) Close() {
	if err := s.srv.Close(); err != nil {
		// log here?
	}
}

// ServeAPI initialize servers
func (s *Server) ServeAPI() error {
	s.srv.srv.Handler = s.createMux()
	err := s.srv.Serve()
	return err
}

// HTTPServer contains an instance of http server and listener
type HTTPServer struct {
	srv *http.Server
	l   net.Listener
}

// Serve starts listening for inbound requests
func (s *HTTPServer) Serve() error {
	return s.srv.Serve(s.l)
}

// Close immediately closes all active net.Listeners and any connections in state StateNew, StateActive, or StateIdle. For a graceful shutdown, use Shutdown.
// Close does not attempt to close (and does not even know about) any hijacked connections, such as WebSockets.
// Close returns any error returned from closing the Server's underlying Listener(s).
func (s *HTTPServer) Close() error {
	return s.l.Close()
}

func (s *Server) makeHTTPHandler(handler httputils.APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handlerFunc := s.handlerWithGlobalMiddleWare(handler)
		vars := mux.Vars(r)
		if vars == nil {
			vars = make(map[string]string)
		}
		// have to take care of error here
		if err := handlerFunc(ctx, w, r, vars), err != nil {
			statusCode := errdefs.GetHTTPErrorStatusCode(err)
			if statusCode >= 500 {
				// log 
			}
		}

	}
}

// InitRouters initializes the list of routers for the server.
func (s *Server) InitRouters(routers ...router.Router) {
	s.routers = append(s.routers, routers...)
}

// createMux initializes the main router
func (s *Server) createMux() *mux.Router {
	m := mux.NewRouter()
	// log here?
	for _, apiRouter := range s.routers {
		for _, r := range apiRouter.Routes() {
			f := s.makeHTTPHandler(r.Handler())
			m.Path(versionMatcher + r.Path()).Methods(r.Method()).Handler(f)
			m.Path(r.Path()).Methods(r.Method()).Handler(f)
		}
	}
	notFoundHandler := httputils.MakeErrorHandler(pageNotFoundError{})
	// error handler when users try to type a legit API
	m.HandleFunc(versionMatcher+"/{path:.*}", notFoundHandler)
	// error handler for when reach a part of the site that user shouldnt go?
	m.NotFoundHandler = notFoundHandler
	m.MethodNotAllowedHandler = notFoundHandler
	return m
}
