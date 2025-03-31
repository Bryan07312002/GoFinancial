package handlers

import (
	"financial/internal/api/router/middlewares"
	"financial/internal/services"

	"encoding/json"
	"net/http"
)

type CurrentBalanceFactory interface {
	CreateCurrentBalance() services.CurrentBalance
}
type CurrentBalanceHandler struct {
	factory CurrentBalanceFactory
}

func NewCurrentBalanceHandler(factory CurrentBalanceFactory) http.Handler {
	return &CurrentBalanceHandler{factory}
}

func (c *CurrentBalanceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middlewares.UserKey).(uint)
	if !ok {
		http.Error(w, "User not found in context", http.StatusInternalServerError)
		return
	}

	service := c.factory.CreateCurrentBalance()
	result, err := service.Run(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(
			w,
			"Failed to encode response",
			http.StatusInternalServerError)
	}
}
