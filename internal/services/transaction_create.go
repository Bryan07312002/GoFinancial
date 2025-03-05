package services

import (
	"financial/internal/db"
	"financial/internal/models"
	"financial/internal/utils"

	"errors"
	"time"

	"github.com/shopspring/decimal"
)

type CreateTransactionService struct {
	transactionRepo db.TransactionRepository
	bankAccountRepo db.BankAccountRepository
}

func NewCreateTransactionService(
	transactionRepo db.TransactionRepository,
	bankAccountRepo db.BankAccountRepository,
) CreateTransactionService {
	return CreateTransactionService{transactionRepo, bankAccountRepo}
}

type CreateTransaction struct {
	Type          string
	Value         decimal.Decimal
	BankAccountID uint
	Establishment string

	Date   *string
	CardID *uint
	Credit *bool
	Method *string
}

// TODO: check date of transaction and update all balances after the transaction
func (c *CreateTransactionService) Run(
	newTransaction CreateTransaction,
	userId uint,
) error {

	bankAccount, err := c.bankAccountRepo.FindByID(newTransaction.BankAccountID)
	if err != nil {
		return err
	}

	if bankAccount.UserID != userId {
		return errors.New("Cant create transaction for another user bank account")
	}

	// FIXME: TODO: check if date is valid
	if newTransaction.Date == nil {
		now := time.Now().String()
		newTransaction.Date = &now
	}

	if newTransaction.Credit == nil {
		credit := false
		newTransaction.Credit = &credit
	}

	if !models.TransactionType(newTransaction.Type).IsValid() {
		return errors.New("invalid transaction type")
	}

	var method models.PaymentMethod
	if newTransaction.Method == nil {
		method = models.Other
	} else {
		if !models.PaymentMethod(*newTransaction.Method).IsValid() {
			return errors.New("invalid payment method")
		}

		method = models.PaymentMethod(*newTransaction.Method)
	}

	time, err := utils.ParseTime(*newTransaction.Date)
	if err != nil {
		return err
	}

	_, err = c.transactionRepo.Create(&models.Transaction{
		Type:          models.TransactionType(newTransaction.Type),
		Method:        method,
		Establishment: *&newTransaction.Establishment,
		Value:         newTransaction.Value,
		BankAccountID: newTransaction.BankAccountID,

		Date:   time,
		CardID: *&newTransaction.CardID,
		Credit: *newTransaction.Credit,
	})

	return err
}
