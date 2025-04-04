package handlers

import (
	"financial/internal/services"

	"net/http"
)

type NewItemsDto struct {
	Items         []services.NewItem `json:"items" validate:"required"`
	TransactionId uint               `json:"transaction_id" validate:"required"`
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

	userID, err := extractBodyAndUserId(r, &form)
	if err != nil {
		writeError(err, w)
		return
	}

	service := a.factory.CreateAddItemsToTransaction()
	if err := service.Run(form.Items, form.TransactionId, userID); err != nil {
		writeError(err, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
