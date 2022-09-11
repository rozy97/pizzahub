package repository

import (
	"time"

	"github.com/rozy97/pizzahub/src/domain"
)

type MenuRepository struct {
}

func NewMenuRepository() *MenuRepository {
	return &MenuRepository{}
}

func (usecase *MenuRepository) GetMenus() ([]domain.Menu, error) {
	menus := []domain.Menu{
		{
			ID:       1,
			Name:     "Pizza Cheese",
			Duration: 3 * time.Second,
		},
		{
			ID:       2,
			Name:     "Pizza BBQ",
			Duration: 5 * time.Second,
		},
	}
	return menus, nil
	// return nil, errors.New("error from repository")
}
