package keyboards

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	buttons "hotPotBot/internal/presentation/buttons/callbackButtons"
)

const DataDelimiter = "&"

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
				DataDelimiter, leftIndex)),

		tgbotapi.NewInlineKeyboardButtonData(
			fmt.Sprintf("(%v/%v)", cur, total),
			"counter_button_data"),

		tgbotapi.NewInlineKeyboardButtonData(
			buttons.RightInlineButton.Title,
			fmt.Sprintf("%s%s%v",
				buttons.RightInlineButton.Data,
				DataDelimiter,
				rightIndex),
		),
	)
}

func NewCarouselKeyboard(
	cur int,
	total int,
	leftIndex int,
	rightIndex int,
) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		getDefaultCarouselRow(cur, total, leftIndex, rightIndex),
	)
}

func NewMyCardsCarouselKeyboard(
	cur int,
	total int,
	leftIndex int,
	rightIndex int,
	userFromId int,
	cardId int,
) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		getDefaultCarouselRow(cur, total, leftIndex, rightIndex),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.ExchangeThisCardInlineButton.Title,
				fmt.Sprintf("%s%s%v%s%v",
					buttons.ExchangeThisCardInlineButton.Data,
					DataDelimiter,
					userFromId,
					DataDelimiter,
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
