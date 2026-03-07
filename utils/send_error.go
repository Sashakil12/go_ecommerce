package utils

import (
	"encoding/json"
	"net/http"
)

func SendError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(message)
}
