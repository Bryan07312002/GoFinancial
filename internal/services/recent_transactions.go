package services

import (
	"financial/internal/db"
	"financial/internal/models"
)

type PaginateTransactions struct {
	transactionRepo db.TransactionRepository
}

func NewRecentTransactions(
	transactionRepo db.TransactionRepository) PaginateTransactions {
	return PaginateTransactions{transactionRepo}
}

func (p *PaginateTransactions) Run(userId uint) ([]models.TransactionWithBadges, error) {
	return p.transactionRepo.GetRecentTransactions(userId)
}
