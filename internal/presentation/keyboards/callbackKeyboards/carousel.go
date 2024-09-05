package keyboards

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	buttons "hotPotBot/internal/presentation/buttons/callbackButtons"
)

const ArrowDataDelimiter = "&"

func NewCarouselKeyboard(
	cur int,
	total int,
	leftIndex int,
	rightIndex int,
) tgbotapi.InlineKeyboardMarkup {

	if total == 1 {
		return tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(
					fmt.Sprintf("(%v/%v)", cur, total),
					"counter_button_data"),
			),
		)
	}

	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.LeftInlineButton.Title,
				fmt.Sprintf("%s%s%v",
					buttons.LeftInlineButton.Data,
					ArrowDataDelimiter, leftIndex)),

			tgbotapi.NewInlineKeyboardButtonData(
				fmt.Sprintf("(%v/%v)", cur, total),
				"counter_button_data"),

			tgbotapi.NewInlineKeyboardButtonData(
				buttons.RightInlineButton.Title,
				fmt.Sprintf("%s%s%v",
					buttons.RightInlineButton.Data,
					ArrowDataDelimiter,
					rightIndex),
			),
		),
	)
}
