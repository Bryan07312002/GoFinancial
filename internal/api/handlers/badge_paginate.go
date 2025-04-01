package handlers

import (
	"financial/internal/api/router/middlewares"
	"financial/internal/services"

	"encoding/json"
	"net/http"
)

type PaginateBadgeFactory interface {
	CreatePaginateBadges() services.PaginateBadges
}

type PaginateBadgesHandler struct {
	factory PaginateBadgeFactory
}

func NewPaginateBadgesHandler(factory PaginateBadgeFactory) http.Handler {
	return &PaginateBadgesHandler{factory}
}

func (p *PaginateBadgesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middlewares.UserKey).(uint)
	if !ok {
		http.Error(
			w,
			"User not found in context",
			http.StatusInternalServerError)
		return
	}

	service := p.factory.CreatePaginateBadges()
	result, err := service.Run(extractPaginationOptions(r), userID)
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
