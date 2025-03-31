package handlers

import (
	"financial/internal/api/router/middlewares"
	"financial/internal/services"

	"encoding/json"
	"net/http"
)

type PaginateTransactionFactory interface {
	CreatePaginateTransaction() services.PaginateTransaction
}

type PaginateTransaction struct {
	factory PaginateTransactionFactory
}

func NewPaginateTransaction(factory PaginateTransactionFactory) http.Handler {
	return &PaginateTransaction{factory}
}

func (p *PaginateTransaction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middlewares.UserKey).(uint)
	if !ok {
		http.Error(
			w,
			"User not found in context",
			http.StatusInternalServerError)
		return
	}

	service := p.factory.CreatePaginateTransaction()
	result, err := service.Run(
		extractPaginateOptionsWithTimeWindowSearch(r),
		userID,
	)
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
