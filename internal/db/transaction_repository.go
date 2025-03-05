package db

import (
	"financial/internal/models"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction *models.Transaction) (uint, error)
	FindByID(id uint) (models.Transaction, error)
	FindByIDWithDetails(id, userID uint) (models.TransactionWithDetails, error)
	PaginateFromUserID(
		paginteOpt PaginateOptions,
		userID uint,
	) (PaginateResult[models.Transaction], error)
	GetRecentTransactions(userID uint) ([]models.TransactionWithBadges, error)
	GetCurrentBalances(userID uint) (decimal.Decimal, decimal.Decimal, error)
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
		Establishment: t.Establishment,
		Value:         t.Value,
		Date:          t.Date,
		CardID:        t.CardID,
		BankAccountID: t.BankAccountID,
	}
}

func toTransactionWithDetails(t TransactionTable) models.TransactionWithDetails {
	var items []models.ItemWithBadges
	for _, item := range t.Items {
		var badges []models.Badge
		for _, badge := range item.Badges {
			badges = append(badges, ToBadge(badge))
		}

		items = append(items, models.ItemWithBadges{
			Item: models.Item{
				ID:            item.ID,
				Name:          item.Name,
				TransactionID: item.TransactionID,
				Value:         item.Value,
				Quantity:      item.Quantity,
			},
			Badges: badges,
		})
	}

	return models.TransactionWithDetails{
		Transaction: ToTransaction(t),
		Items:       items,
		BankAccount: ToBankAccount(t.BankAccount),
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
		Establishment: t.Establishment,
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

func (c *transactionRepository) FindByIDWithDetails(id, userID uint) (models.TransactionWithDetails, error) {
	var transaction TransactionTable
	err := c.db.
		Preload("Items.Badges").
		Preload("BankAccount").
		Preload("Card").
		Joins("JOIN bank_accounts ON transactions.bank_account_id = bank_accounts.id").
		Where("transactions.id = ? AND bank_accounts.user_id = ?", id, userID).
		First(&transaction).Error

	if err != nil {
		return models.TransactionWithDetails{}, err
	}

	return toTransactionWithDetails(transaction), nil
}

func (b *transactionRepository) PaginateFromUserID(
	paginateOpt PaginateOptions,
	userID uint,
) (PaginateResult[models.Transaction], error) {
	panic("not implemented yet")
}

func (b *transactionRepository) GetRecentTransactions(userID uint) ([]models.TransactionWithBadges, error) {
	var transactions []TransactionTable

	subquery := b.db.Model(&BankAccountTable{}).Select("id").Where("user_id = ?", userID)

	err := b.db.
		Preload("BankAccount").
		Preload("Card").
		Preload("Items.Badges").
		Where("bank_account_id IN (?)", subquery).
		Order("date DESC").
		Limit(5).
		Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	transactionsWithBadges := []models.TransactionWithBadges{}
	for _, transaction := range transactions {
		transactionWithBadges := models.TransactionWithBadges{
			Transaction: ToTransaction(transaction),
		}

		for _, item := range transaction.Items {
			for _, badge := range item.Badges {
				transactionWithBadges.Badges = append(
					transactionWithBadges.Badges, models.Badge{
						Name: badge.Name,
						ID:   badge.ID,
					})
			}
		}

		transactionsWithBadges = append(
			transactionsWithBadges,
			transactionWithBadges,
		)
	}

	return transactionsWithBadges, nil
}

func (b *transactionRepository) GetCurrentBalances(
	userID uint) (decimal.Decimal, decimal.Decimal, error) {

	query := `
        SELECT COALESCE(SUM(
          CASE
            WHEN type='income' THEN value
            ELSE -value
          END), 0) as balance
        FROM transactions
        LEFT JOIN bank_accounts
        ON transactions.bank_account_id=bank_accounts.id
        LEFT JOIN users
        ON bank_accounts.user_id=users.id
        WHERE users.id=?
        AND transactions.credit=false;
    `

	creditBalance := `
        SELECT COALESCE(SUM(
          CASE
            WHEN type='income' THEN value
            ELSE -value
          END), 0) as balance
        FROM transactions
        LEFT JOIN bank_accounts
        ON transactions.bank_account_id=bank_accounts.id
        LEFT JOIN users
        ON bank_accounts.user_id=users.id
        WHERE users.id=?
        AND transactions.credit=true;
    `

	var balance float64
	if err := b.db.Raw(query, userID).Scan(&balance).Error; err != nil {
		return decimal.Decimal{}, decimal.Decimal{}, err
	}

	var credit float64
	if err := b.db.Raw(creditBalance, userID).Scan(&credit).Error; err != nil {
		return decimal.Decimal{}, decimal.Decimal{}, err
	}

	return decimal.NewFromFloat(balance), decimal.NewFromFloat(credit), nil
}

func (b *transactionRepository) Delete(id uint) error {
	result := b.db.Delete(&TransactionTable{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
