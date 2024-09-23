package services

import (
	"hotPotBot/internal/context"
	"hotPotBot/internal/db"
	"hotPotBot/internal/db/models"
)

type ExchangeService struct {
	Ctx *context.AppContext
}

type BuildExchangeRequest struct {
	FromUserId int
	ToUserId   int
	CardId     int
}

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

// AcceptExchange returns (finished, partnerCardId, err)
func (service *ExchangeService) AcceptExchange(myId, partnerId int) (bool, int, error) {

	// если уже удален, значит партнер отменил, надо сообщить в кастомной ошибке и вернуться

	// ставим accept = true в зависимости от myId

	// если второй accept пока false
	//if ... {
	//	return false, 0, nil
	//}

	// удалить exchange

	// если второй accept уже true
	return true, 10000, nil
}

func (service *ExchangeService) DeclineExchange(myId, partnerId int) error {

	// удалить exchange (если уже удален, то пох)

	return nil
}
