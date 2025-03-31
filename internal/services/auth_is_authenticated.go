package services

import (
	"financial/internal/sessions"
)

type IsAuthenticated interface {
	Run(token sessions.Token) (uint, bool)
}

type isAuthenticated struct {
	authRepo sessions.AuthenticationRepository
}

func NewIsAuthenticated(
	authRepo sessions.AuthenticationRepository) IsAuthenticated {
	return &isAuthenticated{authRepo}
}

func (i *isAuthenticated) Run(token sessions.Token) (uint, bool) {
	return i.authRepo.IsAuthenticated(token)
}
