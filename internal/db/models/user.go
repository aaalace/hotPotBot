package models

type User struct {
	Id         int   `db:"id"`
	TelegramId int64 `db:"telegram_id"`
}
