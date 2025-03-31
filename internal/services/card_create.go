package services

import (
	"financial/internal/db"
	"financial/internal/models"
)

type CreateCard interface {
	Run(newCard CreateCardDto) error
}

type createCard struct {
	cardRepo db.CardRepository
}

func NewCreateCardService(cardRepo db.CardRepository) CreateCard {
	return &createCard{cardRepo}
}

type CreateCardDto struct {
	BankAccountID uint

	Name        string
	Description string
}

func (c *createCard) Run(newCard CreateCardDto) error {
	_, err := c.cardRepo.Create(&models.Card{
		Name:          newCard.Name,
		Description:   newCard.Description,
		BankAccountID: newCard.BankAccountID,
	})

	return err
}
