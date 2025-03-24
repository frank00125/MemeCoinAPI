package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

func NewDatabaseConnectionPool() (DatabaseConnectionPoolInterface, error) {
	connectionString := viper.GetString("POSTGRESQL_URL")
	fmt.Println("Connection String: ", connectionString)

	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return pool, nil
}
