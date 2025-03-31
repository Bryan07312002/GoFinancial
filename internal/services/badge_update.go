package services

import "financial/internal/db"

type UpdateBadge interface {
	Run(badgeID uint, dto BadgeUpdateDto, userID uint) error
}

type updateBadge struct {
	badgeRepo db.BadgeRepository
}

func NewUpdateBadge(badgeRepo db.BadgeRepository) UpdateBadge {
	return &updateBadge{badgeRepo}
}

type BadgeUpdateDto struct {
	Name  *string `json:"name,omitempty"`
	Color *string `json:"color,omitempty"`
}

func (u *updateBadge) Run(badgeID uint, dto BadgeUpdateDto, userID uint) error {
	badge, err := u.badgeRepo.FindByID(badgeID, userID)
	if err != nil {
		return err
	}

	if dto.Name != nil {
		badge.Name = *dto.Name
	}

	if dto.Color != nil {
		badge.Color = *dto.Color
	}

	return u.badgeRepo.Update(badge)
}
