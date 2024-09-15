package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"hotPotBot/internal/logger"
	"os"
)

type Config struct {
	TelegramBotToken string
	DatabasePath     string
	S3Path           string
	S3AccessKey      string
	S3SecretKey      string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		logger.Log.Fatalf("Error loading env variables | %v", err.Error())
	}

	databasePath := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRESQL_USER"),
		os.Getenv("POSTGRESQL_PASSWORD"),
		os.Getenv("POSTGRESQL_HOST"),
		os.Getenv("POSTGRESQL_PORT"),
		os.Getenv("POSTGRESQL_DBNAME"))

	return &Config{
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		DatabasePath:     databasePath,
		S3Path:           os.Getenv("S3_PATH"),
		S3AccessKey:      os.Getenv("S3_ACCESS_KEY"),
		S3SecretKey:      os.Getenv("S3_SECRET_KEY"),
	}
}
