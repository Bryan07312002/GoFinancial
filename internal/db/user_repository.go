package db

import (
	"financial/internal/errors"
	"financial/internal/models"

	"gorm.io/gorm"
)

var (
	ErrDuplicateEmail = errors.New(map[string]string{
		"email": "email already registered",
	})
)

// Easier to test another items that depends on this if interface is exported
type UserRepository interface {
	Create(user models.User) (uint, error)
	FindById(id uint) (models.User, error)
	FindByName(name string) (models.User, error)
}

type userRepository struct {
	conn *gorm.DB
}

func toUserTable(user models.User) UserTable {
	return UserTable{
		ID:       user.ID,
		Name:     user.Name,
		Password: user.Password,
	}
}

func toUser(userTable UserTable) models.User {
	return models.User{
		ID:       userTable.ID,
		Name:     userTable.Name,
		Password: userTable.Password,
	}
}

func NewUserRepository(conn *gorm.DB) UserRepository {
	return &userRepository{conn}
}

func (r *userRepository) Create(user models.User) (uint, error) {
	userInstance := toUserTable(user)
	err := r.conn.Create(&userInstance).Error
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.name" {
			return 0, &ErrDuplicateEmail
		}

		return 0, err
	}

	return userInstance.ID, nil
}

func (r *userRepository) FindByName(name string) (models.User, error) {
	var userTable UserTable

	err := r.conn.Where("name = ?", name).First(&userTable).Error
	if err != nil {
		return models.User{}, err
	}

	user := toUser(userTable)
	return user, nil
}

func (r *userRepository) FindById(id uint) (models.User, error) {
	var userTableInstance UserTable
	r.conn.First(&userTableInstance, id)

	return toUser(userTableInstance), nil
}
