package handlers

import (
	"financial/internal/api/router/middlewares"
	"financial/internal/db"
	"financial/internal/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func CreateTransactionDelete(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(middlewares.UserKey).(uint)
		if !ok {
			http.Error(
				w,
				"User not found in context",
				http.StatusInternalServerError)
			return
		}

		vars := mux.Vars(r)
		id := vars["id"]
		TransactionID, err := strconv.ParseUint(id, 10, 0)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		bankAccountRepo := db.NewBankAccountRepository(con)
		transactionRepo := db.NewTransactionRepository(con)
		service := services.NewDeleteTransaction(bankAccountRepo, transactionRepo)

		if err := service.Run(uint(TransactionID), userID); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
