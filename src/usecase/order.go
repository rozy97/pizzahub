package usecase

import "github.com/rozy97/pizzahub/src/domain"

type OrderUsecase struct {
	repository domain.OrderRepository
}

func NewOrderUsecase(repository domain.OrderRepository) *OrderUsecase {
	return &OrderUsecase{
		repository: repository,
	}
}

func (usecase *OrderUsecase) AddOrder(order domain.Order) (domain.Order, error) {
	return usecase.repository.InsertOrder(order)
}
