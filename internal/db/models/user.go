package models

type User struct {
	Id         string `db:"id"`
	TelegramId int    `db:"telegram_id"`
	Weight     int    `db:"weight"`
}
