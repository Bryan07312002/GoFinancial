package handlers

import (
	"financial/internal/services"

	"encoding/json"
	"net/http"
)

type CreateCardFactory interface {
	CreateCreateCard() services.CreateCard
}

type CreateCardHandler struct {
	factory CreateCardFactory
}

func NewCreateCardHandler(factory CreateCardFactory) http.Handler {
	return &CreateCardHandler{factory}
}

func (c *CreateCardHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var form services.CreateCardDto

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	service := c.factory.CreateCreateCard()

	if err := service.Run(form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
