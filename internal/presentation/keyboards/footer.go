package keyboards

import (
	"hotPotBot/internal/presentation/buttons"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var FooterKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(buttons.GetRandomCardButton),
		tgbotapi.NewKeyboardButton(buttons.CardsStorageButton),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(buttons.HotPotStudioButton),
		tgbotapi.NewKeyboardButton(buttons.TutorialButton),
	),
)
