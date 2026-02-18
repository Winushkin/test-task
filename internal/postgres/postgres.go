// Package postgres contains connection to db
package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	User     string `env:"USER"`
	DBName   string `env:"NAME"`
	Password string `env:"PASSWORD"`
	Port     string `env:"PORT"`
	Host     string `env:"HOST"`
	MinConns string `env:"MIN_CONNS"`
	MaxConns string `env:"MAX_CONNS"`
}

func NewConn(ctx context.Context, cfg Config) (*pgxpool.Pool, error) {

	// urlExample := "postgres://username:password@localhost:5432/database_name?sslmode=disable&pool_min_conns=%d&pool_max_conns=%d"
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&pool_min_conns=%s&pool_max_conns=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.MinConns,
		cfg.MaxConns,
	)
	pgPool, err := pgxpool.New(ctx, connString)
	return pgPool, err
}
