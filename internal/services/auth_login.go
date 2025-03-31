package services

import (
	"financial/internal/db"
	"financial/internal/hash"
	"financial/internal/sessions"

	"errors"
)

type Login interface {
	Run(form LoginForm) (sessions.Token, error)
}

type login struct {
	userRepo db.UserRepository
	authRepo sessions.AuthenticationRepository
	hashRepo hash.HashRepository
}

func NewLogin(
	userRepo db.UserRepository,
	authRepo sessions.AuthenticationRepository,
	hashRepo hash.HashRepository,
) Login {
	return &login{
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

func (l *login) Run(form LoginForm) (sessions.Token, error) {
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
