package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/rozy97/pizzahub/src/domain"
)

type OrderHTTPHandler struct {
	usecase domain.OrderUsecase
}

func NewOrderHTTPHandler(usecase domain.OrderUsecase) *OrderHTTPHandler {
	return &OrderHTTPHandler{
		usecase: usecase,
	}
}

func (handler *OrderHTTPHandler) Order(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	if r.Method != "POST" {
		ResponseNotFound(w)
		log.Printf("%v %s %s %v\n", http.StatusNotFound, r.Method, r.URL.Path, time.Since(start))
		return
	}

	payloadByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response := &Response{Message: err.Error()}
		responseByte, _ := json.Marshal(response)
		ResponseJSON(w, http.StatusBadRequest, responseByte)
		log.Printf("%v %s %s %v\n", http.StatusBadRequest, r.Method, r.URL.Path, time.Since(start))
		return
	}

	var order domain.Order
	err = json.Unmarshal(payloadByte, &order)
	if err != nil {
		response := &Response{Message: err.Error()}
		responseByte, _ := json.Marshal(response)
		ResponseJSON(w, http.StatusBadRequest, responseByte)
		log.Printf("%v %s %s %v\n", http.StatusBadRequest, r.Method, r.URL.Path, time.Since(start))
		return
	}

	result, err := handler.usecase.AddOrder(order)
	if err != nil {
		response := &Response{Message: err.Error()}
		responseByte, _ := json.Marshal(response)
		ResponseJSON(w, http.StatusBadRequest, responseByte)
		log.Printf("%v %s %s %v\n", http.StatusBadRequest, r.Method, r.URL.Path, time.Since(start))
		return
	}

	response := &Response{Message: "success", Data: result}
	responseByte, _ := json.Marshal(response)
	ResponseJSON(w, http.StatusOK, responseByte)
	log.Printf("%v %s %s %v\n", http.StatusOK, r.Method, r.URL.Path, time.Since(start))
}
