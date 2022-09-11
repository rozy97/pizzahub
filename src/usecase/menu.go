package usecase

import (
	"github.com/rozy97/pizzahub/src/domain"
)

type MenuUsecase struct {
	repository domain.MenuRepository
}

func NewMenuUsecase(repository domain.MenuRepository) *MenuUsecase {
	return &MenuUsecase{
		repository: repository,
	}
}

func (usecase *MenuUsecase) GetMenus() ([]domain.Menu, error) {
	return usecase.repository.GetMenus()
}
