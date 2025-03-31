package handlers

import (
	"financial/internal/api/router/middlewares"
	"financial/internal/services"

	"encoding/json"
	"net/http"
)

type RecentTransactionsFactory interface {
	CreateRecentTransactions() services.RecentTransactions
}

type RecentTransactionsHandler struct {
	factory RecentTransactionsFactory
}

func NewRecentTransactionsHandler(factory RecentTransactionsFactory) http.Handler {
	return &RecentTransactionsHandler{factory}
}

func (re *RecentTransactionsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middlewares.UserKey).(uint)
	if !ok {
		http.Error(
			w,
			"User not found in context",
			http.StatusInternalServerError)
		return
	}

	service := re.factory.CreateRecentTransactions()
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
