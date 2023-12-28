package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(ctx context.Context, cfg Config) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)

	connConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection config: %s", err)
	}

	pool, err = pgxpool.NewWithConfig(ctx, connConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection: %s", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("database doesn't respond: %s", err)
	}

	return pool, nil
}
