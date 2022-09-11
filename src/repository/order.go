package repository

import "github.com/rozy97/pizzahub/src/domain"

type OrderRepository struct {
}

var orders []domain.Order

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

func (repository *OrderRepository) InsertOrder(order domain.Order) (domain.Order, error) {
	id := len(orders) + 1
	order.ID = id
	orders = append(orders, order)
	return order, nil
}
