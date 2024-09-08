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
	// My Cards Menu
	case buttons.AllCardsInlineButton.Data:
		callbackHandlers.HandleAllCardsButton(ctx, bot, callback)
	case buttons.SingleCardsInlineButton.Data:
		callbackHandlers.HandleSingleCardsButton(ctx, bot, callback)
	case buttons.AlbumCardsInlineButton.Data:
		callbackHandlers.HandleAlbumCardsButton(ctx, bot, callback)
	case buttons.DuplicatesInlineButton.Data:
		callbackHandlers.HandleDuplicatesButton(ctx, bot, callback)

	// Hot Pot Studio Menu
	case buttons.MyAccountInlineButton.Data:
		callbackHandlers.HandleMyAccount(ctx, bot, callback)
	case buttons.FindUserInlineButton.Data:
		callbackHandlers.HandleOtherAccountButton(ctx, bot, callback)
	case buttons.ShopInlineButton.Data:
		callbackHandlers.HandleShopButton(bot, callback)
	case buttons.ExchangeInlineButton.Data:
	case buttons.CraftInlineButton.Data:
		callbackHandlers.HandleCraftButton(bot, callback)
	case buttons.DiceInlineButton.Data:

	// Shop Menu
	case buttons.ShopAllCardsInlineButton.Data:
		callbackHandlers.HandleShopAllCardsButton(ctx, bot, callback)
	case buttons.ShopSingleCardsInlineButton.Data:
		callbackHandlers.HandleShopSingleCardsButton(ctx, bot, callback)
	case buttons.ShopAlbumCardsInlineButton.Data:
		callbackHandlers.HandleShopAlbumCardsButton(ctx, bot, callback)

	// Craft Menu
	case buttons.CraftAlbumInlineButton.Data:
		callbackHandlers.HandleCraftAlbumButton(ctx, bot, callback)

	// Craft agreement
	case buttons.DoCraftInlineButton.Data:
		callbackHandlers.HandleCraftAgreement(ctx, bot, callback)

	// Arrows
	default:
		// pattern for check if data is from left/right arrows
		arrowPattern := fmt.Sprintf(`^%s&\d+$`, buttons.LeftInlineButton.Data)
		arrowRe := regexp.MustCompile(arrowPattern)

		if arrowRe.MatchString(callback.Data) {
			callbackHandlers.HandleArrowButton(ctx, bot, callback)
		} else {
			logger.Log.Warnf("Unknown callback")
		}
	}
}
