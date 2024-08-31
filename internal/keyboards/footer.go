package keyboards

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

var Footer = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("ACCOUNT_BUTTON_TITLE"),
	),
)
