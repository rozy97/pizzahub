package domain

type Chef struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ChefUsecase interface {
	AddChef(chef Chef) (Chef, error)
}

type ChefRepository interface {
	InsertChef(chef Chef) (Chef, error)
	// GetChefs() ([]Chef, error)
	// GetChef(ID int) (Chef, error)
}
