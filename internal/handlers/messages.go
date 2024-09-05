package handlers

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hotPotBot/internal/context"
	"hotPotBot/internal/logger"
	"hotPotBot/internal/presentation/buttons"
	callbackButtons "hotPotBot/internal/presentation/buttons/callbackButtons"
	keyboards "hotPotBot/internal/presentation/keyboards/callbackKeyboards"
	"hotPotBot/internal/presentation/messages"
	"hotPotBot/internal/services"
	"hotPotBot/internal/utils"
)

func HandleMessage(ctx *context.AppContext, bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Text {
	// Footer buttons
	case buttons.GetRandomCardButton:
		handleGetRandomCard(ctx, bot, message)
	case buttons.CardsStorageButton:
		handleCardsStorageMenu(bot, message)
	case buttons.HotPotStudioButton:
		handleHotPotStudioMenu(bot, message)
	case buttons.TutorialButton:
		handleTutorialButton(bot, message)
	// Dynamic messages
	default:
		HandleDynamicMessage(ctx, bot, message)
	}
}

func HandleDynamicMessage(ctx *context.AppContext, bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	ctx.Mu.Lock()
	value := ctx.UserRequests[message.From.ID]
	ctx.Mu.Unlock()

	switch value {
	case callbackButtons.FindUserInlineButton.Data:
		handleOtherAccount(ctx, bot, message)
	default:
		logger.Log.Warnf("Unknown message: %s", message.Text)
	}

}

// Получить карту
func handleGetRandomCard(ctx *context.AppContext, bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
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

// Мои карты
func handleCardsStorageMenu(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, messages.CardsStoragePageTitle)
	msg.ReplyMarkup = keyboards.CardsStorageKeyboard
	_, err := bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending response <open cardsStorage> | %v", err.Error())
	}
}

// Hot Pot Studio
func handleHotPotStudioMenu(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, messages.HotPotStudioPageTitle)
	msg.ReplyMarkup = keyboards.HotPotStudioKeyboard
	_, err := bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending response <open studio> | %v", err.Error())
	}
}

// Туториал
func handleTutorialButton(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, messages.TutorialText)
	_, err := bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending response <open tutorial> | %v", err.Error())
	}
}

// @username
func handleOtherAccount(ctx *context.AppContext, bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	// important!!!
	utils.RemoveUserPreviousRequest(ctx, message.From.ID)

	msg := tgbotapi.NewMessage(message.Chat.ID, messages.UserNotFoundError)

	userService := services.UserService{Ctx: ctx}
	user, err := userService.GetUserByUsername(message.Text)
	if err != nil {
		logger.Log.Errorf("Error in getting user: %v", err.Error())
		_, err = bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("Error sending response <handleOtherAccount 1_2>: %v", err.Error())
		}
		return
	}
	weight, err := userService.CountUserWeight(user.Id)
	if err != nil {
		logger.Log.Errorf("Error in count user weight: %v", err.Error())
		_, err = bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("Error sending response <handleOtherAccount 2_2>: %v", err.Error())
		}
		return
	}

	accountView := utils.GenerateAccountView(user.TelegramUsername, weight)
	msg = tgbotapi.NewMessage(message.Chat.ID, accountView)
	_, err = bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending response <handleOtherAccount>: %v", err.Error())
	}
}
