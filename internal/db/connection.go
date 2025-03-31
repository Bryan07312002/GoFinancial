package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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
		err = applyAutoMigrate(db)
        if err != nil {
			panic("failed to migrate database")
		}
		println("Finish migrateing!")

		return db
	}

	// Add other driver logic (MySQL, Postgres, etc...)
	panic("unsupported driver")
}

func applyAutoMigrate(con *gorm.DB) error {
	return con.AutoMigrate(
		&UserTable{},
		&BankAccountTable{},
		&TransactionTable{},
		&CardTable{},
		&ItemTable{},
		&BadgeTable{},
		&ItemBadgeTable{},
	)
}
