package services

import (
	"errors"
	"financial/internal/db"
	"financial/internal/models"

	"github.com/shopspring/decimal"
)

type AddItemsToTransaction struct {
	itemRepo        db.ItemRepository
	transactionRepo db.TransactionRepository
	bankAccountRepo db.BankAccountRepository
}

func NewAddItemsToTransaction(
	itemRepo db.ItemRepository,
	transactionRepo db.TransactionRepository,
	bankAccountRepo db.BankAccountRepository,
) AddItemsToTransaction {
	return AddItemsToTransaction{itemRepo, transactionRepo, bankAccountRepo}
}

type NewItem struct {
	Name     string
	Value    decimal.Decimal
	Quantity uint
}

func (a *AddItemsToTransaction) Run(
	newItems []NewItem,
	transactionId uint,
	userId uint,
) error {
	bankAccount, err := a.bankAccountRepo.FindBankAccountByTransactionID(transactionId)
	if err != nil {
		return err
	}

	if bankAccount.UserID != userId {
		return errors.New("cant add items to another users transaction")
	}

	var items []models.Item
	for _, newItem := range newItems {
		items = append(items, models.Item{
			Name:          newItem.Name,
			Value:         newItem.Value,
			Quantity:      newItem.Quantity,
			TransactionID: transactionId,
		})
	}

	if _, err := a.itemRepo.CreateMultiple(items); err != nil {
		return err
	}

	return nil
}
