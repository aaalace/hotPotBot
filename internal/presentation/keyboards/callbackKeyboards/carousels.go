package keyboards

import (
	"fmt"
	"hotPotBot/internal/consts"
	buttons "hotPotBot/internal/presentation/buttons/callbackButtons"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewMyCardsCarouselKeyboard(
	cur int,
	total int,
	leftIndex int,
	rightIndex int,
	cardId int,
) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		getDefaultCarouselRow(cur, total, leftIndex, rightIndex),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.ExchangeThisCardInlineButton.Title,
				fmt.Sprintf("%s%s%v",
					buttons.ExchangeThisCardInlineButton.Data,
					consts.InlineDataDelimiter,
					cardId),
			),
		),
	)
}

func NewShopCarouselKeyboard(
	cur int,
	total int,
	leftIndex int,
	rightIndex int,
) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		getDefaultCarouselRow(cur, total, leftIndex, rightIndex),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("👈 Назад",
				buttons.ShopInlineButton.Data),
		),
	)
}

func getDefaultCarouselRow(
	cur int,
	total int,
	leftIndex int,
	rightIndex int,
) []tgbotapi.InlineKeyboardButton {

	if total == 1 {
		return tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				fmt.Sprintf("(%v/%v)", cur, total),
				"counter_button_data"),
		)
	}

	return tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			buttons.LeftInlineButton.Title,
			fmt.Sprintf("%s%s%v",
				buttons.LeftInlineButton.Data,
				consts.InlineDataDelimiter,
				leftIndex)),

		tgbotapi.NewInlineKeyboardButtonData(
			fmt.Sprintf("(%v/%v)", cur, total),
			"not_used"),

		tgbotapi.NewInlineKeyboardButtonData(
			buttons.RightInlineButton.Title,
			fmt.Sprintf("%s%s%v",
				buttons.RightInlineButton.Data,
				consts.InlineDataDelimiter,
				rightIndex),
		),
	)
}
