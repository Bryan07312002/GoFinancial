package main

import "financial/internal/db"

func main() {
	cfg := map[string]string{
		"DATABASE_PATH": "./db.db",
	}

	db.Seed(cfg)
}
