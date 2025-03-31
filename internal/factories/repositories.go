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
	dbCon *gorm.DB
}

func NewRepositoryFactory(dbCon *gorm.DB) RepositoryFactory {
	return &repositoryFactory{dbCon}
}

func (*repositoryFactory) CreateUserRepository() db.UserRepository               { return nil }
func (*repositoryFactory) CreateBankAccountRepository() db.BankAccountRepository { return nil }
func (*repositoryFactory) CreateTransactionRepository() db.TransactionRepository { return nil }
func (*repositoryFactory) CreateItemRepository() db.ItemRepository               { return nil }
func (*repositoryFactory) CreateBadgeRepository() db.BadgeRepository             { return nil }
func (*repositoryFactory) CreateCardRepository() db.CardRepository               { return nil }
func (*repositoryFactory) CreateHashRepository() hash.HashRepository             { return nil }
func (*repositoryFactory) CreateAuthenticationRepository() sessions.AuthenticationRepository {
	return nil
}
func (*repositoryFactory) CreateUuidStrategy() uuid.UuidStrategy { return nil }
