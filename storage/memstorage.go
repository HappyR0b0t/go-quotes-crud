package storage

import (
	"context"
	"errors"
	"math/rand"
	"sync"

	model "example.com/go-scout-ai-crud/model"
)

type MemStorage struct {
	mu     sync.Mutex
	quotes map[int]model.Quote
	nextID int
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		quotes: make(map[int]model.Quote),
		nextID: 1,
	}
}

func (m *MemStorage) AddQuote(ctx context.Context, q model.Quote) (int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	q.ID = m.nextID
	m.nextID++
	m.quotes[q.ID] = q
	return q.ID, nil
}

// func (m *MemStorage) GetAllQuotes(ctx context.Context) ([]model.Quote, error) {
// 	m.mu.Lock()
// 	defer m.mu.Unlock()

// 	var result []model.Quote
// 	for _, q := range m.quotes {
// 		result = append(result, q)
// 	}
// 	return result, nil
// }

// func (m *MemStorage) GetQuoteByID(ctx context.Context, id int) (model.Quote, error) {
// 	m.mu.Lock()
// 	defer m.mu.Unlock()

// 	if _, ok := m.quotes[id]; !ok {
// 		return model.Quote{}, errors.New("id not found")
// 	}
// 	quote := m.quotes[id]
// 	return quote, nil
// }

func (m *MemStorage) GetRandomQuote(ctx context.Context) (model.Quote, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	ids := make([]int, 0, len(m.quotes))
	for k := range m.quotes {
		ids = append(ids, k)
	}
	randomIndex := ids[rand.Intn(len(ids))]
	quote, exists := m.quotes[randomIndex]

	if !exists {
		return model.Quote{}, errors.New("quote not found")
	}
	return quote, nil
}

func (m *MemStorage) GetQuotesByAuthor(ctx context.Context, author string) ([]model.Quote, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	var res []model.Quote

	for _, q := range m.quotes {
		if q.Author == author {
			res = append(res, q)
		}
	}
	if len(res) == 0 {
		return []model.Quote{}, errors.New("quote not found")
	}
	return res, nil
}

func (m *MemStorage) DeleteQuote(ctx context.Context, id int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.quotes[id]; !ok {
		return errors.New("quote not found")
	}
	delete(m.quotes, id)
	return nil
}
