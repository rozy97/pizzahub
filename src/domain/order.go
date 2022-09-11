package domain

type Order struct {
	ID           int    `json:"id"`
	CustomerName string `json:"customer_name"`
	Menu         Menu   `json:"menu,omitempty"`
	Chef         Chef   `json:"chef,omitempty"`
}

type OrderUsecase interface {
	AddOrder(order Order) (Order, error)
	// GetOrders() ([]Order, error)
	// GetOrder(ID int) (Order, error)
}

type OrderRepository interface {
	InsertOrder(order Order) (Order, error)
}
