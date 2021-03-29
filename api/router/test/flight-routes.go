package test

import (
	"context"
	"net/http"
)

func (s *testRouter) getFlightsGet(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {

	w.Header().Set("Content-Type", "application/x-tar")
	return nil
}
