package handlers

import (
	"financial/internal/api/router/middlewares"
	"financial/internal/services"

	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type FindTransactionFactory interface {
	CreateFindTransaction() services.FindTransaction
}

type FindTransactionHandler struct {
	factory FindTransactionFactory
}

func NewFindTransactionHandler(factory FindTransactionFactory) http.Handler {
	return &FindTransactionHandler{factory}
}

func (f *FindTransactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middlewares.UserKey).(uint)
	if !ok {
		http.Error(
			w,
			"User not found in context",
			http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	transactionID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	service := f.factory.CreateFindTransaction()
	result, err := service.Run(uint(transactionID), userID)
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
