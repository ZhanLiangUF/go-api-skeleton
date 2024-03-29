package httputils //import "github.com/ZhanLiangUf/go-api-skeleton/api/httputils

import (
	"encoding/json"
	"net/http"
)

// WriteJSON writes value v to the http response stream as json with standard json encoding
func WriteJSON(w http.ResponseWriter, code int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	return enc.Encode(v)
}
