package services

import (
	"financial/internal/models"

	"errors"
	"testing"
)

func generateDeleteCardService() (
	DeleteCard,
	*CardRepositoryMock,
	*BankAccountRepositoryMock) {

	cardMock := &CardRepositoryMock{}
	bankAccMock := &BankAccountRepositoryMock{}
	service := NewDeleteCard(cardMock, bankAccMock)
	return service, cardMock, bankAccMock
}

func TestDeleteCardService_Run(t *testing.T) {
	const validUserID = uint(123)
	const validCardID = uint(1)

	t.Run("successful card deletion", func(t *testing.T) {
		service, cardMock, bankAccMock := generateDeleteCardService()

		bankAccMock.FindBankAccountByCardIDFunc =
			func(id uint) (models.BankAccount, error) {
				return models.BankAccount{UserID: validUserID}, nil
			}

		var deletedCardID uint
		cardMock.DeleteFunc = func(id uint) error {
			deletedCardID = id
			return nil
		}

		err := service.Run(validCardID, validUserID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if deletedCardID != validCardID {
			t.Errorf("expected card ID %d, got %d", validCardID, deletedCardID)
		}
	})

	t.Run("bank account not found", func(t *testing.T) {
		service, cardMock, bankAccMock := generateDeleteCardService()

		expectedErr := errors.New("bank account not found")
		bankAccMock.FindBankAccountByCardIDFunc =
			func(id uint) (models.BankAccount, error) {
				return models.BankAccount{}, expectedErr
			}

		cardMock.DeleteFunc = func(id uint) error {
			t.Error("Delete should not be called")
			return nil
		}

		err := service.Run(validCardID, validUserID)
		if err != expectedErr {
			t.Fatalf("expected error %v, got %v", expectedErr, err)
		}
	})

	t.Run("unauthorized user", func(t *testing.T) {
		service, cardMock, bankAccMock := generateDeleteCardService()

		bankAccMock.FindBankAccountByCardIDFunc =
			func(id uint) (models.BankAccount, error) {
				return models.BankAccount{UserID: validUserID + 1}, nil
			}

		cardMock.DeleteFunc = func(id uint) error {
			t.Error("Delete should not be called")
			return nil
		}

		err := service.Run(validCardID, validUserID)
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		expectedErrMsg := "cant delete card from another user"
		if err.Error() != expectedErrMsg {
			t.Errorf("expected error message %q, got %q", expectedErrMsg, err.Error())
		}
	})

	t.Run("error during deletion", func(t *testing.T) {
		service, cardMock, bankAccMock := generateDeleteCardService()

		bankAccMock.FindBankAccountByCardIDFunc =
			func(id uint) (models.BankAccount, error) {
				return models.BankAccount{UserID: validUserID}, nil
			}

		expectedErr := errors.New("delete failed")
		cardMock.DeleteFunc = func(id uint) error {
			return expectedErr
		}

		err := service.Run(validCardID, validUserID)
		if err != expectedErr {
			t.Fatalf("expected error %v, got %v", expectedErr, err)
		}
	})
}
