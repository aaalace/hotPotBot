package services

import (
	"hotPotBot/internal/context"
	"hotPotBot/internal/db"
	"hotPotBot/internal/db/models"
	"hotPotBot/internal/logger"
	"strings"
)

type UserService struct {
	Ctx *context.AppContext
}

func (service *UserService) GetUserByTelegramId(tgId int64) (*models.User, error) {
	user := models.User{}

	err := service.Ctx.DB.Get(&user, db.SelectUserQuery, tgId)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (service *UserService) GetUserByUsername(username string) (*models.User, error) {
	user := models.User{}

	username = strings.Replace(username, "@", "", 1)

	err := service.Ctx.DB.Get(&user, db.SelectUserByUsernameQuery, username)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (service *UserService) AddUser(tgId int64, tgUsername string) (*models.User, error) {
	user := models.User{
		TelegramId:       tgId,
		TelegramUsername: tgUsername,
	}

	_, err := service.Ctx.DB.NamedExec(db.AddUserQuery, &user)
	if err != nil {
		return nil, err
	}

	return &user, err
}

func (service *UserService) UpdCorrectUsername(tgId int64, tgUsername string) {
	user, err := service.GetUserByTelegramId(tgId)
	if err != nil {
		logger.Log.Errorf("UpdCorrectUsername error: %s", err.Error())
		return
	}

	if user.TelegramUsername == tgUsername {
		return
	}
	_, err = service.Ctx.DB.NamedExec(db.UpdCorrectUsername, map[string]interface{}{
		"telegram_username": tgUsername,
		"telegram_id":       tgId,
	})
}

func (service *UserService) CountUserWeight(userId int) (int, error) {
	weight := 0
	err := service.Ctx.DB.Get(&weight, db.CountUserWeight, userId)

	return weight, err
}
