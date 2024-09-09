package main

import (
	"fmt"
	"hotPotBot/internal/bot"
	"hotPotBot/internal/config"
	"hotPotBot/internal/context"
	"hotPotBot/internal/db"
)

func main() {
	fmt.Println("Starting Application...")

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
		panic("Can not connect to database")
	}

	botHandler.Start(ctx)
}
