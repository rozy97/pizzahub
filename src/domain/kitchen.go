package domain

type KitchenRepository interface {
	AddChefToKitchen(chef *Chef, msg string)
	GetChefFromKitchen() *Chef
}
