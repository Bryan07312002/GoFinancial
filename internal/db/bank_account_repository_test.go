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

func generateTestBankAccountRepository(
	t *testing.T) (BankAccountRepository, *gorm.DB) {
	conn, err := gorm.Open(sqlite.Open("file::memory:?"),
		&gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					LogLevel: logger.Silent, // Silent to avoid logs in tests
				},
			),
		},
	)

	if err != nil {
		t.Fatalf("failed to connect to in-memory database: %v", err)
	}

	err = conn.AutoMigrate(
		&BankAccountTable{},
		&CardTable{},
		&TransactionTable{},
	)

	if err != nil {
		t.Fatalf("failed to migrate database schema: %v", err)
	}

	return NewBankAccountRepository(conn), conn
}

func TestBankAccountRepository(t *testing.T) {
	t.Run("CreateBankAccount", func(t *testing.T) {
		bankAccountRepo, _ := generateTestBankAccountRepository(t)
		bankAccount := models.BankAccount{
			Name:        "Test Account",
			Description: "Test Description",
			UserID:      1,
		}

		bankAccountID, err := bankAccountRepo.Create(bankAccount)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if bankAccountID <= 0 {
			t.Errorf("expected bank account ID to be greater than 0, got %v", bankAccountID)
		}
	})

	t.Run("FindBankAccountByID", func(t *testing.T) {
		bankAccountRepo, _ := generateTestBankAccountRepository(t)

		// Create a bank account
		bankAccount := models.BankAccount{
			Name:        "Findable Account",
			Description: "Account to be found",
			UserID:      2,
		}

		bankAccountID, err := bankAccountRepo.Create(bankAccount)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		foundAccount, err := bankAccountRepo.FindByID(bankAccountID)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if foundAccount.Name != bankAccount.Name {
			t.Errorf("expected name %v, got %v", bankAccount.Name, foundAccount.Name)
		}

		if foundAccount.Description != bankAccount.Description {
			t.Errorf("expected description %v, got %v", bankAccount.Description, foundAccount.Description)
		}
	})

	t.Run("DeleteBankAccount", func(t *testing.T) {
		bankAccountRepo, _ := generateTestBankAccountRepository(t)

		// Create a bank account to delete
		bankAccount := models.BankAccount{
			Name:        "Deletable Account",
			Description: "Account to be deleted",
			UserID:      3,
		}
		bankAccountID, err := bankAccountRepo.Create(bankAccount)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		err = bankAccountRepo.Delete(bankAccountID)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		_, err = bankAccountRepo.FindByID(bankAccountID)
		if err == nil {
			t.Error("expected error when finding deleted bank account, got nil")
		}
	})

	t.Run("FindBankAccountByCardID", func(t *testing.T) {
		bankAccountRepo, conn := generateTestBankAccountRepository(t)

		// Create a bank account
		bankAccount := models.BankAccount{
			Name:        "Account With Card",
			Description: "Account linked to card",
			UserID:      4,
		}
		bankAccountID, err := bankAccountRepo.Create(bankAccount)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		// Now create a card that references the bank account (using CardTable and BankAccountTable)
		card := &CardTable{
			BankAccountID: bankAccountID,
			Name:          "Test Card",
			Description:   "Card linked to account",
		}

		err = conn.Create(card).Error
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		// Now query the bank account by card ID
		foundBankAccount, err := bankAccountRepo.FindBankAccountByCardID(card.ID)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if foundBankAccount.ID != bankAccountID {
			t.Errorf("expected bank account ID %v, got %v", bankAccountID, foundBankAccount.ID)
		}
	})
}
