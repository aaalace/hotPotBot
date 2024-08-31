package main

import (
	"hotPotBot/internal/bot"
	"hotPotBot/internal/config"
)

func main() {
	cfg := config.NewConfig()
	// db conn
	botHandler := bot.NewBot(cfg)
	botHandler.Start()
}
