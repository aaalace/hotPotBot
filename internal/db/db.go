package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"hotPotBot/internal/config"
	"hotPotBot/internal/logger"
)

func ConnectDatabase(cfg *config.Config) *sqlx.DB {
	logger.Log.Info("Initializing database")

	dsn := cfg.DatabasePath
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		logger.Log.Fatalf("Failed to setup database | %v", err.Error())
		return nil
	}
	return db
}
