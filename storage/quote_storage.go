package storage

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"

	model "example.com/go-scout-ai-crud/model"
)

type QuotesStorage struct {
	mu     sync.Mutex
	quotes map[int]model.Quote
	nextID int
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
	s.quotes[s.nextID] = quote
	s.nextID++
	return quote
}

func (s *QuotesStorage) List(author string) []model.Quote {
	s.mu.Lock()
	defer s.mu.Unlock()

	quotes := make([]model.Quote, 0, len(s.quotes))
	for _, q := range s.quotes {
		if author == "" || q.Author == author {
			quotes = append(quotes, q)
		}
	}
	return quotes
}

func (s *QuotesStorage) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.quotes[id]; !exists {
		return errors.New("quote not found")
	}
	delete(s.quotes, id)
	return nil
}

func (s *QuotesStorage) GetRandom() (model.Quote, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	ids := make([]int, 0, len(s.quotes))
	for k := range s.quotes {
		ids = append(ids, k)
	}
	randomIndex := ids[rand.Intn(len(ids))]
	quote, exists := s.quotes[randomIndex]

	if !exists {
		return model.Quote{}, errors.New("quote not found")
	}
	return quote, nil
}

func (s *QuotesStorage) GetByAuthor(author string) ([]model.Quote, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	quotes := make([]model.Quote, 0)
	fmt.Println(quotes)
	fmt.Println(s.quotes)
	fmt.Println(author)

	for _, v := range s.quotes {
		fmt.Println(v)
		fmt.Println(v.Author)
		if v.Author == author {
			quotes = append(quotes, v)
		}
	}
	fmt.Println(quotes)
	return quotes, nil
}
