package handlers

import (
	"financial/internal/services"

	"encoding/json"
	"net/http"
)

type RegisterUserFactory interface {
	CreateRegisterUser() services.RegisterUser
}

type RegisterUserHandler struct {
	factory RegisterUserFactory
}

func NewRegisterUserHandler(factory RegisterUserFactory) http.Handler {
	return &RegisterUserHandler{factory}
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (re *RegisterUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var form RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	service := re.factory.CreateRegisterUser()

	// FIXME: handle errors correctly
	if err := service.Run(services.RegisterUserDto{
		Name:     form.Email,
		Password: form.Password,
	}); err != nil {
		writeError(err, w)
	}

	w.WriteHeader(http.StatusCreated)
}
