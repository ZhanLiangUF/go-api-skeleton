package middleware // import "github.com/zhanlianguf/go-flights/api/middleware"
import (
	"context"
	"net/http"
)

// CORSMiddleware inject CORS headers to each request when it's configured
type CORSMiddleware struct {
	defaultHeaders string
}

// NewCorsMiddleware creates a new CORSMiddleware with default headers
func NewCORSMiddleware(d string) CORSMiddleware {
	return CORSMiddleware{defaultHeaders: d}
}

// WrapHandler returns a new handler function wrapping the previous one in the request chain
func (c CORSMiddleware) WrapHandler(handler func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error) func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
		corsHeader := c.defaultHeaders
		if corsHeader == "" {
			corsHeader = "*"
		}

		w.Header().Add("Access-Control-Allow-Origin", corsHeader)
		w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, X-Registry-Auth")
		w.Header().Add("Access-Control-Allow-Methods", "HEAD, GET, POST, DELETE, PUT, OPTIONS")
		return handler(ctx, w, r, vars)
	}
}
