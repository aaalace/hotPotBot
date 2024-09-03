package main

import (
	"hotPotBot/internal/bot"
	"hotPotBot/internal/config"
)

func main() {
	// recover?

	configuration := config.NewConfig()
	// db conn
	botHandler := bot.NewBot(configuration)
	botHandler.Start()
}
