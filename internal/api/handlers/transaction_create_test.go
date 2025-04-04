package handlers

import (
	"financial/internal/api/router/middlewares"
	"financial/internal/errors"
	"financial/internal/services"

	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockCreateTransactionFactory struct {
	createCreateTransactionFn func() services.CreateTransaction
}

func (m *mockCreateTransactionFactory) CreateCreateTransaction() services.CreateTransaction {
	if m.createCreateTransactionFn != nil {
		return m.createCreateTransactionFn()
	}
	return &mockCreateTransaction{}
}

// mockCreateTransaction is a test implementation of RegisterUser service
type mockCreateTransaction struct {
	runFn func(newTransaction services.CreateTransactionDto, userId uint) error
}

func (m *mockCreateTransaction) Run(
	newTransaction services.CreateTransactionDto,
	userId uint,
) error {
	if m.runFn != nil {
		return m.runFn(newTransaction, userId)
	}
	return nil
}

func TestCreateTransactionHandler(t *testing.T) {
	validJSON := `{
            "type": "expanse",
            "value": 52.04,
            "bank_account_id": 1,
            "establishment": "amazon",
            "credit": true,
            "method": "Credit"
        }`

	tests := []struct {
		name           string
		requestBody    string
		factorySetup   func() *mockCreateTransactionFactory
		expectedStatus int
		expectedBody   string
	}{
		{
			name:        "successful create",
			requestBody: validJSON,
			factorySetup: func() *mockCreateTransactionFactory {
				return &mockCreateTransactionFactory{
					createCreateTransactionFn: func() services.CreateTransaction {
						return &mockCreateTransaction{
							runFn: func(dto services.CreateTransactionDto, userId uint) error {
								return nil
							},
						}
					},
				}
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   "",
		},
		{
			name:        "invalid JSON",
			requestBody: `{ "test": 1`,
			factorySetup: func() *mockCreateTransactionFactory {
				return &mockCreateTransactionFactory{}
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: string(errors.BadRequestError().
				WithDetails("unexpected EOF").ToJSON()),
		},
		{
			name: "missing type",
			requestBody: `{
            "value": 52.04,
            "bank_account_id": 1,
            "establishment": "amazon",
            "credit": true,
            "method": "Credit"
            }`,
			factorySetup: func() *mockCreateTransactionFactory {
				return &mockCreateTransactionFactory{}
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: string(errors.BadRequestError().
				AddFieldError("type", errors.RequiredField("type").Message).
				ToJSON()),
		},
		{
			name: "missing bank account id",
			requestBody: `{
            "type": "expanse",
            "value": 52.04,
            "establishment": "amazon",
            "credit": true,
            "method": "Credit"
            }`,
			factorySetup: func() *mockCreateTransactionFactory {
				return &mockCreateTransactionFactory{}
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: string(errors.BadRequestError().
				AddFieldError("bank_account_id", errors.RequiredField("bank_account_id").Message).
				ToJSON()),
		},
		{
			name: "missing establishment",
			requestBody: `{
            "type": "expanse",
            "value": 52.04,
            "bank_account_id": 1,
            "credit": true,
            "method": "Credit"
        }`,
			factorySetup: func() *mockCreateTransactionFactory {
				return &mockCreateTransactionFactory{}
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: string(errors.BadRequestError().
				AddFieldError("establishment", errors.RequiredField("establishment").Message).
				ToJSON()),
		},
		{
			name:        "service error",
			requestBody: validJSON,
			factorySetup: func() *mockCreateTransactionFactory {
				return &mockCreateTransactionFactory{
					createCreateTransactionFn: func() services.CreateTransaction {
						return &mockCreateTransaction{
							runFn: func(
								newTransaction services.CreateTransactionDto,
								userId uint,
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
			// Setup
			factory := tt.factorySetup()
			handler := NewCreateTransactionHandler(factory)

			req, err := http.NewRequest(
				"POST",
				"/transactions",
				bytes.NewBufferString(tt.requestBody),
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

			// Execute
			handler.ServeHTTP(rr, req)

			// Verify
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
