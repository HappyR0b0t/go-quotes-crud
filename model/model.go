package model

import (
	"context"
)

type Quote struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Text   string `json:"text"`
}

type QuoteStorage interface {
	AddQuote(ctx context.Context, q Quote) (int, error)
	// GetAllQuotes(ctx context.Context) ([]Quote, error)
	// GetQuoteByID(ctx context.Context, id int) (Quote, error)
	// GetRandomQuote(ctx context.Context) (Quote, error)
	// GetQuotesByAuthor(ctx context.Context, author string) ([]Quote, error)
	// UpdateQuote(ctx context.Context, id int, q Quote) error
	// DeleteQuote(ctx context.Context, id int) error
}
