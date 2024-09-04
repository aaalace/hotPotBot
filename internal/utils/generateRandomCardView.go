package utils

import (
	"fmt"
	"hotPotBot/internal/db/models"
)

func GenerateRandomCardView(card *models.Card, typename string) string {
	textView := fmt.Sprintf(
		"Поздравляем, тебе выпала карта -\n%s\n\nТип: %s\nFame: +%v\n",
		card.Name,
		typename,
		card.Weight)

	return textView
}
