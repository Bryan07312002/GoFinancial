package services

import (
	"errors"
	"financial/internal/models"
	"testing"
)

func generateBadgeUpdate() (
	UpdateBadge,
	*BadgeRepositoryMock,
) {
	badgeRepoMock := &BadgeRepositoryMock{}
	return NewUpdateBadge(badgeRepoMock), badgeRepoMock
}

func TestShouldUpdateExistingBadge(t *testing.T) {
	existingBadge := models.Badge{ID: 1, Name: "", Color: ""}

	newName := "a"
	newColor := "red"
	badgeUpdateDto := UpdateBadgeDto{Name: &newName, Color: &newColor}

	service, bageRepo := generateBadgeUpdate()

	FindByIDCalledTimes := 0
	bageRepo.FindByIDFn = func(id, userID uint) (models.Badge, error) {
		FindByIDCalledTimes += 1
		return existingBadge, nil
	}

	bageRepo.UpdateFn = func(badge models.Badge) error {
		if badge.Name != newName {
			t.Errorf("expected updated name to be: " +
				newName + " but got: " + badge.Name)
		}

		if badge.Color != newColor {
			t.Errorf(
				"expected updated color to be: " +
					newColor + " but got: " + badge.Color,
			)
		}

		return nil
	}

	if err := service.Run(1, badgeUpdateDto, 1); err != nil {
		t.Errorf("should not return error but got error: %s", err.Error())
	}

	if FindByIDCalledTimes != 1 {
		t.Errorf("FindByID should have beeing called 1 time but got: %d",
			FindByIDCalledTimes)
	}
}

func TestShouldReturnErrorIfFindByIDReturnsError(t *testing.T) {
	newName := "a"
	newColor := "red"
	badgeUpdateDto := UpdateBadgeDto{Name: &newName, Color: &newColor}

	expectedError := errors.New("mock error")

	service, bageRepo := generateBadgeUpdate()

	FindByIDCalledTimes := 0
	bageRepo.FindByIDFn = func(id, userID uint) (models.Badge, error) {
		FindByIDCalledTimes += 1
		return models.Badge{}, expectedError
	}

	bageRepo.UpdateFn = func(badge models.Badge) error {
		t.Error("update should not beeing called")
		return nil
	}

	err := service.Run(1, badgeUpdateDto, 1)
	if err == nil {
		t.Errorf("should return error but got nil")
	}

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected: " + expectedError.Error() + " but got: " + err.Error())
	}

	if FindByIDCalledTimes != 1 {
		t.Errorf("FindByID should have beeing called 1 time but got: %d",
			FindByIDCalledTimes)
	}
}

func TestShouldReturnErrorIfUpdateReturnsError(t *testing.T) {
	existingBadge := models.Badge{ID: 1, Name: "", Color: ""}

	newName := "a"
	newColor := "red"
	badgeUpdateDto := UpdateBadgeDto{Name: &newName, Color: &newColor}

	expectedError := errors.New("mock error")

	service, bageRepo := generateBadgeUpdate()

	FindByIDCalledTimes := 0
	bageRepo.FindByIDFn = func(id, userID uint) (models.Badge, error) {
		FindByIDCalledTimes += 1
		return existingBadge, expectedError
	}

	bageRepo.UpdateFn = func(badge models.Badge) error {
		return expectedError
	}

	err := service.Run(1, badgeUpdateDto, 1)
	if err == nil {
		t.Errorf("should return error but got nil")
	}

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected: " + expectedError.Error() + " but got: " + err.Error())
	}

	if FindByIDCalledTimes != 1 {
		t.Errorf("FindByID should have beeing called 1 time but got: %d",
			FindByIDCalledTimes)
	}
}
