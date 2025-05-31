package quote

import (
	"encoding/json"
	"net/http"

	model "example.com/go-scout-ai-crud/model"
	storage "example.com/go-scout-ai-crud/storage"
)

type QuotesHandler struct {
	Store *storage.QuotesStorage
}

func NewQuotesHandler(store *storage.QuotesStorage) *QuotesHandler {
	return &QuotesHandler{Store: store}
}

func (h *QuotesHandler) CreateQuote(w http.ResponseWriter, r *http.Request) {
	var quote model.Quote
	if err := json.NewDecoder(r.Body).Decode(&quote); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	created := h.Store.Create(quote)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(created)
}

func (h *QuotesHandler) ListQuotes(w http.ResponseWriter, r *http.Request) {
	quotes := h.Store.List()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}
