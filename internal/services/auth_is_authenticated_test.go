package services

import (
	"financial/internal/sessions"
	"testing"
)

func generateIsAuthenticatedService() (
	IsAuthenticated,
	*AuthorizationRepositoryMock) {

	authRepo := &AuthorizationRepositoryMock{}

	return IsAuthenticated{authRepo: authRepo}, authRepo
}

func TestIsAuthenticatedService(t *testing.T) {
	t.Run("should return true if repo return true", func(t *testing.T) {
		service, authRepo := generateIsAuthenticatedService()

		authRepo.IsAuthenticatedFunc = func(token sessions.Token) (uint, bool) {
			return 0, true
		}

		_, isAuth := service.Run(sessions.Token(""))

		if !isAuth {
			t.Errorf("Should return true but got false")
		}
	})

	t.Run("should return false if repo return false", func(t *testing.T) {
		service, authRepo := generateIsAuthenticatedService()

		authRepo.IsAuthenticatedFunc = func(token sessions.Token) (uint, bool) {
			return 0, false
		}

		_, isAuth := service.Run(sessions.Token(""))

		if isAuth {
			t.Errorf("Should return false but got true")
		}
	})
}
