package handlers

import (
	"financial/internal/api/router/middlewares"
	"financial/internal/services"

	"encoding/json"
	"net/http"
)

type CreateBankAccount struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateBankAccountFactory interface {
	CreateCreateBankAccount() services.CreateBankAccount
}

type CreateBankAccountHandler struct {
	factory CreateBankAccountFactory
}

func NewCreateBankAccountHandler(factory CreateBankAccountFactory) http.Handler {
	return &CreateBankAccountHandler{factory}
}

func (c *CreateBankAccountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var form services.CreateBankAccountDto

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value(middlewares.UserKey).(uint)
	if !ok {
		http.Error(w, "User not found in context", http.StatusInternalServerError)
		return
	}

	service := c.factory.CreateCreateBankAccount()

	if err := service.Run(services.CreateBankAccountDto{
		UserId:      userID,
		Name:        form.Name,
		Description: form.Description,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
