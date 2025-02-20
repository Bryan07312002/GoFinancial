package db

import (
	"financial/internal/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction *models.Transaction) (uint, error)
	FindByID(id uint) (models.Transaction, error)
	Delete(id uint) error
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

type transactionRepository struct {
	db *gorm.DB
}

func ToTransactionTable(t models.Transaction) TransactionTable {
	method := t.Method.String()

	return TransactionTable{
		ID:            t.ID,
		Type:          t.Type.String(),
		Method:        &method,
		Credit:        t.Credit,
		Value:         t.Value,
		Date:          t.Date,
		CardID:        t.CardID,
		BankAccountID: t.BankAccountID,
	}
}

func ToTransaction(t TransactionTable) models.Transaction {
	var method models.PaymentMethod
	if t.Method != nil {
		method = models.PaymentMethod(*t.Method)
	}

	return models.Transaction{
		ID: t.ID,
		// FIXME: should check if type is valid?
		Type:          models.TransactionType(t.Type),
		Method:        method,
		Credit:        t.Credit,
		Value:         t.Value,
		Date:          t.Date,
		CardID:        t.CardID,
		BankAccountID: t.BankAccountID,
	}
}

func (c *transactionRepository) Create(transaction *models.Transaction) (uint, error) {
	transactionTableInstance := ToTransactionTable(*transaction)
	if err := c.db.Create(&transactionTableInstance).Error; err != nil {
		return 0, err
	}

	return transactionTableInstance.ID, nil
}

func (c *transactionRepository) FindByID(id uint) (models.Transaction, error) {
	var transactionTableInstance TransactionTable

	if err := c.db.First(&transactionTableInstance, id).Error; err != nil {
		return models.Transaction{}, err
	}

	transaction := ToTransaction(transactionTableInstance)
	return transaction, nil
}

func (b *transactionRepository) Delete(id uint) error {
	// Attempt to delete the bank account by ID
	result := b.db.Delete(&TransactionTable{}, id)

	// Handle database errors (e.g., connection issues)
	if result.Error != nil {
		return result.Error
	}

	// Check if no rows were affected (record not found)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	// Successfully deleted
	return nil
}
