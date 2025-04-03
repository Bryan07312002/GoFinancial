package services

import (
	"errors"
	"testing"

	"financial/internal/models"
)

func generateDeleteBankAccountService() (
	DeleteBankAccount,
	*BankAccountRepositoryMock) {

	bankAccMock := &BankAccountRepositoryMock{}
	service := NewDeleteBankAccountService(bankAccMock)
	return service, bankAccMock
}

func TestDeleteBankAccountService_Run(t *testing.T) {
	const validUserID = uint(123)
	const validBankAccountID = uint(1)

	t.Run("successful bank account deletion", func(t *testing.T) {
		service, bankAccMock := generateDeleteBankAccountService()

		bankAccMock.FindByIDFunc = func(id, userID uint) (models.BankAccount, error) {
			return models.BankAccount{UserID: validUserID}, nil
		}

		var deletedID uint
		bankAccMock.DeleteFunc = func(id uint) error {
			deletedID = id
			return nil
		}

		err := service.Run(validBankAccountID, validUserID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if deletedID != validBankAccountID {
			t.Errorf("expected ID %d, got %d", validBankAccountID, deletedID)
		}
	})

	t.Run("bank account not found", func(t *testing.T) {
		service, bankAccMock := generateDeleteBankAccountService()

		expectedErr := errors.New("not found")
		bankAccMock.FindByIDFunc = func(id, UserID uint) (models.BankAccount, error) {
			return models.BankAccount{}, expectedErr
		}

		bankAccMock.DeleteFunc = func(id uint) error {
			t.Error("Delete should not be called")
			return nil
		}

		err := service.Run(validBankAccountID, validUserID)
		if err != expectedErr {
			t.Fatalf("expected error %v, got %v", expectedErr, err)
		}
	})

	t.Run("unauthorized user", func(t *testing.T) {
		service, bankAccMock := generateDeleteBankAccountService()

		bankAccMock.FindByIDFunc = func(id, UserID  uint) (models.BankAccount, error) {
			return models.BankAccount{UserID: validUserID + 1}, nil
		}

		bankAccMock.DeleteFunc = func(id uint) error {
			t.Error("Delete should not be called")
			return nil
		}

		err := service.Run(validBankAccountID, validUserID)
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		expectedErrMsg := "cant delete bank account from another user"
		if err.Error() != expectedErrMsg {
			t.Errorf("expected error message %q, got %q", expectedErrMsg, err.Error())
		}
	})

	t.Run("error during deletion", func(t *testing.T) {
		service, bankAccMock := generateDeleteBankAccountService()

		bankAccMock.FindByIDFunc = func(id, userID uint) (models.BankAccount, error) {
			return models.BankAccount{UserID: validUserID}, nil
		}

		expectedErr := errors.New("delete failed")
		bankAccMock.DeleteFunc = func(id uint) error {
			return expectedErr
		}

		err := service.Run(validBankAccountID, validUserID)
		if err != expectedErr {
			t.Fatalf("expected error %v, got %v", expectedErr, err)
		}
	})
}
