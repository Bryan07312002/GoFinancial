package services

import (
	"financial/internal/db"

	"github.com/shopspring/decimal"
)

type UpdateItemDto struct {
	Name         *string
	Value        *decimal.Decimal
	Quantity     *uint
	addBadges    *[]uint
	removeBadges *[]uint
}

type UpdateItem interface {
	Run(itemId uint, dto UpdateItemDto, userID uint) error
}

type updateItem struct {
	itemRepository db.ItemRepository
}

func NewUpdateItem(itemRepository db.ItemRepository) UpdateItem {
	return &updateItem{itemRepository}
}

// FIXME: check if given badges id belongs to given userID
func (u *updateItem) Run(itemId uint, dto UpdateItemDto, userID uint) error {
	item, err := u.itemRepository.FindByID(itemId)
	if err != nil {
		return err
	}

	if dto.Name != nil {
		item.Name = *dto.Name
	}

	if dto.Value != nil {
		item.Value = *dto.Value
	}

	if dto.Quantity != nil {
		item.Quantity = *dto.Quantity
	}

	if dto.addBadges == nil {
		noAddedBadges := []uint{}
		dto.addBadges = &noAddedBadges
	}

	if dto.removeBadges == nil {
		noRemoveBadges := []uint{}
		dto.removeBadges = &noRemoveBadges
	}

	u.itemRepository.Update(item, *dto.removeBadges, *dto.addBadges)

	return nil
}
