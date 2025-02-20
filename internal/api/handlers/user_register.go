package handlers

import (
	"encoding/json"
	"financial/internal/db"
	"financial/internal/hash"
	"financial/internal/services"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateRegisterUserHandler(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var form RegisterRequest

		if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userRepo := db.NewUserRepository(con)
		hashRepo := hash.NewHashRepository()
		service := services.NewRegisterUserService(userRepo, hashRepo)

		// FIXME: handle errors correctly
		if err := service.Run(services.RegisterUser{
			Name:     form.Email,
			Password: form.Password,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusCreated)
	}
}
