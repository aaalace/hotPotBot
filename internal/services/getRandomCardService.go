package services

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"hotPotBot/internal/logger"
)

type GetRandomCardService struct {
	Bot    *tgbotapi.BotAPI
	ChatId int64
}

func (service *GetRandomCardService) Start() {
	msg := service.Work()
	service.Send(msg)
}

func (service *GetRandomCardService) Send(msg tgbotapi.MessageConfig) {
	_, err := service.Bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending response <random card service> | %v", err)
	}
}

func (service *GetRandomCardService) Work() tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(service.ChatId, "not implemented")

	// some useful work

	return msg
}
