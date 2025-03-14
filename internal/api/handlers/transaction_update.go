package handlers

import (
	"financial/internal/api/router/middlewares"
	"financial/internal/db"
	"financial/internal/services"
	"financial/internal/utils"
	"strconv"
	"time"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type UpdateTransactionDto struct {
	Type          *string  `json:"type,omitempty"`
	Method        *string  `json:"method,omitempty"`
	Credit        *bool    `json:"credit,omitempty"`
	Establishment *string  `json:"establishment,omitempty"`
	Value         *float32 `json:"value,omitempty"`
	Date          *string  `json:"date,omitempty"`
	CardID        *uint    `json:"card_id,omitempty"`
}

func CreateUpdateTransaction(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var form UpdateTransactionDto

		if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userID, ok := r.Context().Value(middlewares.UserKey).(uint)
		if !ok {
			http.Error(w, "User not found in context", http.StatusInternalServerError)
			return
		}

		vars := mux.Vars(r)
		strID := vars["id"]
		id, err := strconv.ParseUint(strID, 10, 0)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var date *time.Time
		if form.Date != nil {
			d, err := utils.ParseTime(*form.Date)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			date = &d
		}

        var value *decimal.Decimal
        if form.Value != nil {
            res := decimal.NewFromFloat32(*form.Value)
            value = &res
        }

		dto := services.UpdateTransactionDto{
			Type: form.Type,
			Date: date,
            Method: form.Method,
            Value: value,
            Credit: form.Credit,
            Establishment: form.Establishment,
            CardID: form.CardID,
		}

		transactionRepo := db.NewTransactionRepository(con)
		service := services.NewUpdateTransaction(transactionRepo)

		if err := service.Run(uint(id), dto, userID); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
