package services

import (
	"financial/internal/db"
	"financial/internal/models"
)

type CreateCardService struct {
	cardRepo db.CardRepository
}

func NewCreateCardService(cardRepo db.CardRepository) CreateCardService {
	return CreateCardService{cardRepo}
}

type CreateCard struct {
	BankAccountID uint

	Name        string
	Description string
}

func (c *CreateCardService) Run(newCard CreateCard) error {
	_, err := c.cardRepo.Create(&models.Card{
		Name:          newCard.Name,
		Description:   newCard.Description,
		BankAccountID: newCard.BankAccountID,
	})

	return err
}
