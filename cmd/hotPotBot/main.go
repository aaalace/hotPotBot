package main

import (
	"hotPotBot/internal/bot"
	"hotPotBot/internal/config"
	"hotPotBot/internal/context"
	"hotPotBot/internal/db"
	"hotPotBot/internal/logger"
)

func main() {
	logger.Log.Info("Start configuring...")

	configuration := config.NewConfig()

	database := db.ConnectDatabase(configuration)
	if database == nil {
		panic("Can not connect to database")
	}

	ctx := &context.AppContext{
		DB:           database,
		UserRequests: make(map[int64]string),
	}

	botHandler := bot.NewBot(configuration)
	if botHandler == nil {
		panic("Can not connect to telegram bot")
	}

	logger.Log.Info("Start polling...")
	botHandler.Start(ctx)
}
