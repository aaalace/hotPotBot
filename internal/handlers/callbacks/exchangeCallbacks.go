package callbackHandlers

import (
	"hotPotBot/internal/consts"
	"hotPotBot/internal/context"
	"hotPotBot/internal/logger"
	"hotPotBot/internal/presentation/messages"
	"hotPotBot/internal/services"
	"hotPotBot/internal/utils"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleProcessExchange(ctx *context.AppContext, bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	utils.AddUserPreviousRequest(ctx, callback.From.ID, callback.Data)

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, messages.WriteUsernameToExchangeTitle)
	_, err := bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending response <HandleInitExchange> | %v", err.Error())
	}
}

func HandleAcceptExchange(ctx *context.AppContext, bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	userService := services.UserService{Ctx: ctx}

	// get current user
	currentUser, err := userService.GetUserByTelegramId(callback.From.ID)
	if err != nil {
		logger.Log.Errorf("Error getting currentUser <HandleAcceptExchange> | %v", err.Error())
		return
	}

	// get partner user
	data := strings.Split(callback.Data, consts.InlineDataDelimiter)
	partnerId, err := strconv.Atoi(data[1])
	if err != nil {
		logger.Log.Errorf("Error getting partnerId from callback <HandleAcceptExchange> | %v", err.Error())
		return
	}
	partnerUser, err := userService.GetUserById(partnerId)
	if err != nil {
		logger.Log.Errorf("Error getting partnerUser <HandleAcceptExchange> | %v", err.Error())
		return
	}

	// accept
	exchangeService := services.ExchangeService{Ctx: ctx}
	exchangeFinished, partnerCardId, err := exchangeService.AcceptExchange(currentUser.Id, partnerUser.Id)
	if err != nil {
		logger.Log.Errorf("Error accept in exchangeService <HandleAcceptExchange> | %v", err.Error())
		return
	}

	deleteMsg := tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, callback.Message.MessageID)
	if _, err = bot.Request(deleteMsg); err != nil {
		log.Printf("Error deleting message <HandleAcceptExchange> | %v", err.Error())
		return
	}

	// waiting partner response
	if !exchangeFinished {
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID,
			messages.WaitPartnerToAcceptExchange(partnerUser.TelegramUsername))
		_, err = bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("Error sending waiting response <HandleAcceptExchange> | %v", err.Error())
		}
		return
	}

	// getting partner's card
	cardService := services.CardService{Ctx: ctx}
	card, err := cardService.GetCardById(partnerCardId)
	if err != nil {
		logger.Log.Errorf("Error in getting card by id <HandleAcceptExchange> | %v", err.Error())
		return
	}

	// success partner response
	msg := tgbotapi.NewMessage(callback.Message.Chat.ID,
		messages.SuccessExchange(partnerUser.TelegramUsername, card))
	_, err = bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending success response <HandleAcceptExchange> | %v", err.Error())
	}
}

func HandleDeclineExchange(ctx *context.AppContext, bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	userService := services.UserService{Ctx: ctx}

	// getting partner user
	data := strings.Split(callback.Data, consts.InlineDataDelimiter)
	partnerId, err := strconv.Atoi(data[1])
	if err != nil {
		logger.Log.Errorf("Error getting partnerId from callback <HandleAcceptExchange> | %v", err.Error())
		return
	}
	partnerUser, err := userService.GetUserById(partnerId)
	if err != nil {
		logger.Log.Errorf("Error getting partnerUser <HandleAcceptExchange> | %v", err.Error())
		return
	}

	// getting current user
	currentUser, err := userService.GetUserByTelegramId(callback.From.ID)
	if err != nil {
		logger.Log.Errorf("Error getting currentUser <HandleAcceptExchange> | %v", err.Error())
		return
	}

	// decline
	exchangeService := services.ExchangeService{Ctx: ctx}
	err = exchangeService.DeclineExchange(currentUser.Id, partnerUser.Id)
	if err != nil {
		logger.Log.Errorf("Error in decline <HandleAcceptExchange> | %v", err.Error())
		return
	}

	deleteMsg := tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, callback.Message.MessageID)
	if _, err = bot.Request(deleteMsg); err != nil {
		log.Printf("Error deleting message <HandleInitExchange> | %v", err.Error())
		return
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, messages.DeclinedExchange)
	_, err = bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending response <HandleInitExchange> | %v", err.Error())
	}
}
