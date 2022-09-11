package repository

import (
	"log"

	"github.com/rozy97/pizzahub/src/domain"
)

type KitchenRepository struct {
	kitchen chan *domain.Chef
}

func NewKitchenRepository(kitchen chan *domain.Chef) *KitchenRepository {
	return &KitchenRepository{
		kitchen: kitchen,
	}
}

func (repository *KitchenRepository) AddChefToKitchen(chef *domain.Chef, msg string) {
	repository.kitchen <- chef
	log.Printf("chef %s %s to kitchen", chef.Name, msg)

}

func (repository *KitchenRepository) GetChefFromKitchen() *domain.Chef {
	chef := <-repository.kitchen
	return chef
}
