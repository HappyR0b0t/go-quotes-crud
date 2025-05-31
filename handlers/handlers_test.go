package quote_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	quote "example.com/go-scout-ai-crud/handlers"
	"example.com/go-scout-ai-crud/model"
	"example.com/go-scout-ai-crud/storage"

	"github.com/gorilla/mux"
)

func setupRouter(handler *quote.QuotesHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/quotes", handler.CreateQuote).Methods("POST")
	r.HandleFunc("/quotes", handler.ListQuotes).Methods("GET")
	r.HandleFunc("/quotes/{id:[0-9]+}", handler.DeleteQuote).Methods("DELETE")
	r.HandleFunc("/quotes/random", handler.GetRandomQuote).Methods("GET")
	return r
}

func TestCreateQuote(t *testing.T) {
	store := storage.NewQuotesStorage()
	handler := quote.NewQuotesHandler(store)
	r := setupRouter(handler)

	quote := model.Quote{Author: "Confucius", Text: "Life is simple."}
	data, _ := json.Marshal(quote)
	req := httptest.NewRequest("POST", "/quotes", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.Code)
	}
}

func TestListQuotes(t *testing.T) {
	store := storage.NewQuotesStorage()
	store.Create(model.Quote{Author: "Confucius", Text: "Simplicity."})
	handler := quote.NewQuotesHandler(store)
	r := setupRouter(handler)

	req := httptest.NewRequest("GET", "/quotes?author=Confucius", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.Code)
	}
}

func TestDeleteQuote(t *testing.T) {
	store := storage.NewQuotesStorage()
	created := store.Create(model.Quote{Author: "Confucius", Text: "Simplicity."})
	handler := quote.NewQuotesHandler(store)
	r := setupRouter(handler)

	url := "/quotes/" + strconv.Itoa(created.ID)
	req := httptest.NewRequest("DELETE", url, nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusNoContent {
		t.Errorf("expected status 204, got %d", resp.Code)
	}
}

func TestGetRandomQuote(t *testing.T) {
	store := storage.NewQuotesStorage()
	store.Create(model.Quote{Author: "Confucius", Text: "Wisdom."})
	handler := quote.NewQuotesHandler(store)
	r := setupRouter(handler)

	req := httptest.NewRequest("GET", "/quotes/random", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.Code)
	}
}
