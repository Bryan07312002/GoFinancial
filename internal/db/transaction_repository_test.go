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
	"time"
)

func generateTestTransactionRepository(t *testing.T) TransactionRepository {
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

	// Migrate the schema
	err = conn.AutoMigrate(&TransactionTable{})
	if err != nil {
		t.Fatalf("failed to migrate schema: %v", err)
	}

	return NewTransactionRepository(conn)
}

func TestTransactionRepository(t *testing.T) {
	t.Run("CreateTransaction", func(t *testing.T) {
		transactionRepo := generateTestTransactionRepository(t)

		transaction := models.Transaction{
			Type:          models.Income,
			Method:        models.CreditCard,
			Credit:        true,
			Value:         decimal.NewFromFloat(100.50),
			Date:          time.Now(),
			CardID:        nil,
			BankAccountID: 1,
		}

		// Create a new transaction
		transactionID, err := transactionRepo.Create(&transaction)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if transactionID <= 0 {
			t.Errorf("expected transaction ID to be greater than 0, got %v", transactionID)
		}
	})

	t.Run("FindTransactionByID", func(t *testing.T) {
		transactionRepo := generateTestTransactionRepository(t)

		// Create a new transaction
		transaction := models.Transaction{
			Type:          models.Expense,
			Method:        models.DebitCard,
			Credit:        false,
			Value:         decimal.NewFromFloat(50.75),
			Date:          time.Now(),
			CardID:        nil,
			BankAccountID: 2,
		}

		// Create the transaction and get the ID
		transactionID, err := transactionRepo.Create(&transaction)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		// Find the transaction by ID
		foundTransaction, err := transactionRepo.FindByID(transactionID)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		// Validate that the retrieved transaction matches the created one
		if foundTransaction.Type != transaction.Type {
			t.Errorf("expected type %v, got %v", transaction.Type, foundTransaction.Type)
		}
		if foundTransaction.Value.Compare(transaction.Value) != 0 {
			t.Errorf("expected value %v, got %v", transaction.Value, foundTransaction.Value)
		}
	})

	t.Run("DeleteTransaction", func(t *testing.T) {
		transactionRepo := generateTestTransactionRepository(t)

		// Create a transaction to delete
		transaction := models.Transaction{
			Type:          models.Expense,
			Method:        models.DebitCard,
			Credit:        false,
			Value:         decimal.NewFromFloat(75.25),
			Date:          time.Now(),
			CardID:        nil,
			BankAccountID: 3,
		}

		// Create the transaction
		transactionID, err := transactionRepo.Create(&transaction)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		// Delete the transaction
		err = transactionRepo.Delete(transactionID)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		// Try to find the deleted transaction (should return an error)
		_, err = transactionRepo.FindByID(transactionID)
		if err == nil {
			t.Error("expected error when finding deleted transaction, got nil")
		}
	})
}
