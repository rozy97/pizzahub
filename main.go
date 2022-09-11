package main

import (
	"log"
	"net/http"

	"github.com/rozy97/pizzahub/src/handler"
	"github.com/rozy97/pizzahub/src/repository"
	"github.com/rozy97/pizzahub/src/usecase"
)

func main() {

	chefRepository := repository.NewChefRepository()
	chefUsecase := usecase.NewChefUsecase(chefRepository)
	chefHandler := handler.NewChefHTTPHandler(chefUsecase)

	menuRepository := repository.NewMenuRepository()
	menuUsecase := usecase.NewMenuUsecase(menuRepository)
	menuHandler := handler.NewMenuHTTPHandler(menuUsecase)

	orderRepository := repository.NewOrderRepository()
	orderUsecase := usecase.NewOrderUsecase(orderRepository)
	orderHandler := handler.NewOrderHTTPHandler(orderUsecase)

	http.HandleFunc("/chef", chefHandler.Chef)
	http.HandleFunc("/menus", menuHandler.Menus)
	http.HandleFunc("/orders", orderHandler.Order)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Println(err.Error())
	}
}
