package httputils // import "github.com/ZhanLiangUf/go-flights/api/httputil

import (
	"context"
	"net/http"

	"github.com/ZhanLiangUF/go-flights/errdefs"
)

// APIFunc is an adapter to allow the use of ordinary functions as API endpoints.
type APIFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error

// MakeErrorHandler creates HTTP handler that decodes an error and returns it in the response
func MakeErrorHandler(err error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the status code and access get version variable from path to see if it's a legit API version
		statusCode := errdefs.GetHTTPErrorStatusCode(err)

	}
}
