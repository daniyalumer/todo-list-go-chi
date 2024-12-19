package helper

import (
	"encoding/json"
	"net/http"
	"strconv"
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

func ParseForm(w http.ResponseWriter, r *http.Request) bool {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return false
	}
	return true
}

func ConvertIdToInteger(IdStr string) (int, error) {
	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		return 0, err
	}
	return Id, nil
}
