package usecase

import (
	"errors"
	"testing"

	"github.com/rozy97/pizzahub/src/domain"
)

func Test_AddOrder(t *testing.T) {
	t.Run("negative test case", func(t *testing.T) {
		orderRepo := OrderRepo{}
		kitchenRepo := KitchenRepo{}
		menuRepo := MenuFailRepo{}
		orderUsecase := NewOrderUsecase(&orderRepo, &kitchenRepo, &menuRepo)
		order := domain.Order{CustomerName: "david", Menu: domain.Menu{ID: 1}}
		order, err := orderUsecase.AddOrder(order)
		if order.ID != 0 {
			t.Errorf("FAIL! expected %v, got %v", 0, order.ID)
		}
		if order.CustomerName != "" {
			t.Errorf("FAIL! expected %v, got %v", "", order.CustomerName)
		}
		if order.Menu.ID != 0 {
			t.Errorf("FAIL! expected %v, got %v", 0, order.Menu.ID)
		}
		if order.Menu.Name != "" {
			t.Errorf("FAIL! expected %v, got %v", 0, order.Menu.Name)
		}
		if order.Chef.ID != 0 {
			t.Errorf("FAIL! expected %v, got %v", 0, order.Chef.ID)
		}
		if order.Chef.Name != "" {
			t.Errorf("FAIL! expected %v, got %v", 0, order.Chef.Name)
		}
		if err == nil {
			t.Errorf("FAIL! expected not %v, got %v", nil, err)
		}
		if err.Error() != "failed get menu" {
			t.Errorf("FAIL! expected %v, got %v", "failed get menu", err.Error())
		}
	})

	t.Run("positive test case", func(t *testing.T) {
		orderRepo := OrderRepo{}
		kitchenRepo := KitchenRepo{}
		menuRepo := MenuRepo{}
		orderUsecase := NewOrderUsecase(&orderRepo, &kitchenRepo, &menuRepo)
		order := domain.Order{CustomerName: "david", Menu: domain.Menu{ID: 1}}
		order, err := orderUsecase.AddOrder(order)
		if order.ID != 12 {
			t.Errorf("FAIL! expected %v, got %v", 0, order.ID)
		}
		if order.CustomerName != "david" {
			t.Errorf("FAIL! expected %v, got %v", "", order.CustomerName)
		}
		if order.Menu.ID != 1 {
			t.Errorf("FAIL! expected %v, got %v", 0, order.Menu.ID)
		}
		if order.Menu.Name != "pizza" {
			t.Errorf("FAIL! expected %v, got %v", 0, order.Menu.Name)
		}
		if order.Chef.ID != 1 {
			t.Errorf("FAIL! expected %v, got %v", 0, order.Chef.ID)
		}
		if order.Chef.Name != "juna" {
			t.Errorf("FAIL! expected %v, got %v", 0, order.Chef.Name)
		}
		if err != nil {
			t.Errorf("FAIL! expected %v, got %v", nil, err)
		}
	})
}

type OrderRepo struct {
}

type MenuFailRepo struct {
}

func (repository *MenuFailRepo) GetMenus() ([]domain.Menu, error) {
	return []domain.Menu{}, errors.New("failed get menus")
}

func (repository *MenuFailRepo) GetMenu(ID int) (domain.Menu, error) {
	return domain.Menu{}, errors.New("failed get menu")
}

func (repository *OrderRepo) InsertOrder(order domain.Order) (domain.Order, error) {
	order.ID = 12
	return order, nil
}
