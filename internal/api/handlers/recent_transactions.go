package handlers

import (
	"financial/internal/api/router/middlewares"
	"financial/internal/db"
	"financial/internal/services"

	"gorm.io/gorm"
	"net/http"
)

func CreateRecentTransactions(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(middlewares.UserKey).(uint)
		if !ok {
			http.Error(
				w,
				"User not found in context",
				http.StatusInternalServerError)
			return
		}

		bankAccountRepo := db.NewTransactionRepository(con)
		service := services.NewRecentTransactions(bankAccountRepo)

		service.Run(userID)

	}
}
