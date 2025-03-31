package services

import (
	"financial/internal/db"
	"financial/internal/models"

	"errors"
	"fmt"
	"github.com/shopspring/decimal"
)

type AddItemsToTransaction interface {
	Run(newItems []NewItem, transactionId uint, userId uint) error
}

type addItemsToTransaction struct {
	itemRepo        db.ItemRepository
	badgeRepo       db.BadgeRepository
	transactionRepo db.TransactionRepository
	bankAccountRepo db.BankAccountRepository
}

func NewAddItemsToTransaction(
	itemRepo db.ItemRepository,
	badgeRepo db.BadgeRepository,
	transactionRepo db.TransactionRepository,
	bankAccountRepo db.BankAccountRepository,
) AddItemsToTransaction {
	return &addItemsToTransaction{
		itemRepo,
		badgeRepo,
		transactionRepo,
		bankAccountRepo,
	}
}

type NewItem struct {
	Name     string          `json:"name"`
	Value    decimal.Decimal `json:"value"`
	Quantity uint            `json:"quantity"`
	Badges   *[]uint         `json:"badges"`
}

func (a *addItemsToTransaction) Run(
	newItems []NewItem,
	transactionId uint,
	userId uint,
) error {
	bankAccount, err := a.bankAccountRepo.
		FindBankAccountByTransactionID(transactionId)
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

	ids, err := a.itemRepo.CreateMultiple(items)
	if err != nil {
		return err
	}

	for i, id := range ids {
		fmt.Printf("%+v", newItems[i].Badges)
		items[i].ID = id
	}

	for i, item := range items {
		if *newItems[i].Badges == nil {
			continue
		}

		for badgeID := range *newItems[i].Badges {
			a.badgeRepo.LinkItemToBadge(item.ID, uint(badgeID))
		}
	}

	return nil
}
