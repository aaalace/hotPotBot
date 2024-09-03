package config

import (
	"github.com/joho/godotenv"
	"hotPotBot/internal/logger"
	"os"
)

type Config struct {
	TelegramBotToken string
	DatabasePath     string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		logger.Log.Fatalf("Error loading .env file: %v", err)
	}

	return &Config{
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		DatabasePath:     os.Getenv("DATABASE_PATH"),
	}
}
