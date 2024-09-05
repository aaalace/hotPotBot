package handlers

import (
	"database/sql"
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hotPotBot/internal/context"
	"hotPotBot/internal/logger"
	"hotPotBot/internal/presentation/keyboards"
	"hotPotBot/internal/presentation/messages"
	"hotPotBot/internal/services"
)

func HandleCommand(ctx *context.AppContext, bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		handleStartCommand(ctx, bot, message)
	case "help":
		handleHelpCommand(bot, message)
	default:
		logger.Log.Warnf("Unknown command: %s", message.Command())
	}
}

func handleStartCommand(ctx *context.AppContext, bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, messages.StartPageTitle)
	msg.ReplyMarkup = keyboards.FooterKeyboard

	userService := services.UserService{Ctx: ctx}
	_, err := userService.GetUserByTelegramId(message.From.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			_, err = userService.AddUser(message.From.ID, message.From.UserName)
			if err != nil {
				logger.Log.Errorf("Failed to add user %v | %v", message.From.ID, err.Error())
				msg = tgbotapi.NewMessage(message.Chat.ID, messages.InternalError)
			}
		} else {
			logger.Log.Errorf("Error in getting user: %v", err)
		}
	}

	_, err = bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending response </start> | %v", err.Error())
	}
}

func handleHelpCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, messages.SupportContactText)
	_, err := bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending response <support contact> | %v", err.Error())
	}
}
