package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message != nil {
		if update.Message.IsCommand() {
			HandleCommand(bot, update.Message)
		} else {
			HandleMessage(bot, update.Message)
		}
	} else if update.CallbackQuery != nil {
		HandleCallback(bot, update.CallbackQuery)
	}
}
