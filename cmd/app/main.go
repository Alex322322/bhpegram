package main

import (
	"log"

	"github.com/Alex322322/bhpegram/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	log.Printf("starting app on port %s", cfg.App.Port)
}