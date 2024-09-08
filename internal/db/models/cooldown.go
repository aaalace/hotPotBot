package models

import "time"

type Cooldown struct {
	Id         int       `db:"id"`
	UserId     int64     `db:"user_id"`
	NextAccept time.Time `db:"next_accept"`
}
