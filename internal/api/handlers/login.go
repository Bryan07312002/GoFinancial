package handlers

import (
	"financial/internal/services"

	"encoding/json"
	"net/http"
)

type LoginFactory interface {
	CreateLogin() services.Login
}

type LoginHandler struct {
	factory LoginFactory
}

func NewLoginHandler(factory LoginFactory) http.Handler {
	return &LoginHandler{factory}
}

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var form LoginForm

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	service := (*l.factory).CreateLogin()
	token, err := service.Run(services.LoginForm{
		Name:     form.Email,
		Password: form.Password,
	})
	if err != nil {
		ReturnError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(LoginResponse{
		Token: string(token),
	})
}
