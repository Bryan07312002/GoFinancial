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
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (re *RegisterUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var form RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validateRequest(form); err != nil {
		writeError(err, w)
		return
	}

	service := re.factory.CreateRegisterUser()

	if err := service.Run(services.RegisterUserDto{
		Name:     form.Email,
		Password: form.Password,
	}); err != nil {
		writeError(err, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
