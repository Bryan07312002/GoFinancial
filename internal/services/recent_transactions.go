package services

import (
	"financial/internal/db"
	"financial/internal/models"
)

type PaginateTransactions struct {
	transactionRepo db.TransactionRepository
}

func NewPaginateTransactions(
	transactionRepo db.TransactionRepository) PaginateTransactions {
	return PaginateTransactions{transactionRepo}
}

func (p *PaginateTransactions) Run() db.PaginateResult[models.BankAccount] {
	panic("not implemented yet")
}
