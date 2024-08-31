package services

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"hotPotBot/internal/logger"
)

type MockService struct{}

func (es *MockService) SendMockAnswer(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "MOCK_ANSWER")

	_, err := bot.Send(msg)
	if err != nil {
		logger.Log.Error("Error sending response to message: " + err.Error())
	}
}
