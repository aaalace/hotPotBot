package keyboards

import (
	buttons "hotPotBot/internal/presentation/buttons/callbackButtons"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var CardsStorageKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(buttons.AllCardsInlineButton.Title,
			buttons.AllCardsInlineButton.Data),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(buttons.SingleCardsInlineButton.Title,
			buttons.SingleCardsInlineButton.Data),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(buttons.AlbumCardsInlineButton.Title,
			buttons.AlbumCardsInlineButton.Data),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(buttons.DuplicatesInlineButton.Title,
			buttons.DuplicatesInlineButton.Data),
	),
)
