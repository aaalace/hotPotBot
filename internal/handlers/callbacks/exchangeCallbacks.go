package callbackHandlers

import (
	"errors"
	"hotPotBot/internal/consts"
	"hotPotBot/internal/context"
	"hotPotBot/internal/db/models"
	"hotPotBot/internal/logger"
	"hotPotBot/internal/presentation/messages"
	"hotPotBot/internal/s3"
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
	exchangeFinished, partnerCardId, myCardId, err :=
		exchangeService.AcceptExchange(currentUser.Id, partnerUser.Id)
	if err != nil {
		if errors.As(err, &services.ExchangeDeclined{}) || errors.As(err, &services.NoThisCards{}) {
			if errors.As(err, &services.NoThisCards{}) {
				_ = exchangeService.DeclineExchange(currentUser.Id, partnerUser.Id)
			}

			msg := tgbotapi.NewMessage(callback.Message.Chat.ID, err.Error())
			_, err = bot.Send(msg)
			if err != nil {
				logger.Log.Errorf("Error sending ExchangeDeclined <HandleAcceptExchange> | %v", err.Error())
			}
			return
		}
		logger.Log.Errorf("Error accept in exchangeService <HandleAcceptExchange> | %v", err.Error())
		return
	}

	// delete previous message after service processes
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

	cardService := services.CardService{Ctx: ctx}
	// getting partner's card
	partnerCard, err := cardService.GetCardById(partnerCardId)
	if err != nil {
		logger.Log.Errorf("Error in getting card by id <HandleAcceptExchange> | %v", err.Error())
		return
	}
	partnerCardTypeName, err := cardService.GetTypeNameByTypeId(partnerCard.TypeId)
	if err != nil {
		logger.Log.Errorf("Error in getting typename <HandleAcceptExchange> | %v", err.Error())
		return
	}
	// getting my card
	myCard, err := cardService.GetCardById(myCardId)
	if err != nil {
		logger.Log.Errorf("Error in getting card by id <HandleAcceptExchange> | %v", err.Error())
		return
	}
	myCardTypeName, err := cardService.GetTypeNameByTypeId(myCard.TypeId)
	if err != nil {
		logger.Log.Errorf("Error in getting typename <HandleAcceptExchange> | %v", err.Error())
		return
	}

	// response sender
	sendResponse := func(chatId int64, newCard *models.Card, typename, partnerUsername string) {
		imageReader, err := s3.DownloadImageFromS3(ctx.S3Client, newCard.ImageUrl)
		if err != nil {
			logger.Log.Errorf("Error in image reader <HandleAcceptExchange> | %v", err.Error())
			return
		}
		photo := tgbotapi.NewPhoto(chatId, tgbotapi.FileReader{
			Name:   newCard.ImageUrl,
			Reader: imageReader,
		})
		view := utils.GenerateExchangeCardView(newCard, typename, partnerUsername)
		photo.Caption = view
		_, err = bot.Send(photo)
		if err != nil {
			logger.Log.Errorf("Error sending success response <HandleAcceptExchange> | %v", err.Error())
		}
	}

	// send to me
	sendResponse(callback.Message.Chat.ID, partnerCard, partnerCardTypeName, partnerUser.TelegramUsername)
	// send to partner
	sendResponse(partnerUser.TelegramId, myCard, myCardTypeName, currentUser.TelegramUsername)
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
