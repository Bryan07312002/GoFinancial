package handlers

import (
	"financial/internal/api/router/middlewares"
	"financial/internal/services"

	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type DeleteTransactionFactory interface {
	CreateDeleteTransaction() services.DeleteTransaction
}

type DeleteTransactionHandler struct {
	factory DeleteTransactionFactory
}

func NewDeleteTransactionHandler(factory DeleteTransactionFactory) http.Handler {
	return &DeleteTransactionHandler{factory}
}

func (d *DeleteTransactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	TransactionID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	service := d.factory.CreateDeleteTransaction()
	if err := service.Run(uint(TransactionID), userID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
