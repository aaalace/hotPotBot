package callbackHandlers

import (
	"hotPotBot/internal/context"
	"hotPotBot/internal/logger"
	buttons "hotPotBot/internal/presentation/buttons/callbackButtons"
	"hotPotBot/internal/presentation/messages"
	"hotPotBot/internal/services"
	"hotPotBot/internal/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleMyAccount - Мой аккаунт
func HandleMyAccount(ctx *context.AppContext, bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	userService := services.UserService{Ctx: ctx}
	user, err := userService.GetUserByTelegramId(callback.From.ID)
	if err != nil {
		logger.Log.Errorf("Error in getting user <HandleMyAccount> | %v", err.Error())
		return
	}
	weight, err := userService.CountUserWeight(user.Id)
	if err != nil {
		logger.Log.Errorf("Error in count user weight <HandleMyAccount> | %v", err.Error())
		return
	}

	accountView := utils.GenerateMyAccountView(callback.From.UserName, weight)
	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, accountView)
	_, err = bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending response <HandleMyAccount> | %v", err.Error())
	}
}

// HandleOtherAccountButton - Найти пользователя
func HandleOtherAccountButton(ctx *context.AppContext, bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	utils.AddUserPreviousRequest(ctx, callback.From.ID, buttons.FindUserInlineButton.Data)

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, messages.OtherAccountPageTitle)
	_, err := bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending response <HandleMyAccount> | %v", err.Error())
	}
}
