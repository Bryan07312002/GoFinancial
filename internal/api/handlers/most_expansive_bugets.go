package handlers

import (
	"financial/internal/services"

	"encoding/json"
	"net/http"
)

type MostExpansiveBudgetsFactory interface {
	CreateMostExpansiveBadges() services.MostExpansiveBadges
}

type MostExpansiveBadgesHandler struct {
	factory MostExpansiveBudgetsFactory
}

func NewMostExpansiveBadgesHandler(factory MostExpansiveBudgetsFactory) http.Handler {
	return &MostExpansiveBadgesHandler{factory}
}

func (m *MostExpansiveBadgesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID, err := extractUserId(r)
	if err != nil {
		writeError(err, w)
		return
	}

	service := m.factory.CreateMostExpansiveBadges()
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
