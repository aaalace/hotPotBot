package models

import "time"

type Cooldown struct {
	Id       string
	User     User
	LastTake time.Time
}
