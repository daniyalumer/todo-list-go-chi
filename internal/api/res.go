package api

import (
	"encoding/json"
	"net/http"
)

func ParseResponse(w http.ResponseWriter, payload interface{}, status uint) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(status))

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
