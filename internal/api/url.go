package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ParseURLParameter(r *http.Request, key string) (string, error) {
	val := chi.URLParam(r, key)
	if val == "" {
		return "", fmt.Errorf("`%s` is required in URL", key)
	}

	return val, nil
}
