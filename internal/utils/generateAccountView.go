package utils

import (
	"fmt"
)

func GenerateAccountView(username string, weight int) string {
	textView := fmt.Sprintf(
		"@%s\nFame: %v",
		username,
		weight)

	return textView
}
