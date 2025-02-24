package services

import (
	"financial/internal/db"
)

type PaginateTransactions struct {
	transactionRepo db.TransactionRepository
}

func NewRecentTransactions(
	transactionRepo db.TransactionRepository) PaginateTransactions {
	return PaginateTransactions{transactionRepo}
}

func (p *PaginateTransactions) Run(userId uint) {
	p.transactionRepo.GetRecentTransactions(userId)
}
