package callbackHandlers

import (
	"hotPotBot/internal/context"
	"hotPotBot/internal/logger"
	"hotPotBot/internal/presentation/messages"
	"hotPotBot/internal/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleInitExchange(ctx *context.AppContext, bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	utils.AddUserPreviousRequest(ctx, callback.From.ID, callback.Data)

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, messages.WriteUsernameToExchangeTitle)
	_, err := bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending response <HandleInitExchange> | %v", err.Error())
	}
}
