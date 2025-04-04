package handlers

import (
	"financial/internal/api/router/middlewares"
	"financial/internal/errors"
	"financial/internal/services"

	"bytes"
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

type mockUpdateBadgeFactory struct {
	createUpdateBadgeFn func() services.UpdateBadge
}

func (m *mockUpdateBadgeFactory) CreateUpdateBadge() services.UpdateBadge {
	if m.createUpdateBadgeFn != nil {
		return m.createUpdateBadgeFn()
	}
	return &mockUpdateBadge{}
}

// mockUpdateBadge is a test implementation of RegisterUser service
type mockUpdateBadge struct {
	runFn func(badgeID uint, dto services.UpdateBadgeDto, userID uint) error
}

func (m *mockUpdateBadge) Run(
	badgeID uint,
	dto services.UpdateBadgeDto,
	userID uint,
) error {
	if m.runFn != nil {
		return m.runFn(badgeID, dto, userID)
	}

	return nil
}

func TestUpdateBadgeHandler(t *testing.T) {
	validJSON := `{"name":"", "color":""}`

	tests := []struct {
		name           string
		requestBody    string
		badgeID        uint
		factorySetup   func() *mockUpdateBadgeFactory
		expectedStatus int
		expectedBody   string
	}{
		{
			name:        "successful update",
			requestBody: validJSON,
			badgeID:     1,
			factorySetup: func() *mockUpdateBadgeFactory {
				return &mockUpdateBadgeFactory{
					createUpdateBadgeFn: func() services.UpdateBadge {
						return &mockUpdateBadge{
							runFn: func(
								badgeID uint,
								dto services.UpdateBadgeDto,
								userID uint,
							) error {
								return nil
							},
						}
					},
				}
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
		{
			name:        "recive correct values in service",
			requestBody: `{"name":"test_name", "color":"test_color"}`,
			badgeID:     1,
			factorySetup: func() *mockUpdateBadgeFactory {
				return &mockUpdateBadgeFactory{
					createUpdateBadgeFn: func() services.UpdateBadge {
						return &mockUpdateBadge{
							runFn: func(
								badgeID uint,
								dto services.UpdateBadgeDto,
								userID uint,
							) error {
								if dto.Name == nil {
									t.Errorf("expect: " + "test_name" + " got: nil")
								}

								if dto.Color == nil {
									t.Errorf("expect: " + "test_color" + " got: nil")
								}
								return nil
							},
						}
					},
				}
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
		{
			name:        "invalid JSON",
			requestBody: `{ "test": 1`,
			badgeID:     1,
			factorySetup: func() *mockUpdateBadgeFactory {
				return &mockUpdateBadgeFactory{}
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: string(errors.BadRequestError().
				WithDetails("unexpected EOF").ToJSON()),
		},
		{
			name:        "service error",
			badgeID:     1,
			requestBody: validJSON,
			factorySetup: func() *mockUpdateBadgeFactory {
				return &mockUpdateBadgeFactory{
					createUpdateBadgeFn: func() services.UpdateBadge {
						return &mockUpdateBadge{
							runFn: func(
								badgeID uint,
								dto services.UpdateBadgeDto,
								userID uint,
							) error {
								return errors.ValidationError().
									WithDetails("just test detail")
							},
						}
					},
				}
			},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody: string(errors.ValidationError().
				WithDetails("just test detail").ToJSON()),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			factory := tt.factorySetup()
			handler := NewUpdateBadgeHandler(factory)

			url := "/budgets/" + strconv.FormatUint(uint64(tt.badgeID), 10)

			req, err := http.NewRequest(
				"PUT",
				url,
				bytes.NewBufferString(tt.requestBody),
			)

			req = mux.SetURLVars(req, map[string]string{
				"id": strings.TrimPrefix(url, "/budgets/"),
			})

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

			if body := rr.Body.String(); body != tt.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v",
					body, tt.expectedBody)
			}
		})
	}
}
