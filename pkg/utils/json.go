package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseWithJson(w http.ResponseWriter, status int, resp interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	bytes, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	_, err = w.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
