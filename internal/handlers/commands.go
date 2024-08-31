package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
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
	msg := tgbotapi.NewMessage(message.Chat.ID, "Hello")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("ACCOUNT_BUTTON_TITLE"),
		),
	)

	_, err := bot.Send(msg)
	if err != nil {
		logger.Log.Error("Error sending response to /start: " + err.Error())
	}
}
