package services

import (
	"financial/internal/db"
	"financial/internal/models"
)

type FindTransaction interface {
	Run(id, userID uint) (models.TransactionWithDetails, error)
}

type findTransaction struct {
	transactionsRepo db.TransactionRepository
}

func NewFindTransaction(transactionsRepo db.TransactionRepository) FindTransaction {
	return &findTransaction{transactionsRepo}
}

func (f *findTransaction) Run(id, userID uint) (models.TransactionWithDetails, error) {
	return f.transactionsRepo.FindByIDWithDetails(id, userID)
}
