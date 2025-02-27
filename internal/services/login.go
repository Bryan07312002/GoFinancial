package services

import (
	"financial/internal/db"
	"financial/internal/hash"
	"financial/internal/sessions"

	"errors"
)

type LoginService struct {
	userRepo db.UserRepository
	authRepo sessions.AuthenticationRepository
	hashRepo hash.HashRepository
}

func NewLoginService(
	userRepo db.UserRepository,
	authRepo sessions.AuthenticationRepository,
	hashRepo hash.HashRepository,
) LoginService {
	return LoginService{
		userRepo,
		authRepo,
		hashRepo,
	}
}

type LoginForm struct {
	Name     string
	Password string
}

var (
	EmailOrPasswordNotMatchError = errors.New("email or password dont match with any record")
)

func (l *LoginService) Run(form LoginForm) (sessions.Token, error) {
	user, err := l.userRepo.FindByName(form.Name)
	if err != nil {
		return sessions.Token(""), EmailOrPasswordNotMatchError
	}

	if !l.hashRepo.Compare(form.Password, user.Password) {
		// TODO: add a proper error return here
		return sessions.Token(""), EmailOrPasswordNotMatchError
	}

	return l.authRepo.CreateToken(user)
}
