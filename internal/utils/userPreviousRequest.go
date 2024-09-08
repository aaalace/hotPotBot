package utils

import (
	"hotPotBot/internal/context"
)

func AddUserPreviousRequest(ctx *context.AppContext, telegramId int64, request string) {
	ctx.Mu.Lock()
	ctx.UserRequests[telegramId] = request
	ctx.Mu.Unlock()
}

func GetRmUserPreviousRequest(ctx *context.AppContext, telegramId int64) string {
	ctx.Mu.Lock()
	response := ctx.UserRequests[telegramId]
	delete(ctx.UserRequests, telegramId)
	ctx.Mu.Unlock()

	return response
}
