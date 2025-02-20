package handlers

import (
	"financial/internal/api/router/middlewares"
	"financial/internal/db"
	"financial/internal/services"

	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

func CreateCreateTransaction(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transaction services.CreateTransaction

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
		if err := service.Run(transaction, userID); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusCreated)
	}
}
