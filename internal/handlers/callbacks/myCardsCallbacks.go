package callbackHandlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hotPotBot/consts"
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
)

// HandleAllCardsButton - Все карты
func HandleAllCardsButton(ctx *context.AppContext, bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	utils.AddUserPreviousRequest(ctx, callback.From.ID, buttons.AllCardsInlineButton.Data)
	HandleShowCardInList(ctx, bot, callback, consts.TypeAll, consts.StartCarouselIndex, false)
}

// HandleSingleCardsButton - Single карты
func HandleSingleCardsButton(ctx *context.AppContext, bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	utils.AddUserPreviousRequest(ctx, callback.From.ID, buttons.SingleCardsInlineButton.Data)
	HandleShowCardInList(ctx, bot, callback, consts.TypeSingle, consts.StartCarouselIndex, false)
}

// HandleAlbumCardsButton - Album карты
func HandleAlbumCardsButton(ctx *context.AppContext, bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	utils.AddUserPreviousRequest(ctx, callback.From.ID, buttons.AlbumCardsInlineButton.Data)
	HandleShowCardInList(ctx, bot, callback, consts.TypeAlbum, consts.StartCarouselIndex, false)
}

// HandleDuplicatesButton - Дубликаты
func HandleDuplicatesButton(ctx *context.AppContext, bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	utils.AddUserPreviousRequest(ctx, callback.From.ID, buttons.DuplicatesInlineButton.Data)
	HandleShowCardInList(ctx, bot, callback, consts.TypeAll, consts.StartCarouselIndex, true)
}

// HandleShowCardInList - стандартный показ карты
func HandleShowCardInList(
	ctx *context.AppContext,
	bot *tgbotapi.BotAPI,
	callback *tgbotapi.CallbackQuery,
	typeId int,
	index int,
	duplicates bool,
) {
	userService := services.UserService{Ctx: ctx}
	user, err := userService.GetUserByTelegramId(callback.From.ID)
	if err != nil {
		logger.Log.Errorf("Error in getting user <HandleShowCardInList> | %v", err.Error())
		return
	}

	// all user's cards with this type
	var cards []*models.Card
	cardService := services.CardService{Ctx: ctx}
	if duplicates {
		cards, err = cardService.GetUserDuplicates(user.Id)
	} else {
		cards, err = cardService.GetUserCardsWithType(user.Id, typeId)
	}
	if err != nil {
		logger.Log.Errorf("Error in getting user's cards <HandleShowCardInList> | %v", err.Error())
		return
	}

	// if 0 cards of this type
	if len(cards) == 0 {
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, messages.NoCardsError)
		_, err = bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("Error sending response <HandleShowCardInList> | %v", err.Error())
		}
		return
	}

	// index - [1..len(cards)]
	index %= len(cards)
	card := cards[index]
	if index == 0 {
		index = len(cards)
	}

	quantity, err := cardService.GetUserCardQuantity(user.Id, card.Id)
	if err != nil {
		logger.Log.Errorf("Error in getting quantity <HandleShowCardInList> | %v", err.Error())
		return
	}

	typeName, err := cardService.GetTypeNameByTypeId(card.TypeId)
	if err != nil {
		logger.Log.Errorf("Error in getting typename <HandleShowCardInList> | %v", err.Error())
		return
	}

	msg := tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, callback.Message.MessageID)
	if _, err = bot.Request(msg); err != nil {
		log.Printf("Error deleting message <HandleShowCardInList> | %v", err.Error())
	}

	imageReader, err := s3.DownloadImageFromS3(ctx.S3Client, card.ImageUrl)

	photo := tgbotapi.NewPhoto(callback.Message.Chat.ID, tgbotapi.FileReader{
		Name:   card.ImageUrl,
		Reader: imageReader,
	})
	view := utils.GenerateCardView(card, typeName, quantity)
	photo.Caption = view
	photo.ReplyMarkup = keyboards.NewCarouselKeyboard(index, len(cards), index-1, index+1)
	_, err = bot.Send(photo)
	if err != nil {
		logger.Log.Errorf("Error sending response <HandleShowCardInList> | %v", err.Error())
	}
}
