package services

import (
	"financial/internal/db"
	"financial/internal/models"
)

type RecentTransactions struct {
	transactionRepo db.TransactionRepository
}

func NewRecentTransactions(
	transactionRepo db.TransactionRepository) RecentTransactions {
	return RecentTransactions{transactionRepo}
}

func (p *RecentTransactions) Run(userId uint) ([]models.TransactionWithBadges, error) {
	return p.transactionRepo.GetRecentTransactions(userId)
}
