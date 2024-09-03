package utils

import (
	"fmt"
	"hotPotBot/internal/db/models"
)

func GenerateAccountView(username string, user *models.User) string {
	textView := fmt.Sprintf("Имя пользователя: @%s\nFame: %v", username, user.Weight)
	return textView
}
