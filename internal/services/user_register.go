package services

import (
	"financial/internal/db"
	"financial/internal/hash"
	"financial/internal/models"
)

type RegisterUserService struct {
	userRepo db.UserRepository
	hashRepo hash.HashRepository
}

func NewRegisterUserService(
	userRepo db.UserRepository,
	hashRepo hash.HashRepository,
) RegisterUserService {
	return RegisterUserService{userRepo, hashRepo}
}

type RegisterUser struct {
	Name     string
	Password string
}

func (s *RegisterUserService) Run(input RegisterUser) error {
	hashedPassword, err := s.hashRepo.Hash(input.Password)

	if err != nil {
		return err
	}

	// id will be created when user is saved in database
	user := models.User{
		Name:     input.Name,
		Password: hashedPassword,
	}

	s.userRepo.Create(user)

	return nil
}
