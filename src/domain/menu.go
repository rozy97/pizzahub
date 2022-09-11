package domain

import "time"

type Menu struct {
	ID       int           `json:"id"`
	Name     string        `json:"name"`
	Duration time.Duration `json:"-"`
}

type MenuUsecase interface {
	GetMenus() ([]Menu, error)
}

type MenuRepository interface {
	GetMenus() ([]Menu, error)
	// GetMenu(ID int) (Menu, error)
	// InsertMenu(menu Menu) (Menu, error)
}
