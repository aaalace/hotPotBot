package handlers

import (
	"fmt"
	"hotPotBot/internal/context"
	messageHandlers "hotPotBot/internal/handlers/messages"
	"hotPotBot/internal/logger"
	"hotPotBot/internal/presentation/buttons"
	callbackButtons "hotPotBot/internal/presentation/buttons/callbackButtons"
	"hotPotBot/internal/utils"
	"regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleMessage(ctx *context.AppContext, bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Text {

	// Footer buttons
	case buttons.GetRandomCardButton:
		messageHandlers.HandleGetRandomCard(ctx, bot, message)
	case buttons.CardsStorageButton:
		messageHandlers.HandleCardsStorageMenu(bot, message)
	case buttons.HotPotStudioButton:
		messageHandlers.HandleHotPotStudioMenu(bot, message)
	case buttons.TutorialButton:
		messageHandlers.HandleTutorialButton(bot, message)

	// Just messages
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
		// pattern for check if data is from exchange this card
		exchangePattern := fmt.Sprintf(`^%s&\d+$`, callbackButtons.ExchangeThisCardInlineButton.Data)
		exchangeRe := regexp.MustCompile(exchangePattern)
		if exchangeRe.MatchString(prevReq) {
			messageHandlers.HandleExchangeToAccount(ctx, bot, message, prevReq)
		} else {
			logger.Log.Warnf("Unknown message")	
		}
	}

}
