package context

import "github.com/jmoiron/sqlx"

type AppContext struct {
	DB *sqlx.DB
}
