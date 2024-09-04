package services

import (
	"hotPotBot/internal/context"
	"hotPotBot/internal/db"
	"hotPotBot/internal/db/models"
	"math/rand"
)

type RandomCardService struct {
	Ctx *context.AppContext
}

func (service *RandomCardService) GetRandomCard(userId int) (*models.Card, error) {
	var ids []int
	err := service.Ctx.DB.Select(&ids, db.SelectAllCardsIds)
	if err != nil {
		return nil, err
	}

	randomCardId := ids[rand.Intn(len(ids))]

	// get random card
	card := models.Card{}
	err = service.Ctx.DB.Get(&card, db.SelectCardById, randomCardId)
	if err != nil {
		return nil, err
	}

	// give card to user
	_, err = service.Ctx.DB.NamedExec(db.GiveUserRandomCard, map[string]interface{}{
		"user_id": userId,
		"card_id": randomCardId,
	})
	if err != nil {
		return nil, err
	}

	return &card, err
}
