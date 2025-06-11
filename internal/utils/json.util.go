package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, status int, msg string) error {
	resp := ErrorResponse{
		Status: status,
		Error:  msg,
	}
	return WriteJSON(w, status, resp)
}

func ReadJSON(r *http.Request, data any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}
