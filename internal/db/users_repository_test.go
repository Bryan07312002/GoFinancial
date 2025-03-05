package db_test

import (
	"financial/internal/db"
	"financial/internal/models"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func generateTestUserRepository(t *testing.T) db.UserRepository {
	conn, err := gorm.Open(sqlite.Open("file::memory:?"),
		&gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					LogLevel: logger.Silent,
				},
			),
		},
	)

	if err != nil {
		t.Fatalf("failed to connect to in-memory database: %v", err)
	}

	// Migrate the schema for testing
	err = conn.AutoMigrate(
		&db.UserTable{}, // Make sure to include all tables in migration
	)

	if err != nil {
		t.Fatalf("failed to migrate database schema: %v", err)
	}

	return db.NewUserRepository(conn)
}

func TestUserRepository(t *testing.T) {
	t.Run("CreateUser", func(t *testing.T) {
		userRepo := generateTestUserRepository(t)

		user := models.User{
			Name:     "John Doe",
			Password: "password123",
		}

		userID, err := userRepo.Create(user)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if userID == 0 {
			t.Errorf("expected user ID to be greater than 0, got %v", userID)
		}
	})

	t.Run("FindUserById", func(t *testing.T) {
		userRepo := generateTestUserRepository(t)

		// Create a user
		user := models.User{
			Name:     "Alice",
			Password: "alice123",
		}
		userID, err := userRepo.Create(user)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		// Find the user by ID
		foundUser, err := userRepo.FindById(userID)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if foundUser.Name != user.Name {
			t.Errorf("expected name %v, got %v", user.Name, foundUser.Name)
		}
		if foundUser.Password != user.Password {
			t.Errorf("expected password %v, got %v", user.Password, foundUser.Password)
		}
	})

	t.Run("FindUserByName", func(t *testing.T) {
		userRepo := generateTestUserRepository(t)

		// Create a user
		user := models.User{
			Name:     "Charlie",
			Password: "charlie123",
		}
		_, err := userRepo.Create(user)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		// Find the user by name
		foundUser, err := userRepo.FindByName("Charlie")
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if foundUser.Name != "Charlie" {
			t.Errorf("expected name %v, got %v", "Charlie", foundUser.Name)
		}
	})
}
