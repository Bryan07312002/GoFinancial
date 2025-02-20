package services

import (
	"financial/internal/models"
	"testing"
)

func generateService() (
	RegisterUserService, *UserRepositoryMock, *HashRepositoryMock) {
	repo := &UserRepositoryMock{}
	hashRepo := &HashRepositoryMock{}

	return RegisterUserService{
		userRepo: repo, hashRepo: hashRepo}, repo, hashRepo
}

func TestRegisterUser(t *testing.T) {
	t.Run("should call create 1 time", func(t *testing.T) {
		service, r, h := generateService()

		createCalledTimes := 0
		r.CreateFunc = func(_ models.User) (uint, error) {
			createCalledTimes = createCalledTimes + 1
			return 0, nil
		}

		HashCalledTimes := 0
		h.HashFunc = func(string string) (string, error) {
			HashCalledTimes += 1

			return "", nil
		}

		service.Run(RegisterUser{
			Name:     "tester",
			Password: "testPassword",
		})

		if createCalledTimes != 1 {
			t.Errorf("Create Function should have being called %d times but"+
				" was called %d times", 1, createCalledTimes)
		}
	})

	t.Run(
		"Hash should be applied in User.Password before calling repoUser.Create",
		func(t *testing.T) {
			service, r, h := generateService()
			input := RegisterUser{
				Name:     "tester",
				Password: "testPassword",
			}

			hashedPassword := "hashed password"
			h.HashFunc = func(string string) (string, error) {
				return hashedPassword, nil
			}

			r.CreateFunc = func(user models.User) (uint, error) {
				if user.Password != hashedPassword {
					t.Errorf("expected: User.Password be '%s'"+
						" but got: %s", hashedPassword, user.Password)
				}
				return 0, nil
			}

			service.Run(input)
		})
}
