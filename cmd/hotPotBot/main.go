package main

import (
	"hotPotBot/internal/bot"
	"hotPotBot/internal/config"
	"hotPotBot/internal/context"
	"hotPotBot/internal/db"
)

func main() {
	// recover

	configuration := config.NewConfig()

	database := db.ConnectDatabase(configuration)
	ctx := &context.AppContext{
		DB:           database,
		UserRequests: make(map[int64]string),
	}

	botHandler := bot.NewBot(configuration)
	botHandler.Start(ctx)
}
