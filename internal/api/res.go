package api

import (
	"encoding/json"
	"net/http"
)

func ResponseWriter(w http.ResponseWriter, payload interface{}, statusCode uint) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(statusCode))

	if payload == nil {
		return nil
	}

	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(err)
		return err
	}

	return nil
}
