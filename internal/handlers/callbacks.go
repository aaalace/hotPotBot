package handlers

import (
	"fmt"
	"hotPotBot/internal/context"
	callbackHandlers "hotPotBot/internal/handlers/callbacks"
	"hotPotBot/internal/logger"
	buttons "hotPotBot/internal/presentation/buttons/callbackButtons"
	"regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

		// pattern for check if data is from exchange this card
		exchangePattern := fmt.Sprintf(`^%s&\d+$`, buttons.ExchangeThisCardInlineButton.Data)
		exchangeRe := regexp.MustCompile(exchangePattern)

		// pattern for accept exchange
		exchangeAcceptPattern := fmt.Sprintf(`^%s&\d+$`, buttons.AcceptExchangeInlineButton.Data)
		excAccRe := regexp.MustCompile(exchangeAcceptPattern)

		// pattern for decline exchange
		exchangeDeclinePattern := fmt.Sprintf(`^%s&\d+$`, buttons.DeclineExchangeInlineButton.Data)
		excDecRe := regexp.MustCompile(exchangeDeclinePattern)

		if arrowRe.MatchString(callback.Data) {
			callbackHandlers.HandleArrowButton(ctx, bot, callback)
		} else if exchangeRe.MatchString(callback.Data) {
			callbackHandlers.HandleProcessExchange(ctx, bot, callback)
		} else if excAccRe.MatchString(callback.Data) {
			callbackHandlers.HandleAcceptExchange(ctx, bot, callback)
		} else if excDecRe.MatchString(callback.Data) {
			callbackHandlers.HandleDeclineExchange(ctx, bot, callback)
		} else {
			logger.Log.Warnf("Unknown callback")
		}
	}
}
