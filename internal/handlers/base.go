package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"hotPotBot/internal/context"
	"hotPotBot/internal/logger"
)

func HandleUpdate(ctx *context.AppContext, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message != nil {
		username := update.Message.From.UserName
		logger.Log.WithField("username", username).Info(
			"New message: " + update.Message.Text)
		if update.Message.IsCommand() {
			HandleCommand(ctx, bot, update.Message)
			return
		}
		HandleMessage(ctx, bot, update.Message)
	} else if update.CallbackQuery != nil {
		username := update.CallbackQuery.From.UserName
		logger.Log.WithField("username", username).Info(
			"New callback: " + update.CallbackQuery.Data)
		HandleCallback(ctx, bot, update.CallbackQuery)
	}
}
