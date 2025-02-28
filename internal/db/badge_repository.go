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
	PaginateFromUserID(
		paginateOpt PaginateOptions,
		userID uint,
	) (PaginateResult[models.Badge], error)
	CreateMultiple(badges []models.Badge) ([]uint, error)
	GetMostExpansives(userID uint) ([]models.BadgeWithValue, error)
	Delete(id uint) error
}

type badgeRepository struct {
	db *gorm.DB
}

func ToBadgeTable(model models.Badge) BadgeTable {
	return BadgeTable{
		ID:    model.ID,
		Name:  model.Name,
		Color: model.Color,
	}
}

func ToBadge(table BadgeTable) models.Badge {
	return models.Badge{
		ID:    table.ID,
		Name:  table.Name,
		Color: table.Color,
	}
}

func NewBadgeRepository(db *gorm.DB) BadgeRepository {
	return &badgeRepository{db}
}

func (b *badgeRepository) Create(badge *models.Badge) (uint, error) {
	badgeTableInstance := BadgeTable{
		Name:  badge.Name,
		Color: badge.Color,
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
	badge := ToBadge(badgeTableInstance)
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
		badges = append(badges, ToBadge(badgeTable))
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
		badges = append(badges, ToBadge(badgeTable))
	}

	return badges, nil
}

func (b *badgeRepository) PaginateFromUserID(
	paginateOpt PaginateOptions,
	userID uint,
) (PaginateResult[models.Badge], error) {
	var count int64
	var tableInstances []BadgeTable

	query := b.db.Model(&BadgeTable{}).
		Where("user_id = ?", userID).
		Count(&count).
		Limit(int(paginateOpt.Take)).
		Offset(int((paginateOpt.Page - 1) * paginateOpt.Take))

	if paginateOpt.SortBy != "" {
		order := paginateOpt.SortBy
		if paginateOpt.SortDesc {
			order += " DESC"
		} else {
			order += " ASC"
		}
		query = query.Order(order)
	}

	query.Find(&tableInstances)

	results := make([]models.Badge, len(tableInstances))
	for i, acc := range tableInstances {
		results[i] = ToBadge(acc)
	}

	totalPages := count / int64(paginateOpt.Take)
	if count%int64(paginateOpt.Take) != 0 {
		totalPages++
	}

	return PaginateResult[models.Badge]{
		Data:        results,
		Total:       uint64(count),
		CurrentPage: paginateOpt.Page,
		PageSize:    paginateOpt.Take,
		TotalPages:  uint(totalPages),
	}, nil
}

func (b *badgeRepository) CreateMultiple(badges []models.Badge) ([]uint, error) {
	var badgeTableInstances []BadgeTable
	var badgeIDs []uint

	for _, badge := range badges {
		badgeTableInstances = append(badgeTableInstances, ToBadgeTable(badge))
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

func (b *badgeRepository) GetMostExpansives(userID uint) ([]models.BadgeWithValue, error) {
	query := `
    select badges.id, badges.name, badges.color, SUM(
      case
               when transactions.type='income' then items.value * items.quantity
               else -items.value * items.quantity
        end) as value
    from badges
    join item_badge on item_badge.badge_table_id=badges.id
    join items on item_badge.item_table_id=items.id
    join transactions on items.transaction_id=transactions.id
    join bank_accounts on transactions.bank_account_id=bank_accounts.id
    join users on bank_accounts.user_id=users.id
    where users.id=?
    group by badges.name
    order by value asc
    limit 5;
    `

	var badges []models.BadgeWithValue
	err := b.db.Raw(query, userID).Scan(&badges).Error
	if err != nil {
		return []models.BadgeWithValue{}, err
	}

	return badges[:], nil
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
