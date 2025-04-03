package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"financial/internal/services"
	"financial/internal/sessions"
)

type mockLoginService struct {
	runFunc func(services.LoginForm) (sessions.Token, error)
}

func (m *mockLoginService) Run(form services.LoginForm) (sessions.Token, error) {
	return m.runFunc(form)
}

type mockFactory struct {
	service *mockLoginService
}

func (m *mockFactory) CreateLogin() services.Login {
	return m.service
}

func TestLoginHandler_Success(t *testing.T) {
	expectedToken := "test_token"
	expectedEmail := "test@example.com"
	expectedPassword := "password123"

	mockSvc := &mockLoginService{
		runFunc: func(form services.LoginForm) (sessions.Token, error) {
			if form.Name != expectedEmail {
				t.Errorf("expected email %q, got %q", expectedEmail, form.Name)
			}
			if form.Password != expectedPassword {
				t.Errorf("expected password %q, got %q", expectedPassword, form.Password)
			}
			return sessions.Token(expectedToken), nil
		},
	}
	factory := &mockFactory{service: mockSvc}
	handler := NewLoginHandler(factory)

	reqBody := `{"email": "test@example.com", "password": "password123"}`
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(reqBody))
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected Content-Type application/json, got %q", contentType)
	}

	var loginResp LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if loginResp.Token != expectedToken {
		t.Errorf("expected token %q, got %q", expectedToken, loginResp.Token)
	}
}

func TestLoginHandler_InvalidJSON(t *testing.T) {
	mockSvc := &mockLoginService{
		runFunc: func(form services.LoginForm) (sessions.Token, error) {
			return "token", nil
		},
	}
	factory := &mockFactory{service: mockSvc}
	handler := NewLoginHandler(factory)

	reqBody := `{invalid json}`
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(reqBody))
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", resp.StatusCode)
	}
}

func TestLoginHandler_InvalidCredentials(t *testing.T) {
	expectedError := services.NameOrPasswordNotMatchError

	mockSvc := &mockLoginService{
		runFunc: func(form services.LoginForm) (sessions.Token, error) {
			return "", expectedError
		},
	}
	factory := &mockFactory{service: mockSvc}
	handler := NewLoginHandler(factory)

	reqBody := `{"email": "wrong@example.com", "password": "wrong"}`
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(reqBody))
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("expected status 401, got %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	bodyStr := string(bodyBytes)

	if errMsg := bodyStr; errMsg != string(expectedError.ToJSON()) {
		t.Errorf("expected error %q, got %q", expectedError.Error(), errMsg)
	}
}

func TestLoginHandler_UknownError(t *testing.T) {
	expectedError := errors.New("internal error")

	mockSvc := &mockLoginService{
		runFunc: func(form services.LoginForm) (sessions.Token, error) {
			return "", expectedError
		},
	}
	factory := &mockFactory{service: mockSvc}
	handler := NewLoginHandler(factory)

	reqBody := `{"email": "wrong@example.com", "password": "wrong"}`
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(reqBody))
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	bodyStr := string(bodyBytes)

	if errMsg := bodyStr; errMsg != expectedError.Error() {
		t.Errorf("expected error %q, got %q", expectedError.Error(), errMsg)
	}
}
