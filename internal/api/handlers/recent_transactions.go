package handlers

import (
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
	userID, err := extractUserId(r)
	if err != nil {
		writeError(err, w)
		return
	}

	service := re.factory.CreateRecentTransactions()
	result, err := service.Run(userID)
	if err != nil {
		writeError(err, w)
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
