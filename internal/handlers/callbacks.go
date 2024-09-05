package handlers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hotPotBot/internal/context"
	callbackHandlers "hotPotBot/internal/handlers/callbacks"
	"hotPotBot/internal/logger"
	buttons "hotPotBot/internal/presentation/buttons/callbackButtons"
	"regexp"
)

func HandleCallback(ctx *context.AppContext, bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	switch callback.Data {
	// My Cards
	case buttons.AllCardsInlineButton.Data:
		callbackHandlers.HandleAllCardsButton(ctx, bot, callback)
	case buttons.SingleCardsInlineButton.Data:
		callbackHandlers.HandleSingleCardsButton(ctx, bot, callback)
	case buttons.AlbumCardsInlineButton.Data:
		callbackHandlers.HandleAlbumCardsButton(ctx, bot, callback)
	case buttons.DuplicatesInlineButton.Data:
		callbackHandlers.HandleDuplicatesButton(ctx, bot, callback)
	// Hot Pot Studio
	case buttons.MyAccountInlineButton.Data:
		callbackHandlers.HandleMyAccount(ctx, bot, callback)
	case buttons.FindUserInlineButton.Data:
		callbackHandlers.HandleOtherAccountButton(ctx, bot, callback)
	case buttons.ShopInlineButton.Data:
	case buttons.ExchangeInlineButton.Data:
	case buttons.CraftInlineButton.Data:
	case buttons.DiceInlineButton.Data:

	default:
		// pattern for check if data is from left/right arrows
		arrowPattern := fmt.Sprintf(`^%s&\d+$`, buttons.LeftInlineButton.Data)
		arrowRe := regexp.MustCompile(arrowPattern)
		if arrowRe.MatchString(callback.Data) {
			callbackHandlers.HandleArrowButton(ctx, bot, callback)
		} else {
			logger.Log.Warnf("Unknown callback: %s", callback.Message.Text)
		}
	}
}
