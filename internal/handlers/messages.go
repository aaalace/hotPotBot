package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hotPotBot/internal/context"
	messageHandlers "hotPotBot/internal/handlers/messages"
	"hotPotBot/internal/logger"
	"hotPotBot/internal/presentation/buttons"
	callbackButtons "hotPotBot/internal/presentation/buttons/callbackButtons"
	"hotPotBot/internal/utils"
)

func HandleMessage(ctx *context.AppContext, bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Text {
	// footer buttons
	case buttons.GetRandomCardButton:
		messageHandlers.HandleGetRandomCard(ctx, bot, message)
	case buttons.CardsStorageButton:
		messageHandlers.HandleCardsStorageMenu(bot, message)
	case buttons.HotPotStudioButton:
		messageHandlers.HandleHotPotStudioMenu(bot, message)
	case buttons.TutorialButton:
		messageHandlers.HandleTutorialButton(bot, message)
	// just messages
	default:
		HandleDynamicMessage(ctx, bot, message)
	}
}

func HandleDynamicMessage(ctx *context.AppContext, bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	prevReq := utils.GetRmUserPreviousRequest(ctx, message.From.ID)

	switch prevReq {
	case callbackButtons.FindUserInlineButton.Data:
		messageHandlers.HandleOtherAccount(ctx, bot, message)
	default:
		logger.Log.Warnf("Unknown message: %s", message.Text)
	}

}
