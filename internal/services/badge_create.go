package services

import (
	"financial/internal/db"
	"financial/internal/models"
)

type CreateBadge interface {
	Run(badge NewBadge, userId uint) error
}

type createBadge struct {
	badgeRepo db.BadgeRepository
}

func NewCreateBadge(badgeRepo db.BadgeRepository) CreateBadge {
	return &createBadge{badgeRepo}
}

type NewBadge struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

func (c *createBadge) Run(badge NewBadge, userId uint) error {
	_, err := c.badgeRepo.Create(&models.Badge{
		Name:  badge.Name,
		Color: badge.Color,
	})

	return err
}
