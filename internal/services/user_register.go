package services

import (
	"financial/internal/db"
	"financial/internal/hash"
	"financial/internal/models"
)

type RegisterUser interface {
	Run(input RegisterUserDto) error
}

type registerUser struct {
	userRepo db.UserRepository
	hashRepo hash.HashRepository
}

func NewRegisterUser(
	userRepo db.UserRepository,
	hashRepo hash.HashRepository,
) RegisterUser {
	return &registerUser{userRepo, hashRepo}
}

type RegisterUserDto struct {
	Name     string
	Password string
}

func (s *registerUser) Run(input RegisterUserDto) error {
	hashedPassword, err := s.hashRepo.Hash(input.Password)
	if err != nil {
		return err
	}

	// id will be created when user is saved in database
	user := models.User{
		Name:     input.Name,
		Password: hashedPassword,
	}

	if _, err := s.userRepo.Create(user); err != nil {
		return err
	}
	return nil
}
