package utils

import (
	"fmt"
	"hotPotBot/internal/presentation/messages"
)

func GenerateCraftAgreement(
	oldAmount int,
	oldName string,
	newName string,
) string {
	return fmt.Sprintf(
		"%s\n%v случайных %s ➡️ 1 случайный %s",
		messages.CraftAgreementTitle,
		oldAmount,
		oldName,
		newName,
	)
}
