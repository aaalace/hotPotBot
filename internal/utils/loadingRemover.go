package utils

import (
	"hotPotBot/internal/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func RemoveLoading(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	callbackConfig := tgbotapi.NewCallback(callback.ID, "")
	if _, err := bot.Request(callbackConfig); err != nil {
		logger.Log.Errorf("Error answering callback query | %v", err.Error())
	}
}
