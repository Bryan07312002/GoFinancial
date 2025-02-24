package db

import (
	"financial/internal/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction *models.Transaction) (uint, error)
	FindByID(id uint) (models.Transaction, error)
	PaginateFromUserID(
		paginteOpt PaginateOptions,
		userID uint,
	) (PaginateResult[models.Transaction], error)
	GetRecentTransactions(userID uint) ([]models.TransactionWithBadges, error)
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

func (b *transactionRepository) PaginateFromUserID(
	paginateOpt PaginateOptions,
	userID uint,
) (PaginateResult[models.Transaction], error) {
	panic("")
}

func (b *transactionRepository) GetRecentTransactions(userID uint) ([]models.TransactionWithBadges, error) {
	var transactions []TransactionTable

	// Subquery to get the user's bank account IDs
	subquery := b.db.Model(&BankAccountTable{}).Select("id").Where("user_id = ?", userID)

	// Main query to fetch transactions, preload relationships, and apply conditions
	err := b.db.
		Preload("BankAccount").
		Preload("Card").
		Preload("Items.Badges"). // Preload items and their badges
		Where("bank_account_id IN (?)", subquery).
		Order("date DESC").
		Limit(5).
		Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	var transactionsWithBadges []models.TransactionWithBadges
	for _, transaction := range transactions {
		transactionWithBadges := models.TransactionWithBadges{
			Transaction: ToTransaction(transaction),
		}

		for _, badge := range transaction.Items {
			transactionWithBadges.Badges = append(
				transactionWithBadges.Badges, models.Badge{
					Name: badge.Name,
					ID:   badge.ID,
				})
		}

		transactionsWithBadges = append(transactionsWithBadges, transactionWithBadges)
	}

	return transactionsWithBadges, nil
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
