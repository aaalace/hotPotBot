package callbackHandlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hotPotBot/internal/consts"
	"hotPotBot/internal/context"
	"hotPotBot/internal/logger"
	buttons "hotPotBot/internal/presentation/buttons/callbackButtons"
	keyboards "hotPotBot/internal/presentation/keyboards/callbackKeyboards"
	"hotPotBot/internal/utils"
	"strconv"
	"strings"
)

// HandleArrowButton - кнопки пролистывания в карусели
func HandleArrowButton(ctx *context.AppContext, bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	index, err := strconv.Atoi(strings.Split(callback.Data, keyboards.DataDelimiter)[1])
	if err != nil {
		logger.Log.Errorf("Error in parsing arrow button <HandleArrowButton> | %v", err.Error())
		return
	}

	prevReq := utils.GetRmUserPreviousRequest(ctx, callback.From.ID)
	utils.AddUserPreviousRequest(ctx, callback.From.ID, prevReq)

	switch prevReq {
	// MyCards
	case buttons.AllCardsInlineButton.Data:
		HandleShowCardInList(ctx, bot, callback, consts.TypeAll, index, false)
	case buttons.ShopInlineButton.Data:
		HandleShowCardInList(ctx, bot, callback, consts.TypeSingle, index, false)
	case buttons.AlbumCardsInlineButton.Data:
		HandleShowCardInList(ctx, bot, callback, consts.TypeAlbum, index, false)
	case buttons.DuplicatesInlineButton.Data:
		HandleShowCardInList(ctx, bot, callback, consts.TypeAll, index, true)

	// Shop
	case buttons.ShopAllCardsInlineButton.Data:
		HandleShowCardInShop(ctx, bot, callback, consts.TypeAll, index)
	case buttons.ShopSingleCardsInlineButton.Data:
		HandleShowCardInShop(ctx, bot, callback, consts.TypeSingle, index)
	case buttons.ShopAlbumCardsInlineButton.Data:
		HandleShowCardInShop(ctx, bot, callback, consts.TypeAlbum, index)

	default:
		logger.Log.Errorf("Unexpected type in arrows handler")
	}
}
