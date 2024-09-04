package services

import (
	"hotPotBot/internal/context"
	"hotPotBot/internal/db"
	"hotPotBot/internal/db/models"
)

type UserService struct {
	Ctx *context.AppContext
}

func (service *UserService) GetUserByTelegramId(telegramId int64) (*models.User, error) {
	user := models.User{}

	err := service.Ctx.DB.Get(&user, db.SelectUserQuery, telegramId)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (service *UserService) AddUser(telegramId int64) (*models.User, error) {
	user := models.User{
		TelegramId: telegramId,
	}

	_, err := service.Ctx.DB.NamedExec(db.AddUserQuery, &user)
	if err != nil {
		return nil, err
	}

	return &user, err
}

func (service *UserService) CountUserWeight(userId int) (int, error) {
	weight := 0
	err := service.Ctx.DB.Get(&weight, db.CountUserWeight, userId)

	return weight, err
}
