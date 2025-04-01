package handlers

import (
	"financial/internal/api/router/middlewares"
	"financial/internal/services"

	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type DeleteItemFacotory interface {
	CreateDeleteItem() services.DeleteItem
}

type DeleteItemHandler struct {
	factory DeleteItemFacotory
}

func NewDeleteItemHandler(factory DeleteItemFacotory) http.Handler {
	return &DeleteItemHandler{factory}
}

func (d *DeleteItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middlewares.UserKey).(uint)
	if !ok {
		http.Error(w, "User not found in context", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	itemID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	service := d.factory.CreateDeleteItem()
	if err := service.Run(uint(itemID), userID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
