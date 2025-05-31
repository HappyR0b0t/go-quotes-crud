package storage

import (
	"strings"
	"testing"

	model "example.com/go-scout-ai-crud/model"
)

func TestCreateQuote(t *testing.T) {
	store := NewQuotesStorage()
	quote := model.Quote{Author: "Confucius", Text: "Life is simple"}
	created := store.Create(quote)

	if created.ID != 1 {
		t.Errorf("expected ID 1, got %d", created.ID)
	}

	if created.Author != quote.Author || created.Text != quote.Text {
		t.Errorf("created quote does not match input")
	}
}

func TestListQuotes(t *testing.T) {
	store := NewQuotesStorage()
	store.Create(model.Quote{Author: "A", Text: "One"})
	store.Create(model.Quote{Author: "B", Text: "Two"})
	store.Create(model.Quote{Author: "A", Text: "Three"})

	all := store.List("")
	if len(all) != 3 {
		t.Errorf("expected 3 quotes, got %d", len(all))
	}

	filtered := store.List("A")
	if len(filtered) != 2 {
		t.Errorf("expected 2 quotes for author A, got %d", len(filtered))
	}
}

func TestDeleteQuote(t *testing.T) {
	store := NewQuotesStorage()
	q := store.Create(model.Quote{Author: "Someone", Text: "Something"})

	err := store.Delete(int(q.ID))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	err = store.Delete(int(q.ID))
	if err == nil {
		t.Errorf("expected error when deleting non-existent quote")
	}
}

func TestGetRandomQuote(t *testing.T) {
	store := NewQuotesStorage()
	store.Create(model.Quote{Author: "One", Text: "A"})
	store.Create(model.Quote{Author: "Two", Text: "B"})

	quote, err := store.GetRandom()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if strings.TrimSpace(quote.Text) == "" {
		t.Errorf("expected a non-empty quote")
	}
}

func TestGetByAuthor(t *testing.T) {
	store := NewQuotesStorage()
	store.Create(model.Quote{Author: "A", Text: "One"})
	store.Create(model.Quote{Author: "B", Text: "Two"})
	store.Create(model.Quote{Author: "A", Text: "Three"})

	quotes, err := store.GetByAuthor("A")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(quotes) != 2 {
		t.Errorf("expected 2 quotes for author A, got %d", len(quotes))
	}
}
