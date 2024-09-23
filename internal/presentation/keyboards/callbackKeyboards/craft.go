package keyboards

import (
	buttons "hotPotBot/internal/presentation/buttons/callbackButtons"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var AfterCraftKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"👈 Назад",
			buttons.CraftInlineButton.Data),
	),
)

var CraftAgreementKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			buttons.DoCraftInlineButton.Title,
			buttons.DoCraftInlineButton.Data),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			buttons.CancelCraftInlineButton.Title,
			buttons.CraftInlineButton.Data),
	),
)

var CraftMenuKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(buttons.CraftAlbumInlineButton.Title,
			buttons.CraftAlbumInlineButton.Data),
	),
)
