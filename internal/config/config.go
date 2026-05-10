package config

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Auth     AuthConfig
}

type AppConfig struct {
	Port string
}

type AuthConfig struct {
	JWTSecret string `validate:"required"`
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string `validate:"required"`
	Name     string `validate:"required"`
	SSLMode  string
}

func Load() (*Config, error) {
	godotenv.Load()

	config := &Config{}

	config.App.Port = getEnv("APP_PORT", "8080")

	config.Database.Host = getEnv("DB_HOST", "localhost")
	config.Database.Port = getEnv("DB_PORT", "5432")
	config.Database.User = getEnv("DB_USER", "postgres")
	config.Database.Password = getEnv("DB_PASSWORD", "")
	config.Database.Name = getEnv("DB_NAME", "bhpegram")
	config.Database.SSLMode = getEnv("DB_SSLMODE", "disable")

	config.Auth.JWTSecret = getEnv("JWT_SECRET", "")

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		return nil, err
	}

	return config, nil
}

// getEnv читает переменную окружения, возвращает defaultValue если не задана
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
