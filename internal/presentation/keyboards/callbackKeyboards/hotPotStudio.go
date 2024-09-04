package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	buttons "hotPotBot/internal/presentation/buttons/callbackButtons"
)

var HotPotStudioKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(buttons.MyAccountInlineButton.Title,
			buttons.MyAccountInlineButton.Data),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(buttons.FindUserInlineButton.Title,
			buttons.FindUserInlineButton.Data),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(buttons.ShopInlineButton.Title,
			buttons.ShopInlineButton.Data),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(buttons.CraftInlineButton.Title,
			buttons.CraftInlineButton.Data),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(buttons.ExchangeInlineButton.Title,
			buttons.ExchangeInlineButton.Data),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(buttons.DiceInlineButton.Title,
			buttons.DiceInlineButton.Data),
	),
)
