package services

import (
	"hotPotBot/internal/context"
	"hotPotBot/internal/db"
	"hotPotBot/internal/db/models"
	"hotPotBot/internal/presentation/messages"
	"math/rand"
)

type CraftService struct {
	Ctx *context.AppContext
}

type NotEnoughCardsForCraft struct{}

func (NotEnoughCardsForCraft) Error() string {
	return messages.NoCardsForCraft
}

func (service *CraftService) CraftCard(userId, oldId, oldAmount, newId int) (*models.Card, error) {
	// get all user's duplicates
	var duplicates []*models.Card
	err := service.Ctx.DB.Select(&duplicates, db.SelectUserDuplicatesByType, userId, oldId)
	if err != nil {
		return nil, err
	}

	// count how many duplicates of oldId type user has
	canBeRemovedCounter := 0
	var duplicatesAmounts []int
	for _, card := range duplicates {
		currentRemoveAbility := 0
		err = service.Ctx.DB.Get(&currentRemoveAbility, db.SelectUserCardQuantity, userId, card.Id)
		currentRemoveAbility -= 1
		duplicatesAmounts = append(duplicatesAmounts, currentRemoveAbility) // already -1 !!!
		if err != nil {
			return nil, err
		}
		canBeRemovedCounter += currentRemoveAbility
	}

	// check if user has enough cards to craft
	if canBeRemovedCounter < oldAmount {
		return nil, NotEnoughCardsForCraft{}
	}

	// removing cards with oldType
	toRemove := oldAmount
	for i, card := range duplicates {
		currentRemove := duplicatesAmounts[i]
		if currentRemove > toRemove {
			currentRemove = toRemove
		}
		_, err = service.Ctx.DB.NamedExec(db.MinusUserCardQuantity, map[string]interface{}{
			"to_remove": currentRemove,
			"user_id":   userId,
			"card_id":   card.Id,
		})
		if err != nil {
			return nil, err
		}
		toRemove -= currentRemove
		if toRemove <= 0 {
			break
		}
	}

	// check for ids of cards with newType
	var availableIds []int
	err = service.Ctx.DB.Select(&availableIds, db.SelectAllCardsIdsByType, newId)
	if err != nil {
		return nil, err
	}

	// get random one
	randomCardId := availableIds[rand.Intn(len(availableIds))]

	// get card with this id
	var newCard models.Card
	err = service.Ctx.DB.Get(&newCard, db.SelectCardById, randomCardId)
	if err != nil {
		return nil, err
	}

	// give new card to user
	_, err = service.Ctx.DB.NamedExec(db.GiveUserCard, map[string]interface{}{
		"user_id": userId,
		"card_id": randomCardId,
	})
	if err != nil {
		return nil, err
	}

	return &newCard, nil
}
