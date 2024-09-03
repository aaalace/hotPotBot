package models

type UserCard struct {
	UserId   int `db:"user_id"`
	CardId   int `db:"card_id"`
	Quantity int `db:"quantity"`
}
