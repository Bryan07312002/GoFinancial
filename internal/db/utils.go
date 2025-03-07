package db

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type PaginateOptions struct {
	Page     uint
	Take     uint
	SortBy   string
	SortDesc bool
	TimeWindowSearch
}

// PaginateResult holds the result of a paginated query
type PaginateResult[T any] struct {
	Data        []T    `json:"data"`
	Total       uint64 `json:"total"`
	CurrentPage uint   `json:"current_page"`
	PageSize    uint   `json:"page_size"`
	TotalPages  uint   `json:"total_pages"`
}

type TimeWindowSearch struct {
	Start  time.Time `json:"from"` // default should be 'time.Now()'
	Finish time.Time `json:"to"`   // default should be zero
}

func InitDatabase(driver Driver, config map[string]string) *gorm.DB {
	if driver == SqliteDriver {
		path := config["DATABASE_PATH"]

		println("Connecting to database...")
		db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		println("Connected to database!")

		println("Migrateing tables...")
		err = db.AutoMigrate(
			&UserTable{},
			&BankAccountTable{},
			&TransactionTable{},
			&CardTable{},
			&ItemTable{},
			&BadgeTable{},
			&ItemBadgeTable{},
		)
		if err != nil {
			panic("failed to migrate database")
		}
		println("Finish migrateing!")

		return db
	}

	// Add other driver logic (MySQL, Postgres, etc...)
	panic("unsupported driver")
}
