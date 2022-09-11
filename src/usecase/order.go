package usecase

import (
	"fmt"
	"time"

	"github.com/rozy97/pizzahub/src/domain"
)

type OrderUsecase struct {
	repository  domain.OrderRepository
	kitchenRepo domain.KitchenRepository
	menuRepo    domain.MenuRepository
}

func NewOrderUsecase(repository domain.OrderRepository, kitchenRepo domain.KitchenRepository, menuRepo domain.MenuRepository) *OrderUsecase {
	return &OrderUsecase{
		repository:  repository,
		kitchenRepo: kitchenRepo,
		menuRepo:    menuRepo,
	}
}

func (usecase *OrderUsecase) AddOrder(order domain.Order) (domain.Order, error) {
	// get menu from database
	menu, err := usecase.menuRepo.GetMenu(order.Menu.ID)
	if err != nil {
		return domain.Order{}, err
	}

	// get chef from queue
	chef := usecase.kitchenRepo.GetChefFromKitchen()

	// process cooking
	usecase.processCook(menu)
	order.Chef.ID = chef.ID
	order.Chef.Name = chef.Name
	order.Menu.Name = menu.Name

	// give back chef to queue
	usecase.kitchenRepo.AddChefToKitchen(chef, fmt.Sprintf("back after cook %s", order.Menu.Name))

	return usecase.repository.InsertOrder(order)
}

func (usecase *OrderUsecase) processCook(menu domain.Menu) {
	time.Sleep(time.Duration(menu.Duration) * time.Second)
}
