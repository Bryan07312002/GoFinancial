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

type mockUpdateBankAccountFactory struct {
	createUpdateBankAccount func() services.UpdateBankAccount
}

func (m *mockUpdateBankAccountFactory) CreateUpdateBankAccount() services.UpdateBankAccount {
	if m.createUpdateBankAccount != nil {
		return m.createUpdateBankAccount()
	}
	return &mockUpdateBankAccount{}
}

type mockUpdateBankAccount struct {
	runFn func(bankAccountID uint, dto services.UpdateBankAccountDto, userID uint) error
}

func (m *mockUpdateBankAccount) Run(
	bankAccountID uint,
	dto services.UpdateBankAccountDto,
	userID uint,
) error {
	if m.runFn != nil {
		return m.runFn(bankAccountID, dto, userID)
	}

	return nil
}

func TestUpdateBankAccountHandler(t *testing.T) {
	updatedName := "updated name"
	updatedDescription := "updated description"
	validJSON := `{"name":"` + updatedName + `", "description":"` + updatedDescription + `"}`

	tests := []struct {
		name           string
		requestBody    string
		badgeID        uint
		factorySetup   func() *mockUpdateBankAccountFactory
		expectedStatus int
		expectedBody   string
	}{
		{
			name:        "successful update",
			requestBody: validJSON,
			badgeID:     1,
			factorySetup: func() *mockUpdateBankAccountFactory {
				return &mockUpdateBankAccountFactory{
					createUpdateBankAccount: func() services.UpdateBankAccount {
						return &mockUpdateBankAccount{
							runFn: func(
								bankAccountID uint,
								dto services.UpdateBankAccountDto,
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
			requestBody: validJSON,
			badgeID:     1,
			factorySetup: func() *mockUpdateBankAccountFactory {
				return &mockUpdateBankAccountFactory{
					createUpdateBankAccount: func() services.UpdateBankAccount {
						return &mockUpdateBankAccount{
							runFn: func(
								bankAccountID uint,
								dto services.UpdateBankAccountDto,
								userID uint,
							) error {
								if *dto.Name != updatedName {
									t.Error(
										"expected: " + updatedName + " got: " + *dto.Name)
								}

								if *dto.Description != updatedDescription {
									t.Error(
										"expected: " + updatedDescription + " got: " + *dto.Description)
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
			factorySetup: func() *mockUpdateBankAccountFactory {
				return nil
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: string(errors.BadRequestError().
				WithDetails("unexpected EOF").ToJSON()),
		},
		{
			name:        "service error",
			badgeID:     1,
			requestBody: validJSON,
			factorySetup: func() *mockUpdateBankAccountFactory {
				return &mockUpdateBankAccountFactory{
					createUpdateBankAccount: func() services.UpdateBankAccount {
						return &mockUpdateBankAccount{
							runFn: func(
								bankAccountID uint,
								dto services.UpdateBankAccountDto,
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
			handler := NewUpdateBankAccountHandler(factory)

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
