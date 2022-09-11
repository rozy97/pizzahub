package repository

import (
	"errors"

	"github.com/rozy97/pizzahub/src/domain"
)

type MenuRepository struct {
}

var menus = []domain.Menu{
	{
		ID:       1,
		Name:     "Pizza Cheese",
		Duration: 3,
	},
	{
		ID:       2,
		Name:     "Pizza BBQ",
		Duration: 5,
	},
}

func NewMenuRepository() *MenuRepository {
	return &MenuRepository{}
}

func (usecase *MenuRepository) GetMenus() ([]domain.Menu, error) {

	return menus, nil
}

func (usecase *MenuRepository) GetMenu(ID int) (domain.Menu, error) {
	for _, menu := range menus {
		if menu.ID == ID {
			return menu, nil
		}
	}

	return domain.Menu{}, errors.New("menu not available")
}
