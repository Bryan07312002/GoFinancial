package db

import (
	"financial/internal/models"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
)

func generateTestBadgeRepository(t *testing.T) (BadgeRepository, *gorm.DB) {
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

	// Migrate the schema for BadgeTable, ItemTable, and ItemBadgeTable
	err = applyAutoMigrate(conn)
    if err != nil {
		t.Fatalf("failed to migrate schema: %v", err)
	}

	return NewBadgeRepository(conn), conn
}

func createFakeData(
	con *gorm.DB,
	t *testing.T,
) (UserTable, BankAccountTable, TransactionTable, ItemTable) {
	user := UserTable{
		Name:     "test",
		Password: "test-pass",
	}
	if err := con.Create(&user).Error; err != nil {
		t.Error(err)
	}

	bankAcc := BankAccountTable{
		UserID:      user.ID,
		Name:        "test-acc",
		Description: "",
	}
	if err := con.Create(&bankAcc).Error; err != nil {
		t.Error(err)
	}

	method := "credit_card"
	trans := TransactionTable{
		Credit:        false,
		CardID:        nil,
		BankAccountID: bankAcc.ID,
		Type:          "expanse",
		Method:        &method,
		Value:         decimal.NewFromFloat(100.0),
		Date:          time.Now(),
	}
	if err := con.Create(&trans).Error; err != nil {
		t.Error(err)
	}

	item := ItemTable{
		Name:          "Test Item",
		Value:         decimal.NewFromFloat(100.0),
		Quantity:      1,
		TransactionID: 1, // assuming a transaction with ID 1
	}
	if err := con.Create(&item).Error; err != nil {
		t.Error(err)
	}

	return user, bankAcc, trans, item
}

func TestBadgeRepository(t *testing.T) {
	t.Run("Should create badge successfully", func(t *testing.T) {
		badgeRepo, _ := generateTestBadgeRepository(t)

		badge := models.Badge{
			Name: "Gold Badge",
		}

		// Create the badge and check for errors
		badgeID, err := badgeRepo.Create(&badge)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if badgeID <= 0 {
			t.Errorf("expected badge ID to be greater than 0, got %v", badgeID)
		}
	})

	t.Run("Should get badges from item", func(t *testing.T) {
		badgeRepo, conn := generateTestBadgeRepository(t)
		_, _, _, itemTable := createFakeData(conn, t)

		// Create a badge
		badge := models.Badge{
			ID:   1,
			Name: "Silver Badge",
		}

		badgeID, err := badgeRepo.Create(&badge)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		badgeRepo.LinkItemToBadge(itemTable.ID, badgeID)

		badges, err := badgeRepo.FindByItem(itemTable.ID)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(badges) == 0 {
			t.Error("expected at least one badge, got none")
		}
	})

	t.Run("Should not get badges to unrelated items", func(t *testing.T) {
		badgeRepo, conn := generateTestBadgeRepository(t)
		_, _, transactionTable, itemTable := createFakeData(conn, t)

		// create second item with badge
		badge := models.Badge{
			ID:   1,
			Name: "Silver Badge",
		}

		badgeID, err := badgeRepo.Create(&badge)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		badgeRepo.LinkItemToBadge(itemTable.ID, badgeID)

		// create second item with another badge
		item2 := ItemTable{
			Name:          "Test Item",
			Value:         decimal.NewFromFloat(100.0),
			Quantity:      1,
			TransactionID: transactionTable.ID,
		}
		if err := conn.Create(&item2).Error; err != nil {
			t.Error(err)
		}

		badgeID2, err := badgeRepo.Create(&models.Badge{
			ID:   1,
			Name: "Silver Badge",
		})
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		badgeRepo.LinkItemToBadge(item2.ID, badgeID2)

		badges, err := badgeRepo.FindByItem(itemTable.ID)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(badges) != 1 {
			t.Errorf("expected item to have 1 badge but got: %d", len(badges))
		}
	})

	t.Run("FindBadgeByItem - Fail (No badges for item)", func(t *testing.T) {
		badgeRepo, _ := generateTestBadgeRepository(t)

		// Try to find badges for an item with ID 999 (which doesn't exist)
		badges, err := badgeRepo.FindByItem(999)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(badges) != 0 {
			t.Errorf("expected no badges, got %v", badges)
		}
	})

	t.Run("FindBadgeByTransaction - Success", func(t *testing.T) {
		badgeRepo, conn := generateTestBadgeRepository(t)
		_, _, transactionTable, itemTable := createFakeData(conn, t)

		badge := models.Badge{
			Name: "Silver Badge",
		}

		badgeID, err := badgeRepo.Create(&badge)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		badgeRepo.LinkItemToBadge(itemTable.ID, badgeID)
		badges, err := badgeRepo.FindByTransaction(transactionTable.ID)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(badges) == 0 {
			t.Error("expected at least one badge, got none")
		}
	})

	t.Run("FindBadgeByTransaction - Fail (No badges for transaction)", func(t *testing.T) {
		badgeRepo, _ := generateTestBadgeRepository(t)

		// Try to find badges for a transaction with ID 999 (which doesn't exist)
		badges, err := badgeRepo.FindByTransaction(999)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(badges) != 0 {
			t.Errorf("expected no badges, got %v", badges)
		}
	})

	t.Run("CreateMultipleBadges - Success", func(t *testing.T) {
		badgeRepo, _ := generateTestBadgeRepository(t)

		badges := []models.Badge{
			{Name: "Gold Badge"},
			{Name: "Silver Badge"},
			{Name: "Bronze Badge"},
		}

		// Create multiple badges
		createdIDs, err := badgeRepo.CreateMultiple(badges)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(createdIDs) != 3 {
			t.Errorf("expected 3 badges to be created, got %v", len(createdIDs))
		}
	})

	t.Run("CreateMultipleBadges - Fail (Empty list)", func(t *testing.T) {
		badgeRepo, _ := generateTestBadgeRepository(t)

		// Create an empty list of badges
		badges := []models.Badge{}

		// Try to create an empty list
		createdIDs, err := badgeRepo.CreateMultiple(badges)
		if err == nil {
			t.Error("expected error when creating empty list, got nil")
		}

		if len(createdIDs) != 0 {
			t.Errorf("expected no IDs to be returned, got %v", createdIDs)
		}
	})

	t.Run("DeleteBadge - Success", func(t *testing.T) {
		badgeRepo, _ := generateTestBadgeRepository(t)

		// Create a badge
		badge := models.Badge{
			Name: "Delete Test Badge",
		}
		badgeID, err := badgeRepo.Create(&badge)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		// Delete the badge
		err = badgeRepo.Delete(badgeID)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		// Try to find the deleted badge (should return an error)
		_, err = badgeRepo.FindByID(badgeID, 1)
		if err == nil {
			t.Error("expected error when finding deleted badge, got nil")
		}
	})

	t.Run("DeleteBadge - Fail (Non-existent ID)", func(t *testing.T) {
		badgeRepo, _ := generateTestBadgeRepository(t)

		// Try to delete a badge that doesn't exist
		err := badgeRepo.Delete(999)
		if err == nil {
			t.Error("expected error when deleting non-existent badge, got nil")
		}
	})

	t.Run("FindBadgeByID - Success", func(t *testing.T) {
		badgeRepo, _ := generateTestBadgeRepository(t)

		// Create a badge
		badge := models.Badge{
			Name: "Gold Badge",
		}
		badgeID, err := badgeRepo.Create(&badge)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		// Find the badge by ID
		foundBadge, err := badgeRepo.FindByID(badgeID, 1)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if foundBadge.ID != badgeID {
			t.Errorf("expected badge ID %v, got %v", badgeID, foundBadge.ID)
		}

		if foundBadge.Name != "Gold Badge" {
			t.Errorf("expected badge name 'Gold Badge', got %v", foundBadge.Name)
		}
	})

	t.Run("FindBadgeByID - Fail (Non-existent ID)", func(t *testing.T) {
		badgeRepo, _ := generateTestBadgeRepository(t)

		// Try to find a badge that doesn't exist
		_, err := badgeRepo.FindByID(999, 1) // Assume this ID doesn't exist
		if err == nil {
			t.Error("expected error when finding badge with non-existent ID, got nil")
		}
	})
}
