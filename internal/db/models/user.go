package models

type User struct {
	Id         string
	TelegramId int64
	Weight     int64
	Cards      []Card
}
