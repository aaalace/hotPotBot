package context

import (
	"github.com/jmoiron/sqlx"
	"sync"
)

type AppContext struct {
	DB           *sqlx.DB
	UserRequests map[int64]string // временное решение
	Mu           sync.Mutex
}
