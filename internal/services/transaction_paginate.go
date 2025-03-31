package services

import (
	"financial/internal/db"
	"financial/internal/models"
)

type PaginateTransaction interface {
	Run(
		paginationOption db.PaginateOptionsWithTimeWindowSearch,
		userID uint,
	) (db.PaginateResult[models.TransactionWithDetails], error)
}

type paginateTransaction struct {
	transactionRepo db.TransactionRepository
}

func NewPaginateTransaction(
	transactionRepo db.TransactionRepository) PaginateTransaction {
	return &paginateTransaction{transactionRepo}
}

func (p *paginateTransaction) Run(
	paginationOption db.PaginateOptionsWithTimeWindowSearch,
	userID uint,
) (db.PaginateResult[models.TransactionWithDetails], error) {
	return p.transactionRepo.
		PaginateTransactionWithDetailsFromUserID(paginationOption, userID)
}
