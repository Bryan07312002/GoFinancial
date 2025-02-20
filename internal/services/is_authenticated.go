package services

import (
	"financial/internal/sessions"
)

type IsAuthenticated struct {
	authRepo sessions.AuthenticationRepository
}

func NewIsAuthenticatedService(
	authRepo sessions.AuthenticationRepository) IsAuthenticated {
	return IsAuthenticated{authRepo}
}

func (i *IsAuthenticated) Run(token sessions.Token) (uint, bool) {
	return i.authRepo.IsAuthenticated(token)
}
