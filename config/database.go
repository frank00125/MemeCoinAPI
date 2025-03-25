package config

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/spf13/viper"
)

func NewDatabaseConnectionPool() (*sql.DB, error) {
	connectionString := viper.GetString("POSTGRESQL_URL")

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	return db, nil
}
