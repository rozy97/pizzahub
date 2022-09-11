package handler

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseNotFound(w http.ResponseWriter) {
	response := &Response{Message: http.StatusText(http.StatusNotFound)}
	responseByte, _ := json.Marshal(response)
	ResponseJSON(w, http.StatusNotFound, responseByte)
}

func ResponseJSON(w http.ResponseWriter, statusCode int, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(body)
}
