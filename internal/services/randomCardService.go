package services

import (
	"database/sql"
	"errors"
	"hotPotBot/internal/consts"
	"hotPotBot/internal/context"
	"hotPotBot/internal/db"
	"hotPotBot/internal/db/models"
	"hotPotBot/internal/presentation/messages"
	"math/rand"
	"time"
)

type RandomCardService struct {
	Ctx *context.AppContext
}

type NotEnoughTime struct {
	TimeLeft time.Duration
}

func (err NotEnoughTime) Error() string {
	return messages.SmallCooldownError + err.TimeLeft.Round(time.Second).String()
}

func (service *RandomCardService) GetRandomCard(userId int) (*models.Card, error) {
	// get cooldown
	cooldown := models.Cooldown{}
	err := service.Ctx.DB.Get(&cooldown, db.GetCooldown, userId)
	// unknown error case
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	nextAccept := time.Now().Add(consts.FixedCooldown)
	if err == nil {
		timeLeft := cooldown.NextAccept.Sub(time.Now())

		// check for ability to take card
		if timeLeft > 0 {
			notEnoughTime := NotEnoughTime{TimeLeft: timeLeft}
			return nil, notEnoughTime
		}

		_, err = service.Ctx.DB.Exec(db.UpdateCooldown, nextAccept, userId)
	} else {
		_, err = service.Ctx.DB.NamedExec(db.AddCooldown, map[string]interface{}{
			"user_id":     userId,
			"next_accept": nextAccept,
		})
	}

	// get randomCardId
	var ids []int
	err = service.Ctx.DB.Select(&ids, db.SelectAllCardsIds)
	if err != nil {
		return nil, err
	}
	randomCardId := ids[rand.Intn(len(ids))]

	// get card with this id
	card := models.Card{}
	err = service.Ctx.DB.Get(&card, db.SelectCardById, randomCardId)
	if err != nil {
		return nil, err
	}

	// give card to user
	_, err = service.Ctx.DB.NamedExec(db.GiveUserCard, map[string]interface{}{
		"user_id": userId,
		"card_id": randomCardId,
	})
	if err != nil {
		return nil, err
	}

	return &card, err
}
