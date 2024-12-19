package api

import (
	"fmt"
	"net/http"
)

func ParseForm(w http.ResponseWriter, r *http.Request) bool {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return false
	}
	return true
}

func ParseURLParameter(r *http.Request, key string) (string, error) {
	val := r.URL.Query().Get(key)
	if val == "" {
		return "", fmt.Errorf("missing parameter %s", key)
	}

	return val, nil
}
