package test

import (
	"context"
	"fmt"
	"net/http"
)

func (s *testRouter) getFlightsGet(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	fmt.Fprintf(w, vars["version"])
	return nil
}
