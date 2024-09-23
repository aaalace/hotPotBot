package db

import (
	"hotPotBot/internal/config"
	"hotPotBot/internal/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDatabase(cfg *config.Config) *sqlx.DB {
	logger.Log.Info("Initializing database")

	dsn := cfg.DatabasePath
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		logger.Log.Fatalf("Failed to setup database | %v", err.Error())
		return nil
	}
	if db == nil {
		logger.Log.Fatalf("Failed to connect database")
		return nil
	}

	return db
}
