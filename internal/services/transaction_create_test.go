package services

import (
	"financial/internal/models"

	"github.com/shopspring/decimal"
	"testing"
	"time"
)

func generateCreateTransactionService() (
	CreateTransactionService,
	*BankAccountRepositoryMock,
	*TransactionRepositoryMock) {
	bankAccMock := &BankAccountRepositoryMock{}
	txMock := &TransactionRepositoryMock{}

	service := NewCreateTransactionService(txMock, bankAccMock)
	return service, bankAccMock, txMock
}

func TestCreateTransactionService(t *testing.T) {
	const validUserID = uint(123)
	const validBankAccountID = uint(1)
	validValue := decimal.NewFromInt(100)

	t.Run("successful expense transaction creation", func(t *testing.T) {
		service, bankAccMock, txMock := generateCreateTransactionService()

		bankAccMock.FindByIDFunc = func(id uint) (models.BankAccount, error) {
			return models.BankAccount{UserID: validUserID}, nil
		}

		var createdTx *models.Transaction
		txMock.CreateFunc = func(tx *models.Transaction) (uint, error) {
			createdTx = tx
			return 1, nil
		}

		method := models.DebitCard.String()
		input := CreateTransaction{
			Type:          models.Expense.String(),
			Method:        &method,
			Value:         validValue,
			BankAccountID: validBankAccountID,
		}

		err := service.Run(input, validUserID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if createdTx.Type != models.Expense {
			t.Errorf(
				"expected transaction type %q, got %q",
				models.Expense,
				createdTx.Type)
		}

		if createdTx.Method != models.DebitCard {
			t.Errorf(
				"expected payment method %q, got %q",
				models.DebitCard,
				createdTx.Method)
		}
	})

	t.Run("successful transfer transaction", func(t *testing.T) {
		service, bankAccMock, txMock := generateCreateTransactionService()

		bankAccMock.FindByIDFunc = func(id uint) (models.BankAccount, error) {
			return models.BankAccount{UserID: validUserID}, nil
		}

		var createdTx *models.Transaction
		txMock.CreateFunc = func(tx *models.Transaction) (uint, error) {
			createdTx = tx
			return 1, nil
		}

		input := CreateTransaction{
			Type:          models.Transfer.String(),
			Value:         validValue,
			BankAccountID: validBankAccountID,
		}

		err := service.Run(input, validUserID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if createdTx.Type != models.Transfer {
			t.Errorf(
				"expected transaction type %q, got %q",
				models.Transfer,
				createdTx.Type)
		}
		if createdTx.Method != models.Other {
			t.Error("expected nil payment method for transfer")
		}
	})

	t.Run("invalid payment method for transfer", func(t *testing.T) {
		service, bankAccMock, txMock := generateCreateTransactionService()

		bankAccMock.FindByIDFunc = func(id uint) (models.BankAccount, error) {
			return models.BankAccount{UserID: validUserID}, nil
		}

		txMock.CreateFunc = func(
			transaction *models.Transaction) (uint, error) {
			return 0, nil
		}

		method := "invalid method"
		input := CreateTransaction{
			Type:          models.Transfer.String(),
			Method:        &method,
			Value:         validValue,
			BankAccountID: validBankAccountID,
		}

		err := service.Run(input, validUserID)

		if err == nil {
			t.Fatal("expected error for transfer with payment method")
		}
	})

	t.Run("default values when not provided", func(t *testing.T) {
		service, bankAccMock, txMock := generateCreateTransactionService()

		bankAccMock.FindByIDFunc = func(id uint) (models.BankAccount, error) {
			return models.BankAccount{UserID: validUserID}, nil
		}

		var createdTx *models.Transaction
		txMock.CreateFunc = func(tx *models.Transaction) (uint, error) {
			createdTx = tx
			return 1, nil
		}

		input := CreateTransaction{
			Type:          models.Income.String(),
			Value:         validValue,
			BankAccountID: validBankAccountID,
		}

		err := service.Run(input, validUserID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if time.Since(createdTx.Date) > time.Second {
			t.Error("expected Date to be set to current time")
		}
	})

	t.Run("custom date value", func(t *testing.T) {
		service, bankAccMock, txMock := generateCreateTransactionService()

		bankAccMock.FindByIDFunc = func(id uint) (models.BankAccount, error) {
			return models.BankAccount{UserID: validUserID}, nil
		}

		var _ *models.Transaction
		txMock.CreateFunc = func(tx *models.Transaction) (uint, error) {
			_ = tx
			return 1, nil
		}

		customDate := time.Now().String()
		input := CreateTransaction{
			Type:          models.Expense.String(),
			Value:         validValue,
			BankAccountID: validBankAccountID,
			Date:          &customDate,
		}

		err := service.Run(input, validUserID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// TODO: test if date is correct
		// if !createdTx.Date.Equal(customDate) {
		// 	t.Errorf("expected Date %v, got %v", customDate, createdTx.Date)
		// }
	})

	t.Run("invalid transaction type", func(t *testing.T) {
		service, bankAccMock, _ := generateCreateTransactionService()

		bankAccMock.FindByIDFunc = func(id uint) (models.BankAccount, error) {
			return models.BankAccount{UserID: validUserID}, nil
		}

		input := CreateTransaction{
			Type:          "invalid_type",
			BankAccountID: validBankAccountID,
			Value:         validValue,
		}

		err := service.Run(input, validUserID)
		if err == nil {
			t.Fatal("expected error for invalid transaction type")
		}
	})

	t.Run("decimal value handling", func(t *testing.T) {
		service, bankAccMock, txMock := generateCreateTransactionService()

		bankAccMock.FindByIDFunc = func(id uint) (models.BankAccount, error) {
			return models.BankAccount{UserID: validUserID}, nil
		}

		var createdTx *models.Transaction
		txMock.CreateFunc = func(tx *models.Transaction) (uint, error) {
			createdTx = tx
			return 1, nil
		}

		testValue := decimal.NewFromFloat(123.45)
		input := CreateTransaction{
			Type:          models.Income.String(),
			Value:         testValue,
			BankAccountID: validBankAccountID,
		}

		err := service.Run(input, validUserID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !createdTx.Value.Equal(testValue) {
			t.Errorf("expected Value %v, got %v", testValue, createdTx.Value)
		}
	})
}
