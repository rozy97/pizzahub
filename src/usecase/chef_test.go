package usecase

import (
	"errors"
	"testing"

	"github.com/rozy97/pizzahub/src/domain"
)

func Test_AddChef(t *testing.T) {
	t.Run("negative test case", func(t *testing.T) {
		chefRepo := ChefRepo{}
		kitchenRepo := KitchenRepo{}
		chefUsecase := NewChefUsecase(&chefRepo, &kitchenRepo)
		chef := domain.Chef{ID: -1, Name: "andi"}

		chef, err := chefUsecase.AddChef(chef)
		if chef.ID != 0 {
			t.Errorf("FAIL! expected %v, got %v", 0, chef.ID)
		}
		if chef.Name != "" {
			t.Errorf("FAIL! expected %v, got %v", "", chef.Name)
		}
		if err == nil {
			t.Errorf("FAIL! expected not %v, got %v", nil, err)
		}
		if err.Error() != "error insert chef" {
			t.Errorf("FAIL! expected: %v, got: %v", "error insert chef", err.Error())
		}
	})

	t.Run("positive test case", func(t *testing.T) {
		chefRepo := ChefRepo{}
		kitchenRepo := KitchenRepo{}
		chefUsecase := NewChefUsecase(&chefRepo, &kitchenRepo)
		chef := domain.Chef{Name: "andi"}

		chef, err := chefUsecase.AddChef(chef)
		if chef.ID != 100 {
			t.Errorf("FAIL! expected %v, got %v", 100, chef.ID)
		}
		if chef.Name != "andi" {
			t.Errorf("FAIL! expected %v, got %v", "andi", chef.Name)
		}
		if err != nil {
			t.Errorf("FAIL! expected %v, got %v", nil, err)
		}
	})
}

type ChefRepo struct{}

func (repository *ChefRepo) InsertChef(chef domain.Chef) (domain.Chef, error) {
	if chef.ID == -1 {
		return domain.Chef{}, errors.New("error insert chef")
	}

	chef.ID = 100
	return chef, nil
}

type KitchenRepo struct{}

func (r *KitchenRepo) AddChefToKitchen(chef *domain.Chef, msg string) {}
func (r *KitchenRepo) GetChefFromKitchen() *domain.Chef {
	return &domain.Chef{ID: 1, Name: "juna"}
}
