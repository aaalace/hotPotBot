package main

import (
	"hotPotBot/internal/bot"
	"hotPotBot/internal/config"
	"hotPotBot/internal/context"
	"hotPotBot/internal/db"
)

func main() {
	configuration := config.NewConfig()

	database := db.ConnectDatabase(configuration)
	if database == nil {
		return
	}

	ctx := &context.AppContext{
		DB:           database,
		UserRequests: make(map[int64]string),
	}

	botHandler := bot.NewBot(configuration)
	if botHandler == nil {
		return
	}

	botHandler.Start(ctx)
}
