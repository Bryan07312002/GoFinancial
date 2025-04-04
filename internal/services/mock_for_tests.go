package services

import (
	"financial/internal/db"
	"financial/internal/models"
	"financial/internal/sessions"

	"github.com/shopspring/decimal"
)

type HashRepositoryMock struct {
	HashFunc    func(string string) (string, error)
	CompareFunc func(s1, s2 string) bool
}

func (m *HashRepositoryMock) Hash(string string) (string, error) {
	return m.HashFunc(string)
}

func (m *HashRepositoryMock) Compare(s1, s2 string) bool {
	return m.CompareFunc(s1, s2)
}

// -------------------------------
//
//	DB repositories
//
// -------------------------------

type UserRepositoryMock struct {
	CreateFunc     func(user models.User) (uint, error)
	FindByIdFunc   func(id uint) (models.User, error)
	FindByNameFunc func(name string) (models.User, error)
	DeleteFunc     func(id uint) error
}

func (m *UserRepositoryMock) Create(user models.User) (uint, error) {
	return m.CreateFunc(user)
}

func (m *UserRepositoryMock) FindById(id uint) (models.User, error) {
	return m.FindByIdFunc(id)
}

func (m *UserRepositoryMock) FindByName(name string) (models.User, error) {
	return m.FindByNameFunc(name)
}

func (m *UserRepositoryMock) Delete(id uint) error {
	return m.DeleteFunc(id)
}

type BankAccountRepositoryMock struct {
	CreateFunc func(bankAccount models.BankAccount) (uint, error)

	FindByIDFunc                       func(ID, userId uint) (models.BankAccount, error)
	FindBankAccountByCardIDFunc        func(cardID uint) (models.BankAccount, error)
	FindBankAccountByTransactionIDFunc func(
		transactionID uint) (models.BankAccount, error)

	PaginateFromUserIDFunc func(
		paginteOpt db.PaginateOptions,
		userID uint,
	) (db.PaginateResult[models.BankAccount], error)
	UpdateFunc func(bankAccount models.BankAccount) error
	DeleteFunc func(ID uint) error
}

func (b *BankAccountRepositoryMock) Create(
	bankAccount models.BankAccount) (uint, error) {
	return b.CreateFunc(bankAccount)
}

func (b *BankAccountRepositoryMock) FindByID(
	id, usedID uint) (models.BankAccount, error) {
	return b.FindByIDFunc(id, usedID)
}

func (b *BankAccountRepositoryMock) FindBankAccountByCardID(
	cardID uint) (models.BankAccount, error) {
	return b.FindBankAccountByCardIDFunc(cardID)
}

func (b *BankAccountRepositoryMock) FindBankAccountByTransactionID(
	transactionID uint) (models.BankAccount, error) {
	return b.FindBankAccountByTransactionIDFunc(transactionID)
}

func (b *BankAccountRepositoryMock) Update(
	bankAccount models.BankAccount) error {
	return b.UpdateFunc(bankAccount)
}

func (b *BankAccountRepositoryMock) PaginateFromUserID(
	paginteOpt db.PaginateOptions,
	userID uint,
) (db.PaginateResult[models.BankAccount], error) {
	return b.PaginateFromUserID(paginteOpt, userID)
}

func (b *BankAccountRepositoryMock) Delete(id uint) error {
	return b.DeleteFunc(id)
}

type CardRepositoryMock struct {
	CreateFunc   func(card *models.Card) (uint, error)
	FindByIDFunc func(id uint) (*models.Card, error)
	DeleteFunc   func(id uint) error
}

func (c *CardRepositoryMock) FindByID(id uint) (*models.Card, error) {
	return c.FindByIDFunc(id)
}

func (c *CardRepositoryMock) Create(card *models.Card) (uint, error) {
	return c.CreateFunc(card)
}

func (c *CardRepositoryMock) Delete(id uint) error {
	return c.DeleteFunc(id)
}

type TransactionRepositoryMock struct {
	CreateFunc                                   func(transaction *models.Transaction) (uint, error)
	FindByIDFunc                                 func(id, userID uint) (models.Transaction, error)
	FindByIDWithDetailsFunc                      func(id, userID uint) (models.TransactionWithDetails, error)
	PaginateTransactionWithDetailsFromUserIDFunc func(
		paginteOpt db.PaginateOptionsWithTimeWindowSearch,
		userID uint,
	) (db.PaginateResult[models.TransactionWithDetails], error)
	GetRecentTransactionsFunc func(userID uint) ([]models.TransactionWithBadges, error)
	GetCurrentBalancesFunc    func(userID uint) (decimal.Decimal, decimal.Decimal, error)
	UpdateFunc                func(t models.Transaction) error
	DeleteFunc                func(id uint) error
}

func (t *TransactionRepositoryMock) Delete(id uint) error {
	return t.DeleteFunc(id)
}

func (t *TransactionRepositoryMock) Update(transaction models.Transaction) error {
	return t.UpdateFunc(transaction)
}

func (t *TransactionRepositoryMock) FindByIDWithDetails(
	id,
	userID uint,
) (models.TransactionWithDetails, error) {
	return t.FindByIDWithDetailsFunc(id, userID)
}

func (t *TransactionRepositoryMock) GetRecentTransactions(
	userID uint,
) ([]models.TransactionWithBadges, error) {
	return t.GetRecentTransactionsFunc(userID)
}

func (t *TransactionRepositoryMock) GetCurrentBalances(
	userID uint,
) (decimal.Decimal, decimal.Decimal, error) {
	return t.GetCurrentBalancesFunc(userID)
}

func (t *TransactionRepositoryMock) FindByID(id, userID uint) (models.Transaction, error) {
	return t.FindByIDFunc(id, userID)
}

func (t *TransactionRepositoryMock) PaginateTransactionWithDetailsFromUserID(
	paginateOpt db.PaginateOptionsWithTimeWindowSearch,
	userID uint,
) (db.PaginateResult[models.TransactionWithDetails], error) {
	return t.PaginateTransactionWithDetailsFromUserIDFunc(paginateOpt, userID)
}

func (t *TransactionRepositoryMock) Create(
	transaction *models.Transaction,
) (uint, error) {
	return t.CreateFunc(transaction)
}

type BadgeRepositoryMock struct {
	CreateFn             func(badge *models.Badge) (uint, error)
	CreateMultipleFn     func(badges []models.Badge) ([]uint, error)
	LinkItemToBadgeFn    func(itemID uint, badgeID uint) error
	FindByIDFn           func(id, userID uint) (models.Badge, error)
	FindByItemFn         func(itemID uint) ([]models.Badge, error)
	FindByTransactionFn  func(transactionID uint) ([]models.Badge, error)
	PaginateFromUserIDFn func(
		paginateOpt db.PaginateOptions,
		userID uint,
	) (db.PaginateResult[models.Badge], error)
	GetMostExpansivesFn func(userID uint) ([]models.BadgeWithValue, error)
	UpdateFn            func(badge models.Badge) error
	DeleteFn            func(id uint) error
}

func (b *BadgeRepositoryMock) Create(badge *models.Badge) (uint, error) {
	return b.CreateFn(badge)
}

func (b *BadgeRepositoryMock) CreateMultiple(badges []models.Badge) ([]uint, error) {
	return b.CreateMultipleFn(badges)
}

func (b *BadgeRepositoryMock) LinkItemToBadge(itemID uint, badgeID uint) error {
	return b.LinkItemToBadgeFn(itemID, badgeID)
}

func (b *BadgeRepositoryMock) FindByID(id, userID uint) (models.Badge, error) {
	return b.FindByIDFn(id, userID)
}

func (b *BadgeRepositoryMock) FindByItem(itemID uint) ([]models.Badge, error) {
	return b.FindByItem(itemID)
}

func (b *BadgeRepositoryMock) FindByTransaction(
	transactionID uint) ([]models.Badge, error) {
	return b.FindByTransaction(transactionID)
}

func (b *BadgeRepositoryMock) PaginateFromUserID(
	paginateOpt db.PaginateOptions,
	userID uint,
) (db.PaginateResult[models.Badge], error) {
	return b.PaginateFromUserIDFn(paginateOpt, userID)
}

func (b *BadgeRepositoryMock) GetMostExpansives(
	userID uint) ([]models.BadgeWithValue, error) {
	return b.GetMostExpansivesFn(userID)
}

func (b *BadgeRepositoryMock) Update(badge models.Badge) error {
	return b.UpdateFn(badge)
}

func (b *BadgeRepositoryMock) Delete(id uint) error {
	return b.DeleteFn(id)
}

// -------------------------------
//
//	SESSIONS repositories
//
// -------------------------------

type AuthorizationRepositoryMock struct {
	CreateTokenFunc     func(user models.User) (sessions.Token, error)
	IsAuthenticatedFunc func(token sessions.Token) (uint, bool)
}

func (a *AuthorizationRepositoryMock) CreateToken(user models.User) (
	sessions.Token, error) {
	return a.CreateTokenFunc(user)
}

func (a *AuthorizationRepositoryMock) IsAuthenticated(
	token sessions.Token) (uint, bool) {
	return a.IsAuthenticatedFunc(token)
}
