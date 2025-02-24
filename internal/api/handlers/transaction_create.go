package handlers

import (
	"financial/internal/api/router/middlewares"
	"financial/internal/db"
	"financial/internal/services"

	"encoding/json"
	"net/http"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type CreateTransaction struct {
	Type          string          `json:"type"`
	Value         decimal.Decimal `json:"value"`
	BankAccountID uint            `json:"bank_account_id"`
	Date          *string         `json:"date"`
	CardID        *uint           `json:"card_id"`
	Credit        *bool           `json:"credit"`
	Method        *string         `json:"method"`
}

func CreateCreateTransaction(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transaction CreateTransaction

		if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userID, ok := r.Context().Value(middlewares.UserKey).(uint)
		if !ok {
			http.Error(w, "User not found in context", http.StatusInternalServerError)
			return
		}

		transactionRepo := db.NewTransactionRepository(con)
		bankAccountRepo := db.NewBankAccountRepository(con)
		service := services.NewCreateTransactionService(
			transactionRepo, bankAccountRepo,
		)
		if err := service.Run(services.CreateTransaction{
			Type:          transaction.Type,
			Value:         transaction.Value,
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
}
