package usecase

import (
	"testing"

	"github.com/rozy97/pizzahub/src/domain"
)

func Test_GetMenus(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		menuRepo := MenuRepo{}
		menuUsecase := NewMenuUsecase(&menuRepo)

		menus, err := menuUsecase.GetMenus()
		for _, menu := range menus {

			if menu.ID != 1 {
				t.Errorf("FAIL! expected %v, got %v", 1, menu.ID)
			}
			if menu.Name != "pizza" {
				t.Errorf("FAIL! expected %v, got %v", "pizza", menu.Name)
			}
			if err != nil {
				t.Errorf("FAIL! expected %v, got %v", nil, err)
			}
		}
	})
}

type MenuRepo struct{}

func (repository *MenuRepo) GetMenus() ([]domain.Menu, error) {
	return []domain.Menu{
		{
			ID:       1,
			Name:     "pizza",
			Duration: 1,
		},
	}, nil
}

func (repository *MenuRepo) GetMenu(ID int) (domain.Menu, error) {
	return domain.Menu{ID: ID, Name: "pizza"}, nil
}
