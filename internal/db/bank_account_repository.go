package db

import (
	"financial/internal/models"

	"gorm.io/gorm"
)

// Easier to test another items that depends on this if interface is exported
type BankAccountRepository interface {
	Create(bankAccount models.BankAccount) (uint, error)

	FindByID(ID, userID uint) (models.BankAccount, error)
	FindBankAccountByCardID(cardID uint) (models.BankAccount, error)
	FindBankAccountByTransactionID(
		transactionID uint) (models.BankAccount, error)

	PaginateFromUserID(
		paginteOpt PaginateOptions,
		userID uint,
	) (PaginateResult[models.BankAccount], error)
	Update(bankAccount models.BankAccount) error
	Delete(ID uint) error
}

func toBankAccountTable(bankAccount models.BankAccount) BankAccountTable {
	return BankAccountTable{
		ID:          bankAccount.ID,
		UserID:      bankAccount.UserID,
		Name:        bankAccount.Name,
		Description: bankAccount.Description,
	}
}

func toBankAccount(bankAccountTable BankAccountTable) models.BankAccount {
	return models.BankAccount{
		ID:          bankAccountTable.ID,
		UserID:      bankAccountTable.UserID,
		Name:        bankAccountTable.Name,
		Description: bankAccountTable.Description,
	}
}

type bankAccountRepository struct {
	db *gorm.DB
}

func NewBankAccountRepository(db *gorm.DB) BankAccountRepository {
	return &bankAccountRepository{db}
}

func (b *bankAccountRepository) Create(bankAccount models.BankAccount) (uint, error) {
	bankAccountTableInstance := toBankAccountTable(bankAccount)

	if err := b.db.Create(&bankAccountTableInstance).Error; err != nil {
		return 0, err
	}

	return bankAccountTableInstance.ID, nil
}

func (r *bankAccountRepository) FindByID(ID, userID uint) (models.BankAccount, error) {
	var banckAccountTableInstance BankAccountTable

	if err := r.db.
		Model(&BankAccountTable{}).
		Where("user_id=?", userID).
		First(&banckAccountTableInstance, ID).Error; err != nil {

		return models.BankAccount{}, err
	}

	return toBankAccount(banckAccountTableInstance), nil
}

func (b *bankAccountRepository) FindBankAccountByCardID(
	cardID uint) (models.BankAccount, error) {
	var cardTableInstance CardTable

	if err := b.db.First(&cardTableInstance, cardID).Error; err != nil {
		return models.BankAccount{}, err
	}

	var bankAccountTableInstance BankAccountTable
	if err := b.db.First(&bankAccountTableInstance, cardTableInstance.BankAccountID).Error; err != nil {
		return models.BankAccount{}, err
	}

	return toBankAccount(bankAccountTableInstance), nil
}

func (c *bankAccountRepository) FindBankAccountByTransactionID(
	transactionID uint) (models.BankAccount, error) {
	// Create a variable to hold the result
	var bankAccountTableInstance BankAccountTable

	// Query the database for the transaction and join with the BankAccountTable
	err := c.db.
		Raw(`
			SELECT bank_accounts.*
			FROM transactions t
			JOIN bank_accounts bank_accounts ON t.bank_account_id = bank_accounts.id
			WHERE t.id = ?`, transactionID).
		Scan(&bankAccountTableInstance).Error

	if err != nil {
		return models.BankAccount{}, err
	}

	return toBankAccount(bankAccountTableInstance), nil
}

func (b *bankAccountRepository) PaginateFromUserID(
	paginateOpt PaginateOptions,
	userID uint,
) (PaginateResult[models.BankAccount], error) {
	var totalRecords int64
	var dbAccounts []BankAccountTable

	b.db.Model(&BankAccountTable{}).Where("user_id = ?", userID).Count(&totalRecords)

	offset := (paginateOpt.Page - 1) * paginateOpt.Take

	query := b.db.Model(&BankAccountTable{}).
		Where("user_id = ?", userID).
		Limit(int(paginateOpt.Take)).
		Offset(int(offset))

	if paginateOpt.SortBy != "" {
		order := paginateOpt.SortBy
		if paginateOpt.SortDesc {
			order += " DESC"
		} else {
			order += " ASC"
		}
		query = query.Order(order)
	}

	query.Find(&dbAccounts)

	results := make([]models.BankAccount, len(dbAccounts))
	for i, acc := range dbAccounts {
		results[i] = toBankAccount(acc)
	}

	totalPages := totalRecords / int64(paginateOpt.Take)
	if totalRecords%int64(paginateOpt.Take) != 0 {
		totalPages++
	}

	return PaginateResult[models.BankAccount]{
		Data:        results,
		Total:       uint64(totalRecords),
		CurrentPage: paginateOpt.Page,
		PageSize:    paginateOpt.Take,
		TotalPages:  uint(totalPages),
	}, nil
}

func (b *bankAccountRepository) Update(bankAccount models.BankAccount) error {
	result := b.db.Save(bankAccount)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (b *bankAccountRepository) Delete(id uint) error {
	result := b.db.Delete(&BankAccountTable{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
