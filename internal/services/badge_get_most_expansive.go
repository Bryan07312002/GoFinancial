package services

import (
	"financial/internal/db"
	"financial/internal/models"
)

type MostExpansiveBadges interface {
	Run(userID uint) ([]models.BadgeWithValue, error)
}

type mostExpansiveBadges struct {
	badgeRepo db.BadgeRepository
}

func NewMostExpansiveBadges(badgeRepo db.BadgeRepository) MostExpansiveBadges {
	return &mostExpansiveBadges{badgeRepo}

}

func (m *mostExpansiveBadges) Run(userID uint) ([]models.BadgeWithValue, error) {
	return m.badgeRepo.GetMostExpansives(userID)
}
