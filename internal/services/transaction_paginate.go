package services

import (
	"financial/internal/db"
	"financial/internal/models"
)

type PaginateTransaction struct {
	transactionRepo db.TransactionRepository
}

func NewPaginateTransaction(
	transactionRepo db.TransactionRepository) PaginateTransaction {
	return PaginateTransaction{transactionRepo}
}

func (p *PaginateTransaction) Run(
	paginationOption db.PaginateOptions,
	userID uint,
) (db.PaginateResult[models.TransactionWithBadges], error) {
	return p.transactionRepo.
		PaginateTransactionWithBadgesFromUserID(paginationOption, userID)
}
