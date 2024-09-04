package services

import (
	"hotPotBot/internal/context"
	"hotPotBot/internal/db"
)

type CardService struct {
	Ctx *context.AppContext
}

func (service *CardService) GetTypeNameByTypeId(typeId int) (string, error) {
	typeName := ""
	err := service.Ctx.DB.Get(&typeName, db.SelectTypeNameById, typeId)

	return typeName, err
}
