package database

import (
	"fmt"
	"time"

	"github.com/Alex322322/bhpegram/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	maxRetries = 5
	retryDelay = 2 * time.Second
)

func NewPostgresDB(cfg config.DatabaseConfig) (*sqlx.DB, error) {
	var db *sqlx.DB
	var err error

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode,
	)

	for i := range maxRetries {
		db, err = sqlx.Connect("postgres", dsn)
		if err == nil {
			break
		}

		fmt.Printf("failed to connect to db, attempt %d/%d: %v\n", i+1, maxRetries, err)
		time.Sleep(retryDelay)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after %d attempts: %w", maxRetries, err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}
