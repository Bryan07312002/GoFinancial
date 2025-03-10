package db

import (
	"financial/internal/models"

	"gorm.io/gorm"
)

func ToItem(itemTable ItemTable) models.Item {
	return models.Item{
		ID:            itemTable.ID,
		TransactionID: itemTable.TransactionID,
		Name:          itemTable.Name,
		Value:         itemTable.Value,
		Quantity:      itemTable.Quantity,
	}
}

func ToItemTable(item models.Item) ItemTable {
	return ItemTable{
		ID:            item.ID,
		TransactionID: item.TransactionID,
		Name:          item.Name,
		Value:         item.Value,
		Quantity:      item.Quantity,
	}
}

type ItemRepository interface {
	CreateMultiple(items []models.Item) ([]uint, error)
	FindByID(ID uint) (models.Item, error)
	Delete(id uint) error
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(con *gorm.DB) ItemRepository {
	return &itemRepository{db: con}
}

func (i *itemRepository) Create(item models.Item) (uint, error) {
	itemTableInstance := ToItemTable(item) // Convert to database model

	err := i.db.Create(&itemTableInstance).Error
	if err != nil {
		return 0, err
	}

	return itemTableInstance.ID, nil
}

func (i *itemRepository) CreateMultiple(items []models.Item) ([]uint, error) {
	if len(items) == 0 {
		return []uint{}, nil
	}

	// Convert the slice of models.Item to a slice of ItemTable
	var itemTableInstances []ItemTable
	for _, item := range items {
		itemTableInstances = append(itemTableInstances, ToItemTable(item))
	}

	// Insert all items into the database at once
	result := i.db.Create(&itemTableInstances).Debug()
	if result.Error != nil {
		return nil, result.Error
	}

	// Extract the IDs of the created items
	var ids []uint
	for _, item := range itemTableInstances {
		ids = append(ids, item.ID)
	}

	return ids, nil
}

func (i *itemRepository) FindByID(ID uint) (models.Item, error) {
	var itemTableInstance ItemTable

	if err := i.db.Preload("Badges").First(&itemTableInstance, ID).Error; err != nil {
		return models.Item{}, err
	}

	return ToItem(itemTableInstance), nil
}

func (i *itemRepository) Delete(id uint) error {
	result := i.db.Delete(&ItemTable{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
