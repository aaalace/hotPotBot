package messageHandlers

import (
	"hotPotBot/internal/consts"
	"hotPotBot/internal/context"
	"hotPotBot/internal/logger"
	"hotPotBot/internal/presentation/messages"
	"hotPotBot/internal/services"
	"hotPotBot/internal/utils"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleOtherAccount - "@username"
func HandleOtherAccount(ctx *context.AppContext, bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msgUserNotFound := tgbotapi.NewMessage(message.Chat.ID, messages.UserNotFoundError)

	userService := services.UserService{Ctx: ctx}
	user, err := userService.GetUserByUsername(message.Text)
	if err != nil {
		logger.Log.Errorf("Error in getting user <HandleOtherAccount> | %v", err.Error())
		_, err = bot.Send(msgUserNotFound)
		if err != nil {
			logger.Log.Errorf("Error sending response <HandleOtherAccount> | %v", err.Error())
		}
		return
	}
	weight, err := userService.CountUserWeight(user.Id)
	if err != nil {
		logger.Log.Errorf("Error in count user weight <HandleOtherAccount> | %v", err.Error())
		return
	}

	accountView := utils.GenerateAccountView(user.TelegramUsername, weight)
	msg := tgbotapi.NewMessage(message.Chat.ID, accountView)
	_, err = bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending response <HandleOtherAccount>: %v", err.Error())
	}
}

// HandleExchangeToAccount - "@username"
func HandleExchangeToAccount(ctx *context.AppContext, bot *tgbotapi.BotAPI, message *tgbotapi.Message, prevReq string) {
	msgUserNotFound := tgbotapi.NewMessage(message.Chat.ID, messages.UserNotFoundError)
	
	userToUsername := message.Text
	userFromId := message.From.ID

	// getting both USERS
	userService := services.UserService{Ctx: ctx}
	userTo, err := userService.GetUserByUsername(userToUsername)
	if err != nil {
		logger.Log.Errorf("Error in getting user <HandleExchangeToAccount> | %v", err.Error())
		_, err = bot.Send(msgUserNotFound)
		if err != nil {
			logger.Log.Errorf("Error sending response <HandleExchangeToAccount> | %v", err.Error())
		}
		return
	}
	userFrom, err := userService.GetUserByTelegramId(userFromId)
	if (err != nil) {
		logger.Log.Errorf("Error getting current user <HandleExchangeToAccount> | %v", err.Error())
		return
	}
	
	// getting current user's CARD
	data := strings.Split(prevReq, consts.InlineDataDelimiter)
	cardId, err := strconv.Atoi(data[1])
	if (err != nil) {
		logger.Log.Errorf("Error in parse <HandleExchangeToAccount> | %v", err.Error())
		return
	}
	cardService := services.CardService{Ctx: ctx}
	card, err := cardService.GetCardById(cardId)
	if (err != nil) {
		logger.Log.Errorf("Error in getting card by id <HandleExchangeToAccount> | %v", err.Error())
		return
	}

	// init/continue exchange
	exchangeService := services.ExchangeService{Ctx: ctx}
	exchangeInfo, err := exchangeService.BuildExchange(&services.BuildExchangeRequest{
		ToUserId:   userTo.Id,
		FromUserId: userFrom.Id,
		CardId: card.Id,
	})
	if (err != nil) {
		logger.Log.Errorf("Error in building exchange <HandleExchangeToAccount> | %v", err.Error())
		return
	}
	if exchangeInfo.Built {
		// TODO send agreements
	}
	
	// send to USER who CONTINUE exchange
	msgToOtherUser := tgbotapi.NewMessage(userTo.TelegramId, messages.WantToCotinueExchange("@"+userFrom.TelegramUsername, card.Name))
	_, err = bot.Send(msgToOtherUser)
	if err != nil {
		logger.Log.Errorf("Error sending notification <HandleExchangeToAccount>: %v", err.Error())
	}
	
	// send to USER who INIT exchange
	msg := tgbotapi.NewMessage(userFrom.TelegramId, messages.SuccessfulExchangeInit("@"+userTo.TelegramUsername))
	_, err = bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending response <HandleExchangeToAccount>: %v", err.Error())
	}
}
