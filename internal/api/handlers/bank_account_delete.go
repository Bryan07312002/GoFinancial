package handlers

import (
	"financial/internal/api/router/middlewares"
	"financial/internal/services"

	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type DeleteBankAccountFactory interface {
	CreateDeleteBankAccount() services.DeleteBankAccount
}

type DeleteBankAccountHandler struct {
	factory DeleteBankAccountFactory
}

func NewDeleteBankAccountHandler(factory DeleteBankAccountFactory) http.Handler {
	return &DeleteBankAccountHandler{factory}
}

func (d *DeleteBankAccountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	bankAccountID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	service := d.factory.CreateDeleteBankAccount()
	if err := service.Run(uint(bankAccountID), userID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
