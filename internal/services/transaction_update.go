package services

import (
	"financial/internal/db"
	"financial/internal/models"

	"github.com/shopspring/decimal"
	"time"
)

type UpdateTransaction struct {
	transactionRepo db.TransactionRepository
}

type UpdateTransactionDto struct {
	Type          *string          `json:"type,omitempty"`
	Method        *string          `json:"method,omitempty"`
	Credit        *bool            `json:"credit,omitempty"`
	Establishment *string          `json:"establishment,omitempty"`
	Value         *decimal.Decimal `json:"value,omitempty"`
	Date          *time.Time       `json:"date,omitempty"`
	CardID        *uint            `json:"card_id,omitempty"`
}

func NewUpdateTransaction(
	transactionRepo db.TransactionRepository) UpdateTransaction {
	return UpdateTransaction{transactionRepo}
}

func (u *UpdateTransaction) Run(
	transactionID uint,
	dto UpdateTransactionDto,
	userID uint,
) error {
	transaction, err := u.transactionRepo.FindByID(transactionID, userID)
	if err != nil {
		return err
	}

	updateTransactionFields(&transaction, &dto)

	return u.transactionRepo.Update(transaction)
}

func updateTransactionFields(transaction *models.Transaction, dto *UpdateTransactionDto) {
	if dto.Type != nil {
		transaction.Type = models.TransactionType(*dto.Type)
	}

	if dto.Method != nil {
		transaction.Method = models.PaymentMethod(*dto.Method)
	}

	if dto.Credit != nil {
		transaction.Credit = *dto.Credit
	}

	if dto.Establishment != nil {
		transaction.Establishment = *dto.Establishment
	}

	if dto.Value != nil {
		transaction.Value = *dto.Value
	}

	if dto.Date != nil {
		transaction.Date = *dto.Date
	}

	if dto.CardID != nil {
		transaction.CardID = dto.CardID
	}
}
