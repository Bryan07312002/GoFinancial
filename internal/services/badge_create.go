package services

import (
	"financial/internal/db"
	"financial/internal/models"
)

type CreateBadge struct {
	badgeRepo db.BadgeRepository
}

func NewCreateBadge(badgeRepo db.BadgeRepository) CreateBadge {
	return CreateBadge{badgeRepo}
}

type NewBadge struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

func (c *CreateBadge) Run(badge NewBadge, userId uint) error {
	_, err := c.badgeRepo.Create(&models.Badge{
		Name:  badge.Name,
		Color: badge.Color,
	})

	return err
}
