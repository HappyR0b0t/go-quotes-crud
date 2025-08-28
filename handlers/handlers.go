package quote

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "example.com/go-scout-ai-crud/model"
	"github.com/gorilla/mux"
)

type QuotesHandler struct {
	Storage model.QuoteStorage
}

func NewQuotesHandler(store model.QuoteStorage) *QuotesHandler {
	return &QuotesHandler{Storage: store}
}

// Create a quote #DONE
func (h *QuotesHandler) CreateQuote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var quote model.Quote
	if err := json.NewDecoder(r.Body).Decode(&quote); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	id, err := h.Storage.AddQuote(ctx, quote)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	quote.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quote)
}

// Get quotes by author
func (h *QuotesHandler) ListQuotes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	author := r.URL.Query().Get("author")
	quotes, err := h.Storage.GetQuotesByAuthor(ctx, author)

	if err != nil {
		http.Error(w, "failed to fetch quotes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}

// Delete a quote #DONE
func (h *QuotesHandler) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.Storage.DeleteQuote(ctx, id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Get a random quote #DONE
func (h *QuotesHandler) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	quote, err := h.Storage.GetRandomQuote(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quote)
}

func (h *QuotesHandler) Index(w http.ResponseWriter, r *http.Request) {
	response := "Henlo!"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
