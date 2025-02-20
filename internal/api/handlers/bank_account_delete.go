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

func CreateBankAccountDelete(con *gorm.DB) http.HandlerFunc {
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
		bankAccountID, err := strconv.ParseUint(id, 10, 0)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		bankAccountRepo := db.NewBankAccountRepository(con)
		service := services.NewDeleteBankAccountService(bankAccountRepo)

		if err := service.Run(uint(bankAccountID), userID); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
