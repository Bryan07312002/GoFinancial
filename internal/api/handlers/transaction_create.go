package handlers

import (
	"financial/internal/services"

	"net/http"
	"github.com/shopspring/decimal"
)


type CreateTransactionFactory interface {
	CreateCreateTransaction() services.CreateTransaction
}

type CreateTransactionRequest struct {
	Type          string          `json:"type" validate:"required"`
	Value         decimal.Decimal `json:"value" validate:"required"`
	BankAccountID uint            `json:"bank_account_id" validate:"required"`
	Establishment string          `json:"establishment" validate:"required"`
	Date          *string         `json:"date"`
	CardID        *uint           `json:"card_id"`
	Credit        *bool           `json:"credit"`
	Method        *string         `json:"method"`
}

type CreateTransactionHandler struct {
	factory CreateTransactionFactory
}

func NewCreateTransactionHandler(factory CreateTransactionFactory) http.Handler {
	return &CreateTransactionHandler{factory}
}

func (c *CreateTransactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var transaction CreateTransactionRequest

	userID, err := extractBodyAndUserId(r, &transaction)
	if err != nil {
		writeError(err, w)
		return
	}

	service := c.factory.CreateCreateTransaction()
	if err := service.Run(services.CreateTransactionDto{
		Type:          transaction.Type,
		Value:         transaction.Value,
		Establishment: transaction.Establishment,
		BankAccountID: transaction.BankAccountID,
		Date:          transaction.Date,
		CardID:        transaction.CardID,
		Credit:        transaction.Credit,
		Method:        transaction.Method,
	}, userID); err != nil {
		writeError(err, w)
        return
	}

	w.WriteHeader(http.StatusCreated)
}
