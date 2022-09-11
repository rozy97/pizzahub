package usecase

import "github.com/rozy97/pizzahub/src/domain"

type OrderUsecase struct {
	repository  domain.OrderRepository
	kitchenRepo domain.KitchenRepository
}

func NewOrderUsecase(repository domain.OrderRepository, kitchenRepo domain.KitchenRepository) *OrderUsecase {
	return &OrderUsecase{
		repository:  repository,
		kitchenRepo: kitchenRepo,
	}
}

func (usecase *OrderUsecase) AddOrder(order domain.Order) (domain.Order, error) {
	return usecase.repository.InsertOrder(order)
}
