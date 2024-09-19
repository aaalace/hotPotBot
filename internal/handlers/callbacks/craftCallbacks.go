package callbackHandlers

import (
	"errors"
	"hotPotBot/internal/consts"
	"hotPotBot/internal/context"
	"hotPotBot/internal/db/models"
	"hotPotBot/internal/logger"
	buttons "hotPotBot/internal/presentation/buttons/callbackButtons"
	keyboards "hotPotBot/internal/presentation/keyboards/callbackKeyboards"
	"hotPotBot/internal/presentation/messages"
	"hotPotBot/internal/s3"
	"hotPotBot/internal/services"
	"hotPotBot/internal/utils"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleCraftButton - Крафт
func HandleCraftButton(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, messages.CraftTitle)
	msg.ReplyMarkup = keyboards.CraftMenuKeyboard
	_, err := bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending response <HandleCraftButton> | %v", err.Error())
	}
}

// HandleCraftAlbumButton - Крафт album
func HandleCraftAlbumButton(ctx *context.AppContext, bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	utils.AddUserPreviousRequest(ctx, callback.From.ID, buttons.CraftAlbumInlineButton.Data)

	// Delete craft menu
	delMsg := tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, callback.Message.MessageID)
	if _, err := bot.Request(delMsg); err != nil {
		log.Printf("Error deleting message <HandleCraftAlbumButton> | %v", err.Error())
	}

	view := utils.GenerateCraftAgreement(10, "Single", "Album")
	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, view)
	msg.ReplyMarkup = keyboards.CraftAgreementKeyboard
	_, err := bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending response <HandleCraftAlbum> | %v", err.Error())
	}
}

// HandleCraftAgreement - Согласие на крафт
func HandleCraftAgreement(ctx *context.AppContext, bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	prevReq := utils.GetRmUserPreviousRequest(ctx, callback.From.ID)

	// Delete agreement
	delMsg := tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, callback.Message.MessageID)
	if _, err := bot.Request(delMsg); err != nil {
		log.Printf("Error deleting message <HandleCraftAgreement> | %v", err.Error())
	}

	userService := services.UserService{Ctx: ctx}
	user, err := userService.GetUserByTelegramId(callback.From.ID)
	if err != nil {
		logger.Log.Errorf("Error in getting user <HandleCraftAgreement> | %v", err.Error())
		return
	}

	var craftedCard *models.Card
	var serviceErr error
	craftService := services.CraftService{Ctx: ctx}
	switch prevReq {
	case buttons.CraftAlbumInlineButton.Data:
		craftedCard, serviceErr = craftService.CraftCard(
			user.Id,
			consts.TypeSingle,
			consts.CraftAlbumPrice,
			consts.TypeAlbum)
	}

	if serviceErr != nil {
		if errors.As(serviceErr, &services.NotEnoughCardsForCraft{}) {
			msg := tgbotapi.NewMessage(callback.Message.Chat.ID, serviceErr.Error())
			_, err = bot.Send(msg)
			if err != nil {
				logger.Log.Errorf(
					"Error in sending response (NoCardsForCraft) <HandleCraftAgreement> | %v",
					serviceErr.Error(),
				)
			}
			return
		}
		return
	}
	if craftedCard == nil {
		logger.Log.Error("Error in craft (card - nil) <HandleCraftAgreement>")
		return
	}

	cardService := services.CardService{Ctx: ctx}
	typeName, err := cardService.GetTypeNameByTypeId(craftedCard.TypeId)
	if err != nil {
		logger.Log.Errorf("Error in getting typename <HandleCraftAgreement> | %v", err.Error())
		return
	}

	imageReader, err := s3.DownloadImageFromS3(ctx.S3Client, craftedCard.ImageUrl)

	photo := tgbotapi.NewPhoto(callback.Message.Chat.ID, tgbotapi.FileReader{
		Name:   craftedCard.ImageUrl,
		Reader: imageReader,
	})
	view := utils.GenerateCraftCardView(craftedCard, typeName)
	photo.Caption = view
	photo.ReplyMarkup = keyboards.AfterCraftKeyboard
	_, err = bot.Send(photo)
	if err != nil {
		logger.Log.Errorf("Error sending response <HandleCraftAgreement> | %v", err.Error())
	}
}
