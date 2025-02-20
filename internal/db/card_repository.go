package db

import (
	"financial/internal/models"
	"gorm.io/gorm"
)

type CardRepository interface {
	Create(card *models.Card) (uint, error)
	FindByID(id uint) (*models.Card, error)
	Delete(id uint) error
}

func NewCardRepository(db *gorm.DB) CardRepository {
	return &cardRepository{db}
}

type cardRepository struct {
	db *gorm.DB
}

func (c *cardRepository) Create(card *models.Card) (uint, error) {
	cardTableInstance := &CardTable{
		BankAccountID: card.ID,
		Name:          card.Name,
		Description:   card.Description,
	}

	if err := c.db.Create(cardTableInstance).Error; err != nil {
		return 0, err
	}

	return cardTableInstance.ID, nil
}

func (c *cardRepository) FindByID(id uint) (*models.Card, error) {
	// Create a variable to hold the result
	var cardTableInstance CardTable

	// Query the database for the CardTable instance with the given ID
	if err := c.db.First(&cardTableInstance, id).Error; err != nil {
		return nil, err
	}

	// Convert the CardTable instance to a models.Card instance
	card := &models.Card{
		ID:          cardTableInstance.BankAccountID,
		Name:        cardTableInstance.Name,
		Description: cardTableInstance.Description,
	}

	return card, nil
}

func (c *cardRepository) Delete(id uint) error {
	// Attempt to delete the bank account by ID
	result := c.db.Delete(&CardTable{}, id)

	// Handle database errors (e.g., connection issues)
	if result.Error != nil {
		return result.Error
	}

	// Check if no rows were affected (record not found)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	// Successfully deleted
	return nil
}
