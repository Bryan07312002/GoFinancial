package services

import (
	"errors"
	"testing"

	"financial/internal/models"
)

// Helper function to generate service and mocks
func generateDeleteTransactionService() (
	DeleteTransaction,
	*BankAccountRepositoryMock,
	*TransactionRepositoryMock) {
	bankAccMock := &BankAccountRepositoryMock{}
	txMock := &TransactionRepositoryMock{}

	service := NewDeleteTransaction(bankAccMock, txMock)
	return service, bankAccMock, txMock
}

func TestDeleteTransactionService_Run(t *testing.T) {
	const validUserID = uint(123)
	const validTransactionID = uint(1)

	t.Run("successful transaction deletion", func(t *testing.T) {
		service, bankAccMock, txMock := generateDeleteTransactionService()

		bankAccMock.FindBankAccountByTransactionIDFunc =
			func(id uint) (models.BankAccount, error) {
				return models.BankAccount{UserID: validUserID}, nil
			}

		var deletedTxID uint
		txMock.DeleteFunc = func(id uint) error {
			deletedTxID = id
			return nil
		}

		err := service.Run(validTransactionID, validUserID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if deletedTxID != validTransactionID {
			t.Errorf("expected transaction ID %d, got %d", validTransactionID, deletedTxID)
		}
	})

	t.Run("bank account not found", func(t *testing.T) {
		service, bankAccMock, txMock := generateDeleteTransactionService()

		expectedErr := errors.New("bank account not found")
		bankAccMock.FindBankAccountByTransactionIDFunc =
			func(id uint) (models.BankAccount, error) {
				return models.BankAccount{}, expectedErr
			}

		txMock.DeleteFunc = func(id uint) error {
			t.Error("Delete should not be called")
			return nil
		}

		err := service.Run(validTransactionID, validUserID)
		if err != expectedErr {
			t.Fatalf("expected error %v, got %v", expectedErr, err)
		}
	})

	t.Run("unauthorized user", func(t *testing.T) {
		service, bankAccMock, txMock := generateDeleteTransactionService()

		bankAccMock.FindBankAccountByTransactionIDFunc =
			func(id uint) (models.BankAccount, error) {
				return models.BankAccount{UserID: validUserID + 1}, nil // Different user ID
			}

		txMock.DeleteFunc = func(id uint) error {
			t.Error("Delete should not be called")
			return nil
		}

		err := service.Run(validTransactionID, validUserID)
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		expectedErrMsg := "cant delete another user transaction"
		if err.Error() != expectedErrMsg {
			t.Errorf("expected error message %q, got %q", expectedErrMsg, err.Error())
		}
	})

	t.Run("error during deletion", func(t *testing.T) {
		service, bankAccMock, txMock := generateDeleteTransactionService()

		bankAccMock.FindBankAccountByTransactionIDFunc =
			func(id uint) (models.BankAccount, error) {
				return models.BankAccount{UserID: validUserID}, nil
			}

		expectedErr := errors.New("delete failed")
		txMock.DeleteFunc = func(id uint) error {
			return expectedErr
		}

		err := service.Run(validTransactionID, validUserID)
		if err != expectedErr {
			t.Fatalf("expected error %v, got %v", expectedErr, err)
		}
	})
}
