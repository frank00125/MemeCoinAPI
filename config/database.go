package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DBPool interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	Ping(ctx context.Context) error
	Close()
}

var connectionPool DBPool

func InitDatabase() error {
	connectionString := getEnvVar("POSTGRESQL_DATABASE_URL")
	if connectionString == "" {
		panic("NO DATABASE URL PROVIDED, shutting down...")
	}

	ctx := context.Background()

	poolConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return fmt.Errorf("error parsing config: %w", err)
	}

	connectionPool, err = pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return fmt.Errorf("error creating pool: %w", err)
	}

	err = connectionPool.Ping(ctx)
	if err != nil {
		connectionPool.Close()
		return fmt.Errorf("error pinging database: %w", err)
	}

	return nil
}

func GetConnection() DBPool {
	return connectionPool
}
