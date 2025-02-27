package handlers

import (
	"financial/internal/db"
	"financial/internal/hash"
	"financial/internal/services"
	"financial/internal/sessions"

	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

type AuthResponse struct {
	Token string `json:"token"`
}

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateLoginHandler(con *gorm.DB, jwtKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var form LoginForm

		if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userRepo := db.NewUserRepository(con)
		hashRepo := hash.NewHashRepository()
		sessionsRepo := sessions.NewAuthenticationRepository(jwtKey)

		service := services.NewLoginService(userRepo, &sessionsRepo, hashRepo)

		token, err := service.Run(services.LoginForm{
            Name: form.Email,
            Password: form.Password,
        })
		if err != nil {
            ReturnError(err, w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(AuthResponse{
			Token: string(token),
		})
	}
}
