package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type ChefHTTPHandler struct {
}

func NewChefHTTPHandler() *ChefHTTPHandler {
	return &ChefHTTPHandler{}
}

func (handler *ChefHTTPHandler) Chef(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	if r.Method != "POST" {
		ResponseNotFound(w)
		log.Printf("%v %s %s %v\n", http.StatusNotFound, r.Method, r.URL.Path, time.Since(start))
		return
	}

	response := &Response{Message: "success"}
	responseByte, _ := json.Marshal(response)
	ResponseJSON(w, http.StatusOK, responseByte)
	log.Printf("%v %s %s %v\n", http.StatusOK, r.Method, r.URL.Path, time.Since(start))
}
