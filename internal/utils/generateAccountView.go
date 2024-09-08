package utils

import (
	"fmt"
	"hotPotBot/internal/presentation/messages"
)

func GenerateMyAccountView(username string, weight int) string {
	textView := fmt.Sprintf(
		"%s\n\n@%s\nFame: %v",
		messages.MyAccountPageTitle,
		username,
		weight)

	return textView
}

func GenerateAccountView(username string, weight int) string {
	textView := fmt.Sprintf(
		"@%s\nFame: %v",
		username,
		weight)

	return textView
}
