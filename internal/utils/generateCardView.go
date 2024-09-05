package utils

import (
	"fmt"
	"hotPotBot/internal/db/models"
)

func GenerateCardView(card *models.Card, typename string, quantity int) string {
	textView := fmt.Sprintf(
		"%s\n\nТип: %s\nFame: %v\nКоличество: %v",
		card.Name,
		typename,
		card.Weight,
		quantity)

	return textView
}

func GenerateRandomCardView(card *models.Card, typename string) string {
	textView := fmt.Sprintf(
		"Поздравляем, тебе выпала карта -\n%s\n\nТип: %s\nFame: +%v\n",
		card.Name,
		typename,
		card.Weight)

	return textView
}
