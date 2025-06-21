package storage

import (
	"context"
	"errors"

	model "example.com/go-scout-ai-crud/model"
	"github.com/jackc/pgx/v5"
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

func (s *PostgresStorage) List(author string) ([]model.Quote, error) {
	var (
		rows pgx.Rows
		err  error
	)

	if author != "" {
		rows, err = s.db.Query(context.Background(), "SELECT id, author, text FROM quotes WHERE author=$1", author)
	} else {
		rows, err = s.db.Query(context.Background(), "SELECT id, author, text FROM quotes")
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []model.Quote
	for rows.Next() {
		var q model.Quote
		if err := rows.Scan(&q.ID, &q.Author, &q.Text); err != nil {
			return nil, err
		}
		quotes = append(quotes, q)
	}

	return quotes, nil
}

func (s *PostgresStorage) Delete(id int) error {
	cmd, err := s.db.Exec(context.Background(), "DELETE FROM quotes WHERE id=$1", id)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return errors.New("quote not found")
	}
	return nil
}

func (s *PostgresStorage) GetRandom() (model.Quote, error) {
	var q model.Quote
	err := s.db.QueryRow(
		context.Background(),
		"SELECT id, author, text FROM quotes ORDER BY RANDOM() LIMIT 1",
	).Scan(&q.ID, &q.Author, &q.Text)

	if err != nil {
		return model.Quote{}, err
	}
	return q, nil
}
