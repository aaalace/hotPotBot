package messageHandlers

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hotPotBot/internal/context"
	"hotPotBot/internal/logger"
	keyboards "hotPotBot/internal/presentation/keyboards/callbackKeyboards"
	"hotPotBot/internal/presentation/messages"
	"hotPotBot/internal/services"
	"hotPotBot/internal/utils"
)

// HandleGetRandomCard - Получить карту
func HandleGetRandomCard(ctx *context.AppContext, bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	userService := services.UserService{Ctx: ctx}
	user, err := userService.GetUserByTelegramId(message.From.ID)
	if err != nil {
		logger.Log.Errorf("Error in getting user: %v", err)
		return
	}

	randomCardService := services.RandomCardService{Ctx: ctx}
	card, err := randomCardService.GetRandomCard(user.Id)
	if err != nil {
		if errors.As(err, &services.NotEnoughTime{}) {
			msg := tgbotapi.NewMessage(message.Chat.ID, err.Error())
			_, err = bot.Send(msg)
			if err != nil {
				logger.Log.Errorf("Error sending response <random card service cooldown> | %v", err.Error())
			}
			return
		}
		logger.Log.Errorf("Error in getting random card: %v", err)
		return
	}

	cardService := services.CardService{Ctx: ctx}
	typeName, err := cardService.GetTypeNameByTypeId(card.TypeId)
	if err != nil {
		logger.Log.Errorf("Error in getting typename: %v", err)
		return
	}

	photo := tgbotapi.NewPhoto(message.Chat.ID, tgbotapi.FilePath(card.ImageUrl))
	view := utils.GenerateRandomCardView(card, typeName)
	photo.Caption = view
	_, err = bot.Send(photo)
	if err != nil {
		logger.Log.Errorf("Error sending response <random card service> | %v", err.Error())
	}
}

// HandleCardsStorageMenu - Мои карты
func HandleCardsStorageMenu(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, messages.CardsStoragePageTitle)
	msg.ReplyMarkup = keyboards.CardsStorageKeyboard
	_, err := bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending response <open cardsStorage> | %v", err.Error())
	}
}

// HandleHotPotStudioMenu - Hot Pot Studio
func HandleHotPotStudioMenu(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, messages.HotPotStudioPageTitle)
	msg.ReplyMarkup = keyboards.HotPotStudioKeyboard
	_, err := bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending response <open studio> | %v", err.Error())
	}
}

// HandleTutorialButton - Туториал
func HandleTutorialButton(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, messages.TutorialText)
	_, err := bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending response <open tutorial> | %v", err.Error())
	}
}
