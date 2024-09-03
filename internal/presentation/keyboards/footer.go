package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"hotPotBot/internal/presentation/buttons"
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
