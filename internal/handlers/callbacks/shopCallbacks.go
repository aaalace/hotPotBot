package callbackHandlers

import (
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

// HandleShopButton - Мой аккаунт
func HandleShopButton(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, messages.ShopTitle)
	msg.ReplyMarkup = keyboards.ShopMenuKeyboard
	_, err := bot.Send(msg)
	if err != nil {
		logger.Log.Errorf("Error sending response <HandleShopButton> | %v", err.Error())
	}
}

// HandleShopAllCardsButton - Все карты в магазине
func HandleShopAllCardsButton(ctx *context.AppContext, bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	utils.AddUserPreviousRequest(ctx, callback.From.ID, buttons.ShopAllCardsInlineButton.Data)
	HandleShowCardInShop(ctx, bot, callback, consts.TypeAll, consts.StartCarouselIndex)
}

// HandleShopSingleCardsButton - Single карты в магазине
func HandleShopSingleCardsButton(ctx *context.AppContext, bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	utils.AddUserPreviousRequest(ctx, callback.From.ID, buttons.ShopSingleCardsInlineButton.Data)
	HandleShowCardInShop(ctx, bot, callback, consts.TypeSingle, consts.StartCarouselIndex)
}

// HandleShopAlbumCardsButton - Album карты в магазине
func HandleShopAlbumCardsButton(ctx *context.AppContext, bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	utils.AddUserPreviousRequest(ctx, callback.From.ID, buttons.ShopAlbumCardsInlineButton.Data)
	HandleShowCardInShop(ctx, bot, callback, consts.TypeAlbum, consts.StartCarouselIndex)
}

func HandleShowCardInShop(
	ctx *context.AppContext,
	bot *tgbotapi.BotAPI,
	callback *tgbotapi.CallbackQuery,
	typeId int,
	index int,
) {

	msg := tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, callback.Message.MessageID)
	if _, err := bot.Request(msg); err != nil {
		log.Printf("Error deleting message <HandleShowCardInShop> | %v", err.Error())
	}

	// all cards with this type
	var cards []*models.Card
	cardService := services.CardService{Ctx: ctx}
	cards, err := cardService.GetCardsByType(typeId)
	if err != nil {
		logger.Log.Errorf("Error in getting cards <HandleShowCardInShop> | %v", err.Error())
		return
	}

	// if 0 cards of this type
	if len(cards) == 0 {
		logger.Log.Fatalf("No cards in shop wtf")
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, messages.InternalError)
		_, err = bot.Send(msg)
		if err != nil {
			logger.Log.Errorf("Error sending shop error response <HandleShowCardInShop> | %v", err.Error())
		}
		return
	}

	// index is in [1..len(cards)]
	index %= len(cards)
	card := cards[index]
	if index == 0 {
		index = len(cards)
	}

	typeName, err := cardService.GetTypeNameByTypeId(card.TypeId)
	if err != nil {
		logger.Log.Errorf("Error in getting typename | %v", err.Error())
		return
	}

	imageReader, err := s3.DownloadImageFromS3(ctx.S3Client, card.ImageUrl)

	photo := tgbotapi.NewPhoto(callback.Message.Chat.ID, tgbotapi.FileReader{
		Name:   card.ImageUrl,
		Reader: imageReader,
	})
	view := utils.GenerateShopCardView(card, typeName)
	photo.Caption = view
	photo.ReplyMarkup = keyboards.NewShopCarouselKeyboard(index, len(cards), index-1, index+1)
	_, err = bot.Send(photo)
	if err != nil {
		logger.Log.Errorf("Error sending response <HandleShowCardInShop> | %v", err.Error())
	}
}
