package services

import (
	"hotPotBot/internal/context"
	"hotPotBot/internal/db/models"
)

type ExchangeService struct {
	Ctx *context.AppContext
}

type BuildExchangeInfo struct {
	Built bool
	FirstUser *models.User
	SecondUser *models.User
	FirstCard *models.Card
	SecondCard *models.Card
}

type BuildExchangeRequest struct {
	FromUserId int
	ToUserId int
	CardId int
}

func (service *ExchangeService) BuildExchange(req *BuildExchangeRequest) (*BuildExchangeInfo, error) {
	info := BuildExchangeInfo{
		Built: false,
	}
	// if FromUserId and ToUserId exists in one row CONTINUE exchange else INIT
	
	return &info, nil
}

type ExchangeRequest struct {
	FromUserId int
	ToUserId int
	CardId int
}

func (service *ExchangeService) Exchange(req *ExchangeRequest) {
	
}