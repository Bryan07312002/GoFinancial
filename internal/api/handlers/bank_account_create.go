package handlers

import (
	"financial/internal/api/router/middlewares"
	"financial/internal/db"
	"financial/internal/services"

	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

type CreateBankAccount struct {
	Name        string
	Description string
}

func CreateCreateBankAccountHandler(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var form services.CreateBankAccount

		if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userID, ok := r.Context().Value(middlewares.UserKey).(uint)
		if !ok {
			http.Error(w, "User not found in context", http.StatusInternalServerError)
			return
		}

		bankAccountRepo := db.NewBankAccountRepository(con)
		service := services.NewCreateBankAccountService(bankAccountRepo)

		if err := service.Run(services.CreateBankAccount{
			UserId:      userID,
			Name:        form.Name,
			Description: form.Description,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
