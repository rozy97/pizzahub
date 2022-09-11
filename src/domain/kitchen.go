package domain

type KitchenRepository interface {
	PublishChef(kitchen chan Chef, chef Chef)
	ConsumeChef(Kitchen chan Chef) Chef
}
