package handlers

import (
	"financial/internal/api/router/middlewares"
	"financial/internal/services"

	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type DeleteBadgeFactory interface{
    CreateDeleteBadge() services.DeleteBadge
}

type DeleteBadgeHandler struct {
	factory DeleteBadgeFactory
}

func NewDeleteBadgeHandler(factory DeleteBadgeFactory) http.Handler {
	return &DeleteBadgeHandler{factory}
}

func (d *DeleteBadgeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	badgeId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	service := d.factory.CreateDeleteBadge()
	if err := service.Run(uint(badgeId), userID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
