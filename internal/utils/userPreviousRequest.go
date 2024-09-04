package utils

import (
	"hotPotBot/internal/context"
)

func AddUserPreviousRequest(ctx *context.AppContext, telegramId int64, request string) {
	ctx.Mu.Lock()
	ctx.UserRequests[telegramId] = request
	ctx.Mu.Unlock()
}

func RemoveUserPreviousRequest(ctx *context.AppContext, telegramId int64) {
	ctx.Mu.Lock()
	delete(ctx.UserRequests, telegramId)
	ctx.Mu.Unlock()
}
