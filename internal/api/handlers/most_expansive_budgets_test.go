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

type mockMostExpansiveBadgesFactory struct {
	mock func() services.MostExpansiveBadges
}

func (m *mockMostExpansiveBadgesFactory) CreateMostExpansiveBadges() services.MostExpansiveBadges {
	if m.mock != nil {
		return m.mock()
	}
	return &mockMostExpansiveBadges{}
}

type mockMostExpansiveBadges struct {
	runFn func(userID uint) ([]models.BadgeWithValue, error)
}

func (m *mockMostExpansiveBadges) Run(userID uint) ([]models.BadgeWithValue, error) {
	return m.runFn(userID)
}

func TestMostExpansiveBadgesHandler_ServeHTTP(t *testing.T) {
	tests := []struct {
		name           string
		factorySetup   func() *mockMostExpansiveBadgesFactory
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "sucessfful request",
			factorySetup: func() *mockMostExpansiveBadgesFactory {
				return &mockMostExpansiveBadgesFactory{
					mock: func() services.MostExpansiveBadges {
						return &mockMostExpansiveBadges{
							runFn: func(userID uint) ([]models.BadgeWithValue, error) {
								return []models.BadgeWithValue{}, nil
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
			factorySetup: func() *mockMostExpansiveBadgesFactory {
				return &mockMostExpansiveBadgesFactory{
					mock: func() services.MostExpansiveBadges {
						return &mockMostExpansiveBadges{
							runFn: func(userID uint) ([]models.BadgeWithValue, error) {
								return []models.BadgeWithValue{}, errors.BadRequestError()
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
			handler := NewMostExpansiveBadgesHandler(factory)

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
