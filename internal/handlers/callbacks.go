package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"hotPotBot/internal/logger"
	buttons "hotPotBot/internal/presentation/buttons/callbackButtons"
)

func HandleCallback(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	switch callback.Data {
	// My Cards Menu
	case buttons.AllCardsInlineButton.Data:
	case buttons.DuplicatesInlineButton.Data:

	// Hot Pot Studio Menu
	case buttons.MyAccountInlineButton.Data:
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "@"+callback.From.UserName)
		_, err := bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("fe")
		}
	case buttons.FindUserInlineButton.Data:
	case buttons.ShopInlineButton.Data:
	case buttons.ExchangeInlineButton.Data:
	case buttons.CraftInlineButton.Data:
	case buttons.DiceInlineButton.Data:

	default:
		logger.Log.Warnf("Unknown callback: %s", callback.Message.Text)
	}
}
