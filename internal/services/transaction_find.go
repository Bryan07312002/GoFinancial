package services

import (
	"financial/internal/db"
	"financial/internal/models"
)

type FindTransaction struct {
	transactionsRepo db.TransactionRepository
}

func NewFindTransaction(transactionsRepo db.TransactionRepository) FindTransaction {
	return FindTransaction{transactionsRepo}
}

func (f *FindTransaction) Run(id, userID uint) (models.TransactionWithDetails, error) {
	return f.transactionsRepo.FindByIDWithDetails(id, userID)
}
