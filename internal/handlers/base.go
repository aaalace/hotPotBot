package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hotPotBot/internal/context"
	"hotPotBot/internal/logger"
	"hotPotBot/internal/services"
)

func HandleUpdate(ctx *context.AppContext, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message != nil {
		id := update.Message.From.ID
		username := update.Message.From.UserName
		updCorrectUsernameMiddleware(ctx, id, username)

		logger.Log.WithField("username", username).Info("New message: " + update.Message.Text)

		if update.Message.IsCommand() {
			HandleCommand(ctx, bot, update.Message)
		} else {
			HandleMessage(ctx, bot, update.Message)
		}
	} else if update.CallbackQuery != nil {
		id := update.CallbackQuery.From.ID
		username := update.CallbackQuery.From.UserName
		updCorrectUsernameMiddleware(ctx, id, username)

		logger.Log.WithField("username", username).Info("New callback: " + update.CallbackQuery.Data)

		HandleCallback(ctx, bot, update.CallbackQuery)
	}
}

func updCorrectUsernameMiddleware(ctx *context.AppContext, tgId int64, tgUsername string) {
	userService := services.UserService{Ctx: ctx}
	userService.UpdCorrectUsername(tgId, tgUsername)
}
