package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/spf13/viper"
)

func NewDatabaseConnectionPool() (*sql.DB, error) {
	connectionString := viper.GetString("POSTGRESQL_URL")
	fmt.Println("Connection String: ", connectionString)

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	return db, nil
}
