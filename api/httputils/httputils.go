package httputils // import "github.com/ZhanLiangUf/go-flights/api/httputil

import (
	"context"
	"net/http"
)

// APIFunc is an adapter to allow the use of ordinary functions as API endpoints.
type APIFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error
