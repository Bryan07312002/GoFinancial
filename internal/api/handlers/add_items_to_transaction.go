package handlers

import (
	"financial/internal/api/router/middlewares"
	"financial/internal/services"

	"encoding/json"
	"net/http"
)

type NewItemsDto struct {
	Items         []services.NewItem `json:"items"`
	TransactionId uint               `json:"transaction_id"`
}
type AddItemsToTransactionFactory interface {
	CreateAddItemsToTransaction() services.AddItemsToTransaction
}
type AddItemsToTransactionHandler struct {
	factory AddItemsToTransactionFactory
}

func NewAddItemsToTransactionHandler(factory AddItemsToTransactionFactory) http.Handler {
	return &AddItemsToTransactionHandler{factory}
}

func (a *AddItemsToTransactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var form NewItemsDto

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value(middlewares.UserKey).(uint)
	if !ok {
		http.Error(
			w,
			"User not found in context",
			http.StatusInternalServerError,
		)
		return
	}

	service := a.factory.CreateAddItemsToTransaction()
	if err := service.Run(form.Items, form.TransactionId, userID); err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
