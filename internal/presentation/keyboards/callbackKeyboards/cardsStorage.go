package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	buttons "hotPotBot/internal/presentation/buttons/callbackButtons"
)

var CardsStorageKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(buttons.AllCardsInlineButton.Title,
			buttons.AllCardsInlineButton.Data),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(buttons.DuplicatesInlineButton.Title,
			buttons.DuplicatesInlineButton.Data),
	),
)
