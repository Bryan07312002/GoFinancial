package hash

import (
	"golang.org/x/crypto/bcrypt"
)

type HashRepository interface {
	Hash(string string) (string, error)
	Compare(input1, input2 string) bool
}

type hashRepository struct {
	cost int
}

func NewHashRepository() HashRepository {
	return &hashRepository{
		cost: bcrypt.DefaultCost,
	}
}

func (h *hashRepository) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (*hashRepository) Compare(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
