package services

import "financial/internal/db"

type DeleteBadge interface {
	Run(badgeId, userID uint) error
}

type deleteBadge struct {
	BadgeRepo db.BadgeRepository
}

func NewDeleteBadge(BadgeRepo db.BadgeRepository) DeleteBadge {
	return &deleteBadge{BadgeRepo}
}

// TODO: check if badge refers to user id received
func (d *deleteBadge) Run(badgeId, userID uint) error {
	return d.BadgeRepo.Delete(badgeId)
}
