package services

import (
	"errors"
	"financial/internal/db"
)

type DeleteItem interface {
	Run(id uint, userId uint) error
}

type deleteItem struct {
	bankRepo db.BankAccountRepository
	itemRepo db.ItemRepository
}

func NewDeleteItem(
	itemRepo db.ItemRepository,
	bankRepo db.BankAccountRepository,
) DeleteItem {
	return &deleteItem{bankRepo, itemRepo}
}

func (d *deleteItem) Run(id uint, userId uint) error {
	item, err := d.itemRepo.FindByID(id)
	if err != nil {
		return err
	}

	bankAccount, err := d.bankRepo.FindBankAccountByTransactionID(item.TransactionID)
	if err != nil {
		return err
	}

	if bankAccount.UserID != userId {
		return errors.New("cant delete item from another users transaction")
	}

	return d.itemRepo.Delete(id)
}
