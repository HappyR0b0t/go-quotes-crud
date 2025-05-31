package main

import (
	"log"
	"net/http"

	handlers "example.com/go-scout-ai-crud/handlers"
	storage "example.com/go-scout-ai-crud/storage"
	"github.com/gorilla/mux"
)

func main() {
	store := storage.NewQuotesStorage()
	quotesHandler := handlers.NewQuotesHandler(store)

	r := mux.NewRouter()

	r.HandleFunc("/quotes", quotesHandler.CreateQuote).Methods("POST")
	r.HandleFunc("/quotes", quotesHandler.ListQuotes).Methods("GET")

	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
