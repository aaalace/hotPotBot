package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"hotPotBot/internal/logger"
)

func HandleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Text {
	case "ACCOUNT_BUTTON_TITLE":
		msg := tgbotapi.NewMessage(message.Chat.ID, "ACCOUNT_PAGE_TITLE")
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("CALLBACK1", "CALLBACK1_DATA"),
			),
		)

		_, err := bot.Send(msg)
		if err != nil {
			logger.Log.Error("Error sending response to /start: " + err.Error())
		}
	default:
		logger.Log.Warnf("Unknown message: %s", message.Text)
	}
}
