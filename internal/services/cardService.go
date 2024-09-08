package services

import (
	"hotPotBot/internal/context"
	"hotPotBot/internal/db"
	"hotPotBot/internal/db/models"
)

type CardService struct {
	Ctx *context.AppContext
}

type CardRequestParams struct {
	UserId     int
	CardTypeId int
	Index      int
}

func (service *CardService) GetCardsByType(typeId int) ([]*models.Card, error) {
	var cards []*models.Card

	err := service.Ctx.DB.Select(&cards, db.SelectCardsByType, typeId)
	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (service *CardService) GetUserCardsWithType(userId int, typeId int) ([]*models.Card, error) {
	var cards []*models.Card

	err := service.Ctx.DB.Select(&cards, db.SelectUserCardsByType, userId, typeId)
	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (service *CardService) GetUserDuplicates(userId int) ([]*models.Card, error) {
	var cards []*models.Card

	err := service.Ctx.DB.Select(&cards, db.SelectUserDuplicates, userId)
	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (service *CardService) GetTypeNameByTypeId(typeId int) (string, error) {
	typeName := ""
	err := service.Ctx.DB.Get(&typeName, db.SelectTypeNameById, typeId)

	return typeName, err
}

func (service *CardService) GetUserCardQuantity(userId int, cardId int) (int, error) {
	quantity := 0
	err := service.Ctx.DB.Get(&quantity, db.SelectUserCardQuantity, userId, cardId)

	return quantity, err
}
