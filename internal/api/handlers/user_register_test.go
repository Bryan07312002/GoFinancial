package handlers

import (
	"bytes"
	"financial/internal/db"
	"financial/internal/errors"
	"financial/internal/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockRegisterUserFactory is a test implementation of RegisterUserFactory
type MockRegisterUserFactory struct {
	createRegisterUserFn func() services.RegisterUser
}

func (m *MockRegisterUserFactory) CreateRegisterUser() services.RegisterUser {
	if m.createRegisterUserFn != nil {
		return m.createRegisterUserFn()
	}
	return &MockRegisterUser{}
}

// MockRegisterUser is a test implementation of RegisterUser service
type MockRegisterUser struct {
	runFn func(dto services.RegisterUserDto) error
}

func (m *MockRegisterUser) Run(dto services.RegisterUserDto) error {
	if m.runFn != nil {
		return m.runFn(dto)
	}
	return nil
}

func TestRegisterUserHandler_ServeHTTP(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    string
		factorySetup   func() *MockRegisterUserFactory
		expectedStatus int
		expectedBody   string
	}{
		{
			name:        "successful registration",
			requestBody: `{"email": "test@example.com", "password": "secret"}`,
			factorySetup: func() *MockRegisterUserFactory {
				return &MockRegisterUserFactory{
					createRegisterUserFn: func() services.RegisterUser {
						return &MockRegisterUser{
							runFn: func(dto services.RegisterUserDto) error {
								if dto.Name != "test@example.com" || dto.Password != "secret" {
									t.Error("unexpected DTO values")
								}
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
			requestBody: `{"email": "test@example.com", "password": "secret"`,
			factorySetup: func() *MockRegisterUserFactory {
				return &MockRegisterUserFactory{}
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "unexpected EOF\n",
		},
		{
			name:        "missing email",
			requestBody: `{"password": "secret"}`,
			factorySetup: func() *MockRegisterUserFactory {
				return &MockRegisterUserFactory{}
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: string(errors.BadRequestError().
				AddFieldError("email", errors.RequiredField("email").Message).ToJSON()),
		},
		{
			name:        "service error",
			requestBody: `{"email": "test@example.com", "password": "secret"}`,
			factorySetup: func() *MockRegisterUserFactory {
				return &MockRegisterUserFactory{
					createRegisterUserFn: func() services.RegisterUser {
						return &MockRegisterUser{
							runFn: func(dto services.RegisterUserDto) error {
								return db.ErrDuplicateEmail
							},
						}
					},
				}
			},
			expectedStatus: http.StatusConflict,
			expectedBody: string(errors.ConflictError().
				AddFieldError("email", errors.ConflictField("email").Message).ToJSON()),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			factory := tt.factorySetup()
			handler := NewRegisterUserHandler(factory)

			req, err := http.NewRequest("POST", "/register", bytes.NewBufferString(tt.requestBody))
			if err != nil {
				t.Fatal(err)
			}

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
