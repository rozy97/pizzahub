package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/rozy97/pizzahub/src/domain"
)

type MenuHTTPHandler struct {
	usecase domain.MenuUsecase
}

func NewMenuHTTPHandler(usecase domain.MenuUsecase) *MenuHTTPHandler {
	return &MenuHTTPHandler{
		usecase: usecase,
	}
}

func (handler *MenuHTTPHandler) Menus(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var response Response
	var status int
	if r.Method != "GET" {
		ResponseNotFound(w)
		log.Printf("%v %s %s %v\n", http.StatusNotFound, r.Method, r.URL.Path, time.Since(start))
		return
	}

	result, err := handler.usecase.GetMenus()
	if err != nil {
		response.Message = err.Error()
		status = http.StatusBadRequest
	} else {
		response.Message = "success"
		response.Data = result
		status = http.StatusOK
	}

	responseByte, _ := json.Marshal(&response)
	ResponseJSON(w, status, responseByte)
	log.Printf("%v %s %s %v\n", status, r.Method, r.URL.Path, time.Since(start))
}
