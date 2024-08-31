package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"hotPotBot/internal/logger"
	"hotPotBot/internal/services"
)

func HandleCallback(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	switch callback.Data {
	case "CALLBACK1_DATA":
		handleCallBack1(bot, callback.Message)
	default:
		logger.Log.Warnf("Unknown callback: %s", callback.Message.Text)
	}
}

func handleCallBack1(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	mockService := services.MockService{}
	mockService.SendMockAnswer(bot, message)
}
