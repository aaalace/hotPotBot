package services

import (
	"database/sql"
	"errors"
	"hotPotBot/internal/context"
	"hotPotBot/internal/db"
	"hotPotBot/internal/db/models"
	"hotPotBot/internal/presentation/messages"
)

// ---------------------------------------------------

type ExchangeDeclined struct{}

func (err ExchangeDeclined) Error() string {
	return messages.ExchangeDeclinedByPartner
}

// ---------------------------------------------------

type NoThisCards struct{}

func (err NoThisCards) Error() string {
	return messages.NoThisCards
}

// ---------------------------------------------------

type ExchangeService struct {
	Ctx *context.AppContext
}

type BuildExchangeRequest struct {
	FromUserId int
	ToUserId   int
	CardId     int
}

// ---------------------------------------------------

func (service *ExchangeService) BuildExchange(req *BuildExchangeRequest) (bool, *models.Exchange, error) {
	// check if we need to continue exchange
	var initialized bool
	err := service.Ctx.DB.Get(&initialized, db.CheckExchangeInitialized, req.ToUserId, req.FromUserId)
	if err != nil {
		return false, nil, err
	}

	var exchange models.Exchange
	if !initialized {
		// init exchange
		err = service.Ctx.DB.QueryRowx(db.InitExchange,
			req.FromUserId, req.ToUserId, req.CardId).StructScan(&exchange)
		if err != nil {
			return false, nil, err
		}

		return false, &exchange, nil
	}
	// continue exchange
	err = service.Ctx.DB.QueryRowx(db.ContinueExchange,
		req.FromUserId, req.ToUserId, req.CardId).StructScan(&exchange)
	if err != nil {
		return false, nil, err
	}

	return true, &exchange, nil
}

// ---------------------------------------------------

// AcceptExchange returns (finished, partnerCardId, myCardId, err)
func (service *ExchangeService) AcceptExchange(myId, partnerId int) (bool, int, int, error) {
	// get exchange
	var exchange models.Exchange
	err := service.Ctx.DB.QueryRowx(db.AcceptExchange, myId, partnerId).StructScan(&exchange)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, 0, 0, ExchangeDeclined{}
		}
		return false, 0, 0, err
	}

	// ---- if one side haven't accepted yet ----

	if !exchange.UserInitAccept || !exchange.UserContinueAccept {
		return false, 0, 0, nil
	}

	// ---- if both sides accepted ----

	// check init user has card
	var userInitCards int
	err = service.Ctx.DB.Get(&userInitCards, db.SelectUserCardQuantity,
		exchange.UserInitId, exchange.CardInitId)
	if err != nil {
		return false, 0, 0, nil
	}
	// check continue user has card
	var userContinueCards int
	err = service.Ctx.DB.Get(&userContinueCards, db.SelectUserCardQuantity,
		exchange.UserContinueId, exchange.CardContinueId)
	if err != nil {
		return false, 0, 0, nil
	}

	if userInitCards < 1 || userContinueCards < 1 {
		return false, 0, 0, NoThisCards{}
	}

	// ---- if both sides accepted and got cards ----

	// ---------------------------------------------------

	// give init user card
	_, err = service.Ctx.DB.NamedExec(db.GiveUserCard, map[string]interface{}{
		"user_id": exchange.UserInitId,
		"card_id": exchange.CardContinueId,
	})
	if err != nil {
		return false, 0, 0, err
	}
	// delete init user card
	_, err = service.Ctx.DB.NamedExec(db.MinusUserCardQuantity, map[string]interface{}{
		"to_remove": 1,
		"user_id":   exchange.UserInitId,
		"card_id":   exchange.CardInitId,
	})
	if err != nil {
		return false, 0, 0, err
	}
	// check init user has card
	err = service.Ctx.DB.Get(&userInitCards, db.SelectUserCardQuantity,
		exchange.UserInitId, exchange.CardInitId)
	if err != nil {
		return false, 0, 0, nil
	}
	if userInitCards < 1 {
		_, err = service.Ctx.DB.Exec(db.DeleteUserCard, exchange.UserInitId, exchange.CardInitId)
		if err != nil {
			return false, 0, 0, nil
		}
	}

	// ---------------------------------------------------

	// give continue user card
	_, err = service.Ctx.DB.NamedExec(db.GiveUserCard, map[string]interface{}{
		"user_id": exchange.UserContinueId,
		"card_id": exchange.CardInitId,
	})
	if err != nil {
		return false, 0, 0, err
	}
	// delete continue user card
	_, err = service.Ctx.DB.NamedExec(db.MinusUserCardQuantity, map[string]interface{}{
		"to_remove": 1,
		"user_id":   exchange.UserContinueId,
		"card_id":   exchange.CardContinueId,
	})
	if err != nil {
		return false, 0, 0, err
	}
	// check continue user has card
	err = service.Ctx.DB.Get(&userContinueCards, db.SelectUserCardQuantity,
		exchange.UserContinueId, exchange.CardContinueId)
	if err != nil {
		return false, 0, 0, nil
	}
	if userContinueCards < 1 {
		_, err = service.Ctx.DB.Exec(db.DeleteUserCard, exchange.UserContinueId, exchange.CardContinueId)
		if err != nil {
			return false, 0, 0, nil
		}
	}

	// ---------------------------------------------------

	// close exchange
	_, err = service.Ctx.DB.Exec(db.DeleteExchange, myId, partnerId)
	if err != nil {
		return false, 0, 0, err
	}

	// partner and my card ids
	partnerCardId := int(exchange.CardInitId.Int32)
	myCardId := int(exchange.CardContinueId.Int32)
	if partnerId == int(exchange.UserContinueId.Int32) {
		partnerCardId = int(exchange.CardContinueId.Int32)
		myCardId = int(exchange.CardInitId.Int32)
	}

	return true, partnerCardId, myCardId, nil
}

// ---------------------------------------------------

func (service *ExchangeService) DeclineExchange(myId, partnerId int) error {
	_, err := service.Ctx.DB.Exec(db.DeleteExchange, myId, partnerId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	return nil
}
