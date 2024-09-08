package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	buttons "hotPotBot/internal/presentation/buttons/callbackButtons"
)

var ShopMenuKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(buttons.ShopAllCardsInlineButton.Title,
			buttons.ShopAllCardsInlineButton.Data),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(buttons.ShopSingleCardsInlineButton.Title,
			buttons.ShopSingleCardsInlineButton.Data),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(buttons.ShopAlbumCardsInlineButton.Title,
			buttons.ShopAlbumCardsInlineButton.Data),
	),
)
