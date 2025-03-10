package handlers

import (
	"financial/internal/api/router/middlewares"
	"financial/internal/db"
	"financial/internal/services"

	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

type NewItems struct {
	Items         []services.NewItem `json:"items"`
	TransactionId uint               `json:"transaction_id"`
}

func CreateAddItemsToTransaction(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var form NewItems

		if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userID, ok := r.Context().Value(middlewares.UserKey).(uint)
		if !ok {
			http.Error(
				w,
				"User not found in context",
				http.StatusInternalServerError,
			)
			return
		}

		itemRepo := db.NewItemRepository(con)
		transactionRepo := db.NewTransactionRepository(con)
		badgeRepo := db.NewBadgeRepository(con)
		bankAccountRepo := db.NewBankAccountRepository(con)
		service := services.NewAddItemsToTransaction(
			itemRepo,
			badgeRepo,
			transactionRepo,
			bankAccountRepo,
		)

		if err := service.Run(
			form.Items,
			form.TransactionId,
			userID,
		); err != nil {
			http.Error(w, "error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
