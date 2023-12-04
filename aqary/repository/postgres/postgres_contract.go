package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresRepository struct {
	DB      *pgxpool.Pool
	Queries *Queries
}

func NewClient(ctx context.Context, url string) (*PostgresRepository, error) {
	db, err := pgxpool.Connect(ctx, url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Ping(ctx)
	if err != nil {
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}

	return &PostgresRepository{
		DB:      db,
		Queries: New(db),
	}, nil
}
