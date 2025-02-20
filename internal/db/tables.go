package db

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type UserTable struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"not null"`
	Password string `gorm:"not null"`
}

func (UserTable) TableName() string {
	return "users"
}

type BankAccountTable struct {
	ID uint `gorm:"primaryKey;autoIncrement"`

	UserID    uint      `gorm:"index;not null"` // Foreign key
	UserTable UserTable `gorm:"foreignKey:UserID"`

	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
}

func (BankAccountTable) TableName() string {
	return "bank_accounts"
}

type CardTable struct {
	ID uint `gorm:"primaryKey;autoIncrement"`

	BankAccountID uint             `gorm:"index;not null"`
	BankAccount   BankAccountTable `gorm:"foreignKey:BankAccountID"`

	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
}

func (CardTable) TableName() string {
	return "cards"
}

type TransactionTable struct {
	ID uint `gorm:"primaryKey;autoIncrement"`

	Type   string  `gorm:"not null;"type:enum('income','expense','transfer')"`
	Method *string `gorm:"not null;"type:enum('credit_card','debit_card','other')"`

	Credit bool `gorm:"not null"`

	// 15 digits, 4 decimal places, Supports values up to $999,999,999,999.9999
	Value decimal.Decimal `gorm:"type:DECIMAL(19,4);not null"`
	Date  time.Time       `gorm:"type:DATETIME;not null;default:CURRENT_TIMESTAMP"`

	CardID *uint     `gorm:"index"`
	Card   CardTable `gorm:"foreignKey:CardID"`

	BankAccountID uint             `gorm:"index;not null"`
	BankAccount   BankAccountTable `gorm:"foreignKey:BankAccountID"`
}

func (TransactionTable) TableName() string {
	return "transactions"
}

type ItemTable struct {
	ID uint `gorm:"primaryKey;autoIncrement"`

	Name string `gorm:"not null"`
	// 15 digits, 4 decimal places, Supports values up to $999,999,999,999.9999
	Value    decimal.Decimal `gorm:"type:DECIMAL(19,4);not null"`
	Quantity uint            `gorm:"not null"`

	TransactionID uint             `gorm:"index"`
	Transaction   TransactionTable `gorm:"foreignKey:TransactionID"`

	Badges []BadgeTable `gorm:"many2many:item_badge;"`
}

func (ItemTable) TableName() string {
	return "items"
}

type ItemBadgeTable struct {
    // FIXME: so here I had to rename the collumns to item_table_id and
    // badge_table_id because thats what the lib was searching in the database
    // this may be a bug have to keep an eye on it
	ItemID  uint `gorm:"column:item_table_id;primaryKey"`
	BadgeID uint `gorm:"column:badge_table_id;primaryKey"`

	Item  ItemTable  `gorm:"foreignKey:ItemID;references:ID"`
	Badge BadgeTable `gorm:"foreignKey:BadgeID;references:ID"`
}

func (ItemBadgeTable) TableName() string {
	return "item_badge"
}

func (ItemBadgeTable) AddUniqueConstraint(db *gorm.DB) {
	db.Exec("ALTER TABLE item_badge ADD CONSTRAINT unique_item_badge UNIQUE (item_id, badge_id)")
}

type BadgeTable struct {
	ID uint `gorm:"primaryKey;autoIncrement"`

	Name string `gorm:"not null"`

	Items []ItemTable `gorm:"many2many:item_badge;"`
}

func (BadgeTable) TableName() string {
	return "badges"
}
