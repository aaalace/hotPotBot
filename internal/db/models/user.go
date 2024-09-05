package models

type User struct {
	Id               int    `db:"id"`
	TelegramId       int64  `db:"telegram_id"`
	TelegramUsername string `db:"telegram_username"`
}
