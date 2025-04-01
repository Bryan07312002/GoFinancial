package factories

import (
	"financial/internal/db"
	"financial/internal/hash"
	"financial/internal/sessions"
	"financial/internal/uuid"

	"gorm.io/gorm"
)

type RepositoryFactory interface {
	// db
	CreateUserRepository() db.UserRepository
	CreateBankAccountRepository() db.BankAccountRepository
	CreateTransactionRepository() db.TransactionRepository
	CreateItemRepository() db.ItemRepository
	CreateBadgeRepository() db.BadgeRepository
	CreateCardRepository() db.CardRepository
	// hash
	CreateHashRepository() hash.HashRepository
	// session
	CreateAuthenticationRepository() sessions.AuthenticationRepository
	// uuid
	CreateUuidStrategy() uuid.UuidStrategy
}

type repositoryFactory struct {
	dbCon     *gorm.DB
	secretKey string
}

func NewRepositoryFactory(dbCon *gorm.DB, secretKey string) RepositoryFactory {
	return &repositoryFactory{dbCon, secretKey}
}

func (r *repositoryFactory) CreateUserRepository() db.UserRepository {
	return db.NewUserRepository(r.dbCon)
}

func (r *repositoryFactory) CreateBankAccountRepository() db.BankAccountRepository {
	return db.NewBankAccountRepository(r.dbCon)
}

func (r *repositoryFactory) CreateTransactionRepository() db.TransactionRepository {
	return db.NewTransactionRepository(r.dbCon)
}

func (r *repositoryFactory) CreateItemRepository() db.ItemRepository {
	return db.NewItemRepository(r.dbCon)
}

func (r *repositoryFactory) CreateBadgeRepository() db.BadgeRepository {
	return db.NewBadgeRepository(r.dbCon)
}

func (r *repositoryFactory) CreateCardRepository() db.CardRepository {
	return db.NewCardRepository(r.dbCon)
}

func (r *repositoryFactory) CreateHashRepository() hash.HashRepository {
	return hash.NewHashRepository()
}

func (r *repositoryFactory) CreateAuthenticationRepository() sessions.AuthenticationRepository {
	return sessions.NewAuthenticationRepository(r.secretKey)
}

func (r *repositoryFactory) CreateUuidStrategy() uuid.UuidStrategy {
    return uuid.NewUuidStrategy()
}
