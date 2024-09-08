package utils

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hotPotBot/internal/logger"
)

func RemoveLoading(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	callbackConfig := tgbotapi.NewCallback(callback.ID, "")
	if _, err := bot.Request(callbackConfig); err != nil {
		logger.Log.Errorf("Error answering callback query | %v", err.Error())
	}
}
