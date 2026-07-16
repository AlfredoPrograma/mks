package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/alfredoprograma/mks/internal/config"
	"github.com/alfredoprograma/mks/internal/queries"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrCannotConnectDB = errors.New("cannot connect to database")
)

func Connect(ctx context.Context, dbConfig config.DBConfig) (*queries.Queries, error) {
	connString := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)

	conn, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotConnectDB, err)
	}

	return queries.New(conn), nil
}
