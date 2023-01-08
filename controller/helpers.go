package controller

import (
	"encoding/json"
	"net/http"
)

type respError struct {
	Error string `json:"error"`
}

func ErrorResponse(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	response, _ := json.Marshal(respError{Error: err.Error()})
	w.Write(response)
}
