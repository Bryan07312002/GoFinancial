package db

import (
	"financial/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
)

func generateTestCardRepository(t *testing.T) CardRepository {
	// Create an in-memory SQLite database
	conn, err := gorm.Open(sqlite.Open("file::memory:?"), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Silent, // Silent to avoid logs in tests
			},
		),
	})
	if err != nil {
		t.Fatalf("failed to connect to in-memory database: %v", err)
	}

	// Migrate the schema for CardTable
	err = conn.AutoMigrate(&CardTable{})
	if err != nil {
		t.Fatalf("failed to migrate schema: %v", err)
	}

	return NewCardRepository(conn)
}

func TestCardRepository(t *testing.T) {
	t.Run("CreateCard", func(t *testing.T) {
		cardRepo := generateTestCardRepository(t)

		card := models.Card{
			ID:          1,
			Name:        "Test Card",
			Description: "This is a test card.",
		}

		// Create the card and check for errors
		cardID, err := cardRepo.Create(&card)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if cardID <= 0 {
			t.Errorf("expected card ID to be greater than 0, got %v", cardID)
		}
	})

	t.Run("FindCardByID", func(t *testing.T) {
		cardRepo := generateTestCardRepository(t)

		// Create a card to find later
		card := models.Card{
			ID:          2,
			Name:        "Test Card 2",
			Description: "This is another test card.",
		}
		cardID, err := cardRepo.Create(&card)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		// Find the card by its ID
		foundCard, err := cardRepo.FindByID(cardID)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		// Validate the returned card data
		if foundCard.Name != card.Name {
			t.Errorf("expected card name %v, got %v", card.Name, foundCard.Name)
		}
		if foundCard.Description != card.Description {
			t.Errorf("expected card description %v, got %v", card.Description, foundCard.Description)
		}
	})

	t.Run("DeleteCard", func(t *testing.T) {
		cardRepo := generateTestCardRepository(t)

		// Create a card to delete later
		card := models.Card{
			ID:          3,
			Name:        "Test Card 3",
			Description: "This card will be deleted.",
		}
		cardID, err := cardRepo.Create(&card)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		// Delete the card
		err = cardRepo.Delete(cardID)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		// Try to find the deleted card (should return an error)
		_, err = cardRepo.FindByID(cardID)
		if err == nil {
			t.Error("expected error when finding deleted card, got nil")
		}
	})
}
