package main

import (
	"log"

	"github.com/Alex322322/bhpegram/internal/config"
	"github.com/Alex322322/bhpegram/internal/database"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := database.NewPostgresDB(cfg.Database)
	if err != nil {
		log.Fatalf("database error: %v", err)
	}
	defer db.Close()

	log.Printf("starting app on port %s", cfg.App.Port)
}