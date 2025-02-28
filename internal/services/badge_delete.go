package services

import "financial/internal/db"

type DeleteBadge struct {
	BadgeRepo db.BadgeRepository
}

func NewDeleteBadge(BadgeRepo db.BadgeRepository) DeleteBadge {
	return DeleteBadge{BadgeRepo}
}

// TODO: check if badge refers to user id received
func (d *DeleteBadge) Run(badgeId, userID uint) error {
	return d.BadgeRepo.Delete(badgeId)
}
