package handlers

import (
	"database/sql"
	"errors"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"hotPotBot/internal/context"
	"hotPotBot/internal/logger"
	buttons "hotPotBot/internal/presentation/buttons/callbackButtons"
	"hotPotBot/internal/services"
	"hotPotBot/internal/utils"
)

func HandleCallback(ctx *context.AppContext, bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	switch callback.Data {
	// My Cards Menu
	case buttons.AllCardsInlineButton.Data:
	case buttons.DuplicatesInlineButton.Data:

	// Hot Pot Studio Menu
	case buttons.MyAccountInlineButton.Data:
		handleMyAccount(ctx, bot, callback)
	case buttons.FindUserInlineButton.Data:
	case buttons.ShopInlineButton.Data:
	case buttons.ExchangeInlineButton.Data:
	case buttons.CraftInlineButton.Data:
	case buttons.DiceInlineButton.Data:

	default:
		logger.Log.Warnf("Unknown callback: %s", callback.Message.Text)
	}
}

func handleMyAccount(ctx *context.AppContext, bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	userService := services.UserService{Ctx: ctx}
	user, err := userService.GetUserByTelegramId(callback.From.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Log.Warnf("User do not exists: %v", err)
		} else {
			logger.Log.Errorf("Error in getting user: %v", err)
		}
		return
	}

	accountView := utils.GenerateAccountView(callback.From.UserName, user)
	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, accountView)

	_, err = bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("fe")
	}
}
