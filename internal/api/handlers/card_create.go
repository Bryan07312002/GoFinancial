package handlers

import (
	"financial/internal/db"
	"financial/internal/services"

	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

func CreateCreateCard(dbCon *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var form services.CreateCard

		if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		cardRepo := db.NewCardRepository(dbCon)
		service := services.NewCreateCardService(cardRepo)

		if err := service.Run(form); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
