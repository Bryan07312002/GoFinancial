package services

import (
	"financial/internal/db"
	"financial/internal/models"
)

type PaginateBadges struct {
	badgeRepo db.BadgeRepository
}

func NewPaginateBadges(
	badgeRepo db.BadgeRepository) PaginateBadges {
	return PaginateBadges{badgeRepo}
}

func (p *PaginateBadges) Run(
	paginateOpt db.PaginateOptions,
	userID uint,
) (db.PaginateResult[models.Badge], error) {
	return p.badgeRepo.PaginateFromUserID(paginateOpt, userID)
}
