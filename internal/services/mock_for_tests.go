package services

import (
	"financial/internal/db"
	"financial/internal/models"
	"financial/internal/sessions"
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

	FindByIDFunc                       func(ID uint) (models.BankAccount, error)
	FindBankAccountByCardIDFunc        func(cardID uint) (models.BankAccount, error)
	FindBankAccountByTransactionIDFunc func(
		transactionID uint) (models.BankAccount, error)

	PaginateFromUserIDFunc func(
		paginteOpt db.PaginateOptions,
		userID uint,
	) (db.PaginateResult[models.BankAccount], error)
	DeleteFunc func(ID uint) error
}

func (b *BankAccountRepositoryMock) Create(
	bankAccount models.BankAccount) (uint, error) {
	return b.CreateFunc(bankAccount)
}

func (b *BankAccountRepositoryMock) FindByID(
	id uint) (models.BankAccount, error) {
	return b.FindByIDFunc(id)
}

func (b *BankAccountRepositoryMock) FindBankAccountByCardID(
	cardID uint) (models.BankAccount, error) {
	return b.FindBankAccountByCardIDFunc(cardID)
}

func (b *BankAccountRepositoryMock) FindBankAccountByTransactionID(
	transactionID uint) (models.BankAccount, error) {
	return b.FindBankAccountByTransactionIDFunc(transactionID)
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
	CreateFunc   func(transaction *models.Transaction) (uint, error)
	FindByIDFunc func(id uint) (models.Transaction, error)
	DeleteFunc   func(id uint) error
}

func (t *TransactionRepositoryMock) Delete(id uint) error {
	return t.DeleteFunc(id)
}

func (t *TransactionRepositoryMock) FindByID(id uint) (models.Transaction, error) {
	return t.FindByIDFunc(id)
}

func (t *TransactionRepositoryMock) Create(transaction *models.Transaction) (uint, error) {
	return t.CreateFunc(transaction)
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
