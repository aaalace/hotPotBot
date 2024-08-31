package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"hotPotBot/internal/keyboards"
	"hotPotBot/internal/logger"
)

func HandleCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		handleStartCommand(bot, message)
	default:
		logger.Log.Warnf("Unknown command: %s", message.Command())
	}
}

func handleStartCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "START_PAGE_TITLE")
	msg.ReplyMarkup = keyboards.Footer

	_, err := bot.Send(msg)
	if err != nil {
		logger.Log.Error("Error sending response to /start: " + err.Error())
	}
}
