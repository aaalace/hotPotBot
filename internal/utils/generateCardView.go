package utils

import (
	"fmt"
	"hotPotBot/internal/db/models"
	"hotPotBot/internal/presentation/messages"
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
		"%s\n%s\n\nТип: %s\nFame: +%v\n",
		messages.SuccessfulRandomCardDropTitle,
		card.Name,
		typename,
		card.Weight)

	return textView
}

func GenerateShopCardView(card *models.Card, typename string) string {
	textView := fmt.Sprintf(
		"%s\n\nТип: %s\nFame: %v\n\nЦена: %v₽\n%s",
		card.Name,
		typename,
		card.Weight,
		card.Price,
		messages.ForPurchaseWrite)

	return textView
}

func GenerateCraftCardView(card *models.Card, typename string) string {
	textView := fmt.Sprintf(
		"%s\n%s\n\nТип: %s\nFame: %v\n\n",
		messages.SuccessfulCraftTitle,
		card.Name,
		typename,
		card.Weight)

	return textView
}
