package main

import (
	"financial/internal/api"
	"fmt"
)

func main() {
	cfg := map[string]string{
		"DATABASE_PATH": "./db.db",
	}

	srv := api.NewServer(&api.Config{
		ServerPort: "8080",
		DBConfig:   &cfg,
	})

	if err := srv.Run(); err != nil {
		fmt.Errorf("Could not start server: %v", err)
	}
}
