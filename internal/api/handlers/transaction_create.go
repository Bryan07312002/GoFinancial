package handlers

import (
	"financial/internal/api/router/middlewares"
	"financial/internal/services"

	"encoding/json"
	"github.com/shopspring/decimal"
	"net/http"
)

type CreateTransactionFactory interface {
	CreateCreateTransaction() services.CreateTransaction
}

type CreateTransactionRequest struct {
	Type          string          `json:"type"`
	Value         decimal.Decimal `json:"value"`
	BankAccountID uint            `json:"bank_account_id"`
	Establishment string          `json:"establishment"`
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

	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value(middlewares.UserKey).(uint)
	if !ok {
		http.Error(w, "User not found in context", http.StatusInternalServerError)
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
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
}
