package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"hotPotBot/internal/logger"
	"hotPotBot/internal/presentation/buttons"
	keyboards "hotPotBot/internal/presentation/keyboards/callbackKeyboards"
	"hotPotBot/internal/presentation/messages"
	"hotPotBot/internal/services"
)

func HandleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Text {
	case buttons.GetCardButton:
		handleGetCardButton(bot, message)
	case buttons.CardsStorageButton:
		handleCardsStorageButton(bot, message)
	case buttons.HotPotStudioButton:
		handleHotPotStudioButton(bot, message)
	case buttons.TutorialButton:
		handleTutorialButton(bot, message)
	default:
		logger.Log.Warnf("Unknown message: %s", message.Text)
	}
}

func handleGetCardButton(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	service := services.GetRandomCardService{
		Bot:    bot,
		ChatId: message.Chat.ID,
	}
	service.Start()
}

func handleCardsStorageButton(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, messages.CardsStoragePageTitle)
	msg.ReplyMarkup = keyboards.CardsStorageKeyboard

	_, err := bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending response <open cardsStorage> | %v", err.Error())
	}
}

func handleHotPotStudioButton(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, messages.HotPotStudioPageTitle)
	msg.ReplyMarkup = keyboards.HotPotStudioKeyboard

	_, err := bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending response <open studio> | %v", err.Error())
	}
}

func handleTutorialButton(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, messages.TutorialText)

	_, err := bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending response <open tutorial> | %v", err.Error())
	}
}
