package services

import "financial/internal/db"

type UpdateBadge struct {
	badgeRepo db.BadgeRepository
}

func NewUpdateBadge(badgeRepo db.BadgeRepository) UpdateBadge {
	return UpdateBadge{badgeRepo}
}

type BadgeUpdateDto struct {
	Name  *string `json:"name,omitempty"`
	Color *string `json:"color,omitempty"`
}

func (u *UpdateBadge) Run(badgeID uint, dto BadgeUpdateDto, userID uint) error {
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
