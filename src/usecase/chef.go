package usecase

import "github.com/rozy97/pizzahub/src/domain"

type ChefUsecase struct {
	repository domain.ChefRepository
}

func NewChefUsecase(repository domain.ChefRepository) *ChefUsecase {
	return &ChefUsecase{
		repository: repository,
	}
}

func (usecase *ChefUsecase) AddChef(chef domain.Chef) (domain.Chef, error) {
	return usecase.repository.InsertChef(chef)
}
