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

func CreateDeleteItem(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(middlewares.UserKey).(uint)
		if !ok {
			http.Error(w, "User not found in context", http.StatusInternalServerError)
			return
		}

		vars := mux.Vars(r)
		id := vars["id"]
		itemID, err := strconv.ParseUint(id, 10, 0)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		itemRepo := db.NewItemRepository(con)
		bankAccRepo := db.NewBankAccountRepository(con)
		service := services.NewDeleteItem(itemRepo, bankAccRepo)

		if err := service.Run(uint(itemID), userID); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
