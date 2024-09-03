package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"hotPotBot/internal/context"
	"hotPotBot/internal/logger"
	"hotPotBot/internal/presentation/buttons"
	keyboards "hotPotBot/internal/presentation/keyboards/callbackKeyboards"
	"hotPotBot/internal/presentation/messages"
	"hotPotBot/internal/services"
)

func HandleMessage(ctx *context.AppContext, bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Text {
	case buttons.GetRandomCardButton:
		handleGetRandomCard(ctx, bot, message)
	case buttons.CardsStorageButton:
		handleCardsStorageMenu(bot, message)
	case buttons.HotPotStudioButton:
		handleHotPotStudioMenu(bot, message)
	case buttons.TutorialButton:
		handleTutorialButton(bot, message)
	default:
		logger.Log.Warnf("Unknown message: %s", message.Text)
	}
}

func handleGetRandomCard(ctx *context.AppContext, bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	randomCardService := services.RandomCardService{
		Ctx: ctx,
	}
	randomCardService.GenerateCard()
	// return some data
	msg := tgbotapi.NewMessage(message.Chat.ID, "not implemented")

	_, err := bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending response <random card service> | %v", err)
	}
}

func handleCardsStorageMenu(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, messages.CardsStoragePageTitle)
	msg.ReplyMarkup = keyboards.CardsStorageKeyboard

	_, err := bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending response <open cardsStorage> | %v", err.Error())
	}
}

func handleHotPotStudioMenu(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
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
