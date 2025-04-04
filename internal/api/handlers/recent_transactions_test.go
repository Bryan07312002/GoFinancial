package handlers

import (
	"financial/internal/api/router/middlewares"
	"financial/internal/errors"
	"financial/internal/models"
	"financial/internal/services"

	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockRecentTransactionsFactory struct {
	mock func() services.RecentTransactions
}

func (m *mockRecentTransactionsFactory) CreateRecentTransactions() services.RecentTransactions {
	if m.mock != nil {
		return m.mock()
	}
	return &mockRecentTransactions{}
}

type mockRecentTransactions struct {
	runFn func(userID uint) ([]models.TransactionWithBadges, error)
}

func (m *mockRecentTransactions) Run(userID uint) ([]models.TransactionWithBadges, error) {
	return m.runFn(userID)
}

func TestRecentTransactions_ServeHTTP(t *testing.T) {
	tests := []struct {
		name           string
		factorySetup   func() *mockRecentTransactionsFactory
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "sucessfful request",
			factorySetup: func() *mockRecentTransactionsFactory {
				return &mockRecentTransactionsFactory{
					mock: func() services.RecentTransactions {
						return &mockRecentTransactions{
							runFn: func(userID uint) ([]models.TransactionWithBadges, error) {
								return []models.TransactionWithBadges{}, nil
							},
						}
					},
				}
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "[]\n",
		},
		{
			name: "service error",
			factorySetup: func() *mockRecentTransactionsFactory {
				return &mockRecentTransactionsFactory{
					mock: func() services.RecentTransactions {
						return &mockRecentTransactions{
							runFn: func(userID uint) ([]models.TransactionWithBadges, error) {
								return []models.TransactionWithBadges{}, errors.BadRequestError()
							},
						}
					},
				}
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   string(errors.BadRequestError().ToJSON()),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			factory := tt.factorySetup()
			handler := NewRecentTransactionsHandler(factory)

			req, err := http.NewRequest(
				"GET",
				"/recent_transactions",
				bytes.NewBufferString(""),
			)

			if err != nil {
				t.Fatal(err)
			}

			req = req.WithContext(
				context.WithValue(req.Context(),
					middlewares.UserKey,
					uint(1),
				))
			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			body := rr.Body.String()
			if body != tt.expectedBody {
				t.Errorf("handler returned unexpected body:got '%v' want '%v'",
					body, tt.expectedBody)
			}
		})
	}
}
