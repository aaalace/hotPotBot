package keyboards

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"hotPotBot/internal/consts"
	buttons "hotPotBot/internal/presentation/buttons/callbackButtons"
)

func NewAcceptanceExchangeKeyboard(
	partnerId int,
) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.AcceptExchangeInlineButton.Title,
				fmt.Sprintf("%s%s%v",
					buttons.AcceptExchangeInlineButton.Data,
					consts.InlineDataDelimiter,
					partnerId),
			),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				buttons.DeclineExchangeInlineButton.Title,
				fmt.Sprintf("%s%s%v",
					buttons.DeclineExchangeInlineButton.Data,
					consts.InlineDataDelimiter,
					partnerId),
			),
		),
	)
}
