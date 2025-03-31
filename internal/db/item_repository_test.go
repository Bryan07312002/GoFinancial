package db

import (
	"financial/internal/models"
	"github.com/shopspring/decimal"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
)

func generateTestItemRepository(t *testing.T) ItemRepository {
	conn, err := gorm.Open(sqlite.Open("file::memory:?"),
		&gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					LogLevel: logger.Silent,
				},
			),
		},
	)

	if err != nil {
		t.Fatalf("failed to connect to in-memory database: %v", err)
	}


	err = applyAutoMigrate(conn)
	if err != nil {
		t.Fatalf("failed to migrate database schema: %v", err)
	}

	return NewItemRepository(conn)
}

func TestItemRepository(t *testing.T) {
	t.Run("CreateItem", func(t *testing.T) {
		itemRepo := generateTestItemRepository(t)
		item := models.Item{
			Name:          "Test Item",
			Value:         decimal.NewFromFloat(100.50),
			Quantity:      5,
			TransactionID: 1,
		}
		items := make([]models.Item, 0)
		items = append(items, item)
		itemID, err := itemRepo.CreateMultiple(items)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if itemID[0] <= 0 {
			t.Errorf("expected item ID to be greater than 0, got %v", itemID)
		}
	})

	t.Run("FindItemByID", func(t *testing.T) {
		itemRepo := generateTestItemRepository(t)
		// Create an item
		item := models.Item{
			Name:          "Findable Item",
			Value:         decimal.NewFromFloat(50.75),
			Quantity:      3,
			TransactionID: 2,
		}
		items := make([]models.Item, 0)
		items = append(items, item)
		itemID, err := itemRepo.CreateMultiple(items)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		foundItem, err := itemRepo.FindByID(itemID[0])
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if foundItem.Name != item.Name {
			t.Errorf("expected name %v, got %v", item.Name, foundItem.Name)
		}

		if foundItem.Value.Compare(item.Value) != 0 {
			t.Errorf("expected value %v, got %v", item.Value, foundItem.Value)
		}
		if foundItem.Quantity != item.Quantity {
			t.Errorf("expected quantity %v, got %v", item.Quantity, foundItem.Quantity)
		}
	})

	t.Run("DeleteItem", func(t *testing.T) {
		itemRepo := generateTestItemRepository(t)

		// Create an item to delete
		item := models.Item{
			Name:          "Deletable Item",
			Value:         decimal.NewFromFloat(75.25),
			Quantity:      10,
			TransactionID: 3,
		}
		items := make([]models.Item, 0)
		items = append(items, item)
		itemID, err := itemRepo.CreateMultiple(items)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		err = itemRepo.Delete(itemID[0])
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		_, err = itemRepo.FindByID(itemID[0])
		if err == nil {
			t.Error("expected error when finding deleted item, got nil")
		}
	})
}
