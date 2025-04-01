package handlers

import (
	"financial/internal/api/router/middlewares"
	"financial/internal/services"

	"encoding/json"
	"net/http"
)

type CreateBadgeFactory interface {
	CreateCreateBadge() services.CreateBadge
}

type CreateBadgeHandler struct {
	factory CreateBadgeFactory
}

func NewCreateBadgeHandler(factory CreateBadgeFactory) http.Handler {
	return &CreateBadgeHandler{factory}
}

func (c *CreateBadgeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var form services.NewBadge

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value(middlewares.UserKey).(uint)
	if !ok {
		http.Error(w, "User not found in context", http.StatusInternalServerError)
		return
	}

	service := c.factory.CreateCreateBadge()
	if err := service.Run(form, userID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
