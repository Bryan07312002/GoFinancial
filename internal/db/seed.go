package db

import (
	"financial/internal/models"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func Seed(config map[string]string) {
	conn := InitDatabase(SqliteDriver, config)

	bankAccs := createBankAccounts(models.User{ID: 1}, conn)
	for _, bankAcc := range bankAccs {
		transactions := createTransactions(bankAcc, conn)
		for _, transaction := range transactions {
			items := createItems(transaction, conn)
			for _, item := range items {
				createBadges(item, conn)
			}
		}
	}

}

func createBankAccounts(user models.User, conn *gorm.DB) []models.BankAccount {
	bankAccs := [2]models.BankAccount{
		{
			UserID:      user.ID,
			Name:        "Bank number one",
			Description: "first bank account",
		},
		{
			UserID:      user.ID,
			Name:        "Bank number two",
			Description: "second bank account",
		}}

	bankAccRepo := NewBankAccountRepository(conn)
	for i, bankAcc := range bankAccs {
		bankAccID, err := bankAccRepo.Create(bankAcc)
		if err != nil {
			panic(err)
		}

		bankAccs[i].ID = bankAccID
	}

	return bankAccs[:]
}

func createTransactions(
	bankAcc models.BankAccount, conn *gorm.DB) []models.Transaction {
	transactions := [3]models.Transaction{
		{
			Type:          models.Income,
			Method:        models.Other,
			Credit:        false,
            Establishment: "company one",
			Value:         decimal.NewFromInt(10.000),
			Date:          time.Now(),
			CardID:        nil,
			BankAccountID: bankAcc.ID,
		},
		{
			Type:          models.Expense,
			Method:        models.DebitCard,
			Credit:        false,
            Establishment: "store two",
			Value:         decimal.NewFromInt(10.000),
			Date:          time.Now(),
			CardID:        nil,
			BankAccountID: bankAcc.ID,
		},
		{
			Type:          models.Expense,
			Method:        models.CreditCard,
            Establishment: "market",
			Credit:        true,
			Value:         decimal.NewFromInt(5.000),
			Date:          time.Now(),
			CardID:        nil,
			BankAccountID: bankAcc.ID,
		}}

	transactionRepo := NewTransactionRepository(conn)
	for i, transaction := range transactions {
		ID, err := transactionRepo.Create(&transaction)
		if err != nil {
			panic(err)
		}

		transactions[i].ID = ID
	}

	return transactions[:]
}

func createItems(transaction models.Transaction, conn *gorm.DB) []models.Item {
	items := [2]models.Item{
		{
			TransactionID: transaction.ID,
			Name:          "item one",
			Value:         transaction.Value.Div(decimal.NewFromInt(2)),
			Quantity:      1,
		},
		{
			TransactionID: transaction.ID,
			Name:          "item two",
			Value:         transaction.Value.Div(decimal.NewFromInt(4)),
			Quantity:      2,
		}}

	itemRepo := NewItemRepository(conn)
	IDs, err := itemRepo.CreateMultiple(items[:])
	if err != nil {
		panic(err)
	}

	for index, itemID := range IDs {
		items[index].ID = itemID
	}

	return items[:]
}

func createBadges(item models.Item, conn *gorm.DB) []models.Badge {
	badges := [2]models.Badge{
		{
			Name: "badge one",
            Color: "136 60 60",
		},
		{
			Name: "badge two",
            Color: "270 60 60",
		}}

	badgeRepo := NewBadgeRepository(conn)
	for _, badge := range badges {
		ID, err := badgeRepo.Create(&badge)
		if err != nil {
			panic(err)
		}

		badgeRepo.LinkItemToBadge(item.ID, ID)

		badge.ID = ID
	}

	return badges[:]
}
