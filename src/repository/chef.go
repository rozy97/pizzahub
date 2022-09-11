package repository

import "github.com/rozy97/pizzahub/src/domain"

type ChefRepository struct {
}

var chefs []domain.Chef

func NewChefRepository() *ChefRepository {
	return &ChefRepository{}
}

func (repository *ChefRepository) InsertChef(chef domain.Chef) (domain.Chef, error) {
	id := len(chefs) + 1
	chef.ID = id
	chefs = append(chefs, chef)
	return chef, nil
}
