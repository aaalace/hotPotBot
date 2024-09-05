package messageHandlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hotPotBot/internal/context"
	"hotPotBot/internal/logger"
	"hotPotBot/internal/presentation/messages"
	"hotPotBot/internal/services"
	"hotPotBot/internal/utils"
)

// HandleOtherAccount - "@username"
func HandleOtherAccount(ctx *context.AppContext, bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, messages.UserNotFoundError)

	userService := services.UserService{Ctx: ctx}
	user, err := userService.GetUserByUsername(message.Text)
	if err != nil {
		logger.Log.Errorf("Error in getting user: %v", err.Error())
		_, err = bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("Error sending response <handleOtherAccount 1_2>: %v", err.Error())
		}
		return
	}
	weight, err := userService.CountUserWeight(user.Id)
	if err != nil {
		logger.Log.Errorf("Error in count user weight: %v", err.Error())
		_, err = bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("Error sending response <handleOtherAccount 2_2>: %v", err.Error())
		}
		return
	}

	accountView := utils.GenerateAccountView(user.TelegramUsername, weight)
	msg = tgbotapi.NewMessage(message.Chat.ID, accountView)
	_, err = bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending response <handleOtherAccount>: %v", err.Error())
	}
}
