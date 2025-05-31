package storage

import (
	"sync"

	model "example.com/go-scout-ai-crud/model"
)

type QuotesStorage struct {
	mu     sync.Mutex
	quotes map[int]model.Quote
	nextID uint
}

func NewQuotesStorage() *QuotesStorage {
	return &QuotesStorage{
		quotes: make(map[int]model.Quote),
		nextID: 1,
	}
}

func (s *QuotesStorage) Create(quote model.Quote) model.Quote {
	s.mu.Lock()
	defer s.mu.Unlock()

	quote.ID = s.nextID
	s.quotes[int(s.nextID)] = quote
	s.nextID++
	return quote
}

func (s *QuotesStorage) List() []model.Quote {
	s.mu.Lock()
	defer s.mu.Unlock()

	quotes := make([]model.Quote, 0, len(s.quotes))
	for _, u := range s.quotes {
		quotes = append(quotes, u)
	}
	return quotes
}
