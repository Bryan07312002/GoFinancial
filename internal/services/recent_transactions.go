package services

import (
	"financial/internal/db"
	"financial/internal/models"
)

type RecentTransactions interface {
	Run(userId uint) ([]models.TransactionWithBadges, error)
}

type recentTransactions struct {
	transactionRepo db.TransactionRepository
}

func NewRecentTransactions(
	transactionRepo db.TransactionRepository) RecentTransactions {
	return &recentTransactions{transactionRepo}
}

func (p *recentTransactions) Run(userId uint) ([]models.TransactionWithBadges, error) {
	return p.transactionRepo.GetRecentTransactions(userId)
}
