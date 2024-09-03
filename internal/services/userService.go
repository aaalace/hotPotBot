package services

import (
	"hotPotBot/internal/context"
	"hotPotBot/internal/db"
	"hotPotBot/internal/db/models"
)

type UserService struct {
	Ctx *context.AppContext
}

func (service *UserService) GetUserByTelegramId(telegramId int) (*models.User, error) {
	user := models.User{}
	err := service.Ctx.DB.Get(&user, db.SelectUserQuery, telegramId)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (service *UserService) AddUser(telegramId int) error {
	user := models.User{
		TelegramId: telegramId,
		Weight:     0,
	}
	_, err := service.Ctx.DB.NamedExec(db.AddUserQuery, &user)

	return err

}
