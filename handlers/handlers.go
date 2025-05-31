package quote

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "example.com/go-scout-ai-crud/model"
	storage "example.com/go-scout-ai-crud/storage"
	"github.com/gorilla/mux"
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
	author := r.URL.Query().Get("author")
	if len(author) != 0 {
		quotes, err := h.Store.GetByAuthor(author)
	} else {
		quotes := h.Store.List()
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}

func (h *QuotesHandler) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.Store.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *QuotesHandler) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	quote, err := h.Store.GetRandom()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quote)
}

// func (h *QuotesHandler) GetByAuthor(w http.ResponseWriter, r *http.Request) {
// 	author := mux.Vars(r)["author"]

// 	quote, err := h.Store.GetByAuthor(author)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusNotFound)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(quote)
// }
