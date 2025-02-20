package db

import (
	"financial/internal/models"
	"gorm.io/gorm"
)

type BadgeRepository interface {
	Create(badge *models.Badge) (uint, error)
	LinkItemToBadge(itemID uint, badgeID uint) error
	FindByID(id uint) (models.Badge, error)
	FindByItem(itemID uint) ([]models.Badge, error)
	FindByTransaction(transactionID uint) ([]models.Badge, error)
	CreateMultiple(badges []models.Badge) ([]uint, error)
	Delete(id uint) error
}

type badgeRepository struct {
	db *gorm.DB
}

func NewBadgeRepository(db *gorm.DB) BadgeRepository {
	return &badgeRepository{db}
}

func (b *badgeRepository) Create(badge *models.Badge) (uint, error) {
	badgeTableInstance := BadgeTable{
		Name: badge.Name,
	}

	if err := b.db.Create(&badgeTableInstance).Error; err != nil {
		return 0, err
	}

	return badgeTableInstance.ID, nil
}

func (r *badgeRepository) LinkItemToBadge(itemID uint, badgeID uint) error {
	// Create a new ItemBadgeTable entry
	itemBadge := ItemBadgeTable{
		ItemID:  itemID,
		BadgeID: badgeID,
	}

	// Insert the link into the item_badge table
	if err := r.db.Create(&itemBadge).Error; err != nil {
		return err
	}

	return nil
}

// FindByID retrieves a badge by its ID
func (r *badgeRepository) FindByID(id uint) (models.Badge, error) {
	var badgeTableInstance BadgeTable
	if err := r.db.First(&badgeTableInstance, id).Error; err != nil {
		return models.Badge{}, err
	}

	// Convert the BadgeTable instance to models.Badge
	badge := models.Badge{
		ID:   badgeTableInstance.ID,
		Name: badgeTableInstance.Name,
	}

	return badge, nil
}

func (b *badgeRepository) FindByItem(itemID uint) ([]models.Badge, error) {
	var badgeTables []BadgeTable

	// Get all badges associated with the item
	if err := b.db.Joins("JOIN item_badge ON badges.id = item_badge.badge_table_id").
		Where("item_badge.item_table_id = ?", itemID).
		Find(&badgeTables).Error; err != nil {
		return nil, err
	}

	var badges []models.Badge
	for _, badgeTable := range badgeTables {
		badges = append(badges, models.Badge{
			ID:   badgeTable.ID,
			Name: badgeTable.Name,
		})
	}

	return badges, nil
}

func (b *badgeRepository) FindByTransaction(transactionID uint) ([]models.Badge, error) {
	var badgeTables []BadgeTable

	// Get all badges associated with the transaction
	if err := b.db.Joins("JOIN item_badge ON badges.id = item_badge.badge_table_id").
		Joins("JOIN items ON items.id = item_badge.item_table_id").
		Where("items.transaction_id = ?", transactionID).
		Find(&badgeTables).Error; err != nil {
		return nil, err
	}

	var badges []models.Badge
	for _, badgeTable := range badgeTables {
		badges = append(badges, models.Badge{
			ID:   badgeTable.ID,
			Name: badgeTable.Name,
		})
	}

	return badges, nil
}

func (b *badgeRepository) CreateMultiple(badges []models.Badge) ([]uint, error) {
	var badgeTableInstances []BadgeTable
	var badgeIDs []uint

	for _, badge := range badges {
		badgeTableInstances = append(badgeTableInstances, BadgeTable{
			Name: badge.Name,
		})
	}

	if err := b.db.Create(&badgeTableInstances).Error; err != nil {
		return nil, err
	}

	// Collect the IDs of the created badges
	for _, badgeTable := range badgeTableInstances {
		badgeIDs = append(badgeIDs, badgeTable.ID)
	}

	return badgeIDs, nil
}

func (b *badgeRepository) Delete(id uint) error {
	// Delete the badge by ID
	result := b.db.Delete(&BadgeTable{}, id)

	// Handle database errors (e.g., connection issues)
	if result.Error != nil {
		return result.Error
	}

	// Check if no rows were affected (record not found)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	// Successfully deleted
	return nil
}
