package main

import (
	"log"
	"net/http"

	"github.com/rozy97/pizzahub/src/handler"
)

func main() {

	chefHandler := handler.NewChefHTTPHandler()
	menuHandler := handler.NewMenuHTTPHandler()
	orderHandler := handler.NewOrderHTTPHandler()

	http.HandleFunc("/chef", chefHandler.Chef)
	http.HandleFunc("/menus", menuHandler.Menus)
	http.HandleFunc("/orders", orderHandler.Order)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Println(err.Error())
	}
}
