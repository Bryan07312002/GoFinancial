package services

import (
	"financial/internal/db"
	"financial/internal/models"
	"financial/internal/sessions"
	"testing"
)

func generateLoginService() (
	Login,
	*UserRepositoryMock,
	*AuthorizationRepositoryMock,
	*HashRepositoryMock,
) {
	userRepo := &UserRepositoryMock{}
	authRepo := &AuthorizationRepositoryMock{}
	hashRepo := &HashRepositoryMock{}

	return NewLogin(userRepo, authRepo, hashRepo), userRepo, authRepo, hashRepo
}

func TestLogin(t *testing.T) {
	t.Run("Should login correctly", func(t *testing.T) {
		service, userRepo, authRepo, hashRepo := generateLoginService()

		userRepo.FindByNameFunc = func(name string) (models.User, error) {
			return models.User{
				ID:       0,
				Name:     "test",
				Password: "teste",
			}, nil
		}

		hashRepo.CompareFunc = func(s1, s2 string) bool { return true }

		authRepo.CreateTokenFunc = func(user models.User) (sessions.Token, error) {
			return "token", nil
		}

		result, err := service.Run(
			LoginForm{
				Name:     "test",
				Password: "teste",
			},
		)

		if result != "token" {
			t.Errorf("expected: %s \n but got: %s", "token", result)
		}

		if err != nil {
			t.Errorf("expected error to be: nil \n but got: %s", err.Error())
		}
	})

	t.Run("Should not login if user not found", func(t *testing.T) {
		service, userRepo, _, hashRepo := generateLoginService()

		userRepo.FindByNameFunc = func(name string) (models.User, error) {
			return models.User{
				ID:       0,
				Name:     "test",
				Password: "teste",
			}, nil
		}

		hashRepo.CompareFunc = func(s1, s2 string) bool { return false }

		res, err := service.Run(
			LoginForm{
				Name:     "test",
				Password: "teste",
			},
		)

		if res != "" {
			t.Errorf("expect empty string but got: %s", res)
		}

		if err == nil {
			t.Errorf("expect empty string but got nil value")
		}
	})

	t.Run("Should not login if passwords dont match", func(t *testing.T) {
		service, userRepo, _, _ := generateLoginService()

		userRepo.FindByNameFunc = func(name string) (models.User, error) {
			return models.User{}, db.ErrRecordNotFound
		}

		res, err := service.Run(
			LoginForm{
				Name:     "test",
				Password: "teste",
			},
		)

		if res != "" {
			t.Errorf("expect empty string but got: %s", res)
		}

		if err == nil {
			t.Errorf("expect empty string but got nil value")
		}
	})
}
