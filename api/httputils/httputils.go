package httputils // import "github.com/ZhanLiangUf/go-flights/api/httputil

import (
	"context"
	"net/http"
)

// APIFunc is an adapter to allow the use of ordinary functions as Docker API endpoints.
// Any function that has the appropriate signature can be registered as an API endpoint (e.g. getVersion).
type APIFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error
