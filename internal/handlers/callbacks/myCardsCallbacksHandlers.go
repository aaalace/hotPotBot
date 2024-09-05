package callbackHandlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hotPotBot/internal/context"
	"hotPotBot/internal/logger"
	buttons "hotPotBot/internal/presentation/buttons/callbackButtons"
	keyboards "hotPotBot/internal/presentation/keyboards/callbackKeyboards"
	"hotPotBot/internal/presentation/messages"
	"hotPotBot/internal/services"
	"hotPotBot/internal/utils"
	"log"
	"strconv"
	"strings"
)

// HandleAllCardsButton - Все карты
func HandleAllCardsButton(ctx *context.AppContext, bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	utils.AddUserPreviousRequest(ctx, callback.From.ID, buttons.AllCardsInlineButton.Data)
	handleShowCardInList(ctx, bot, callback, 0, 1)
}

// HandleSingleCardsButton - Single карты
func HandleSingleCardsButton(ctx *context.AppContext, bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	utils.AddUserPreviousRequest(ctx, callback.From.ID, buttons.SingleCardsInlineButton.Data)
	handleShowCardInList(ctx, bot, callback, 1, 1)
}

// HandleAlbumCardsButton - Album карты
func HandleAlbumCardsButton(ctx *context.AppContext, bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	utils.AddUserPreviousRequest(ctx, callback.From.ID, buttons.AlbumCardsInlineButton.Data)
	handleShowCardInList(ctx, bot, callback, 2, 1)
}

// HandleDuplicatesButton - Дубликаты
func HandleDuplicatesButton(ctx *context.AppContext, bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	utils.AddUserPreviousRequest(ctx, callback.From.ID, buttons.DuplicatesInlineButton.Data)
	handleShowCardInList(ctx, bot, callback, 0, 1)
}

// HandleArrowButton - кнопки пролистывания в карусели
func HandleArrowButton(ctx *context.AppContext, bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	index, err := strconv.Atoi(strings.Split(callback.Data, keyboards.ArrowDataDelimiter)[1])
	if err != nil {
		logger.Log.Errorf("Error in parsing arrow button: %v", err.Error())
		return
	}

	prevReq := utils.GetRmUserPreviousRequest(ctx, callback.From.ID)
	utils.AddUserPreviousRequest(ctx, callback.From.ID, prevReq)

	switch prevReq {
	case buttons.AllCardsInlineButton.Data:
		handleShowCardInList(ctx, bot, callback, 0, index)
	case buttons.ShopInlineButton.Data:
		handleShowCardInList(ctx, bot, callback, 1, index)
	case buttons.AlbumCardsInlineButton.Data:
		handleShowCardInList(ctx, bot, callback, 2, index)
	case buttons.DuplicatesInlineButton.Data:
		handleShowCardInList(ctx, bot, callback, 0, index)
	default:
		logger.Log.Errorf("Unexpected type of list in arrows handler")
	}
}

func handleShowCardInList(
	ctx *context.AppContext,
	bot *tgbotapi.BotAPI,
	callback *tgbotapi.CallbackQuery,
	typeId int,
	index int,
) {
	userService := services.UserService{Ctx: ctx}
	user, err := userService.GetUserByTelegramId(callback.From.ID)
	if err != nil {
		logger.Log.Errorf("Error in getting user: %v", err.Error())
		return
	}

	cardService := services.CardService{Ctx: ctx}
	cards, err := cardService.GetUserCardsWithType(user.Id, typeId)
	if err != nil {
		logger.Log.Errorf("Error in getting user: %v", err.Error())
		return
	}

	// if 0 cards of this type
	if len(cards) == 0 {
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, messages.NoCardsMessage)
		_, err = bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("Error sending response <handleMyAccount>: %v", err.Error())
		}
		return
	}

	index %= len(cards)
	card := cards[index]
	if index == 0 {
		index = len(cards)
	}

	quantity, err := cardService.GetUserCardQuantity(user.Id, card.Id)
	if err != nil {
		logger.Log.Errorf("Error in getting quantity: %v", err)
		return
	}

	typeName, err := cardService.GetTypeNameByTypeId(card.TypeId)
	if err != nil {
		logger.Log.Errorf("Error in getting typename: %v", err)
		return
	}

	msg := tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, callback.Message.MessageID)
	if _, err := bot.Request(msg); err != nil {
		log.Printf("Error deleting message <handleShowCardInList>: %v", err)
	}

	photo := tgbotapi.NewPhoto(callback.Message.Chat.ID, tgbotapi.FilePath(card.ImageUrl))
	view := utils.GenerateCardView(card, typeName, quantity)
	photo.Caption = view
	photo.ReplyMarkup = keyboards.NewCarouselKeyboard(index, len(cards), index-1, index+1)
	_, err = bot.Send(photo)
	if err != nil {
		logger.Log.Errorf("Error sending response <handleShowCardInList>: %v", err.Error())
	}
}
