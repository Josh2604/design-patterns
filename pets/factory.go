package pets

import "github.com/Josh2604/design-patterns/models"

func NewPet(species string) *models.Pet {
	pet := models.Pet{
		Species:     species,
		Breed:       "",
		MinWeight:   0,
		MaxWeight:   0,
		Description: "No description entered yet",
		LifeSpan:    0,
	}

	return &pet
}
