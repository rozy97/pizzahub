package usecase

import "github.com/rozy97/pizzahub/src/domain"

type ChefUsecase struct {
	repository  domain.ChefRepository
	kitchenRepo domain.KitchenRepository
}

func NewChefUsecase(repository domain.ChefRepository, kitchenRepo domain.KitchenRepository) *ChefUsecase {
	return &ChefUsecase{
		repository:  repository,
		kitchenRepo: kitchenRepo,
	}
}

func (usecase *ChefUsecase) AddChef(chef domain.Chef) (domain.Chef, error) {
	chef, err := usecase.repository.InsertChef(chef)
	if err != nil {
		return chef, err
	}

	usecase.kitchenRepo.AddChefToKitchen(&chef, "newly added")
	return chef, nil
}
