package services

import (
	"financial/internal/db"
	"financial/internal/models"
)

type MostExpansiveBadges struct {
	badgeRepo db.BadgeRepository
}

func NewMostExpansiveBadges(badgeRepo db.BadgeRepository) MostExpansiveBadges {
	return MostExpansiveBadges{badgeRepo}

}

func (m *MostExpansiveBadges) Run(userID uint) ([]models.BadgeWithValue, error) {
	return m.badgeRepo.GetMostExpansives(userID)
}
