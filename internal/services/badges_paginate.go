package services

import (
	"financial/internal/db"
	"financial/internal/models"
)

type PaginateBadges interface {
	Run(
		paginateOpt db.PaginateOptions,
		userID uint,
	) (db.PaginateResult[models.Badge], error)
}

type paginateBadges struct {
	badgeRepo db.BadgeRepository
}

func NewPaginateBadges(
	badgeRepo db.BadgeRepository) PaginateBadges {
	return &paginateBadges{badgeRepo}
}

func (p *paginateBadges) Run(
	paginateOpt db.PaginateOptions,
	userID uint,
) (db.PaginateResult[models.Badge], error) {
	return p.badgeRepo.PaginateFromUserID(paginateOpt, userID)
}
