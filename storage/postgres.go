package storage

import (
	"context"

	model "example.com/go-scout-ai-crud/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStorage struct {
	db *pgxpool.Pool
}

func NewPostgresStorage(connString string) (*PostgresStorage, error) {
	dbpool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	return &PostgresStorage{db: dbpool}, err
}

func (s *PostgresStorage) Create(quote model.Quote) model.Quote {
	err := s.db.QueryRow(
		context.Background(),
		"INSERT INTO quotes (author, text) VALUES ($1, $2) RETURNING id",
		quote.Author,
		quote.Text,
	).Scan(&quote.ID)
	if err != nil {
		panic(err)
	}
	return quote
}
