package handlers

import (
	"financial/internal/api/router/middlewares"
	"financial/internal/db"
	"financial/internal/services"

	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

func CreatePaginateBadge(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(middlewares.UserKey).(uint)
		if !ok {
			http.Error(
				w,
				"User not found in context",
				http.StatusInternalServerError)
			return
		}

		badgeRepo := db.NewBadgeRepository(con)
		service := services.NewPaginateBadges(badgeRepo)

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
}
