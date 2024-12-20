package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseRequest(r *http.Request, req interface{}) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return fmt.Errorf("unable to parse request: %v", err)
	}

	return nil
}
