package handlers

import (
	"financial/internal/api/router/middlewares"
	"financial/internal/db"
	"financial/internal/services"

	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CreateCardDelete(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(middlewares.UserKey).(uint)
		if !ok {
			http.Error(w, "User not found in context", http.StatusInternalServerError)
			return
		}

		vars := mux.Vars(r)
		id := vars["id"]
		cardID, err := strconv.ParseUint(id, 10, 0)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		cardRepo := db.NewCardRepository(con)
		bankAccountRepo := db.NewBankAccountRepository(con)
		service := services.NewDeleteCard(cardRepo, bankAccountRepo)

		if err := service.Run(uint(cardID), userID); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
