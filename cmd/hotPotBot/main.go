package main

import (
	"hotPotBot/internal/bot"
	"hotPotBot/internal/config"
	"hotPotBot/internal/context"
	"hotPotBot/internal/db"
	"hotPotBot/internal/logger"
	"hotPotBot/internal/s3"
)

func main() {
	logger.Log.Info("Start configuring...")

	configuration := config.NewConfig()

	database := db.ConnectDatabase(configuration)
	if database == nil {
		panic("Can not connect to database")
	}

	s3Client := s3.ConnectS3Client(configuration)
	if s3Client == nil {
		panic("Can not connect to S3")
	}

	botHandler := bot.NewBot(configuration)
	if botHandler == nil {
		panic("Can not connect to telegram bot")
	}

	logger.Log.Info("Start polling...")
	botHandler.Start(&context.AppContext{
		DB:           database,
		S3Client:     s3Client,
		UserRequests: make(map[int64]string),
	})
}
