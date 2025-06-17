package main

import (
	"log"
	"net/http"

	handlers "example.com/go-scout-ai-crud/handlers"
	storage "example.com/go-scout-ai-crud/storage"
	"github.com/gorilla/mux"
)

func main() {
	// store := storage.NewQuotesStorage()
	store, err := storage.NewPostgresStorage("postgres://user:password@localhost:5432/dbname")
	if err != nil {
		log.Fatal("Could not connect to database: ", err)
	}
	quotesHandler := handlers.NewQuotesHandler(store)

	r := mux.NewRouter()

	r.HandleFunc("/quotes", quotesHandler.CreateQuote).Methods("POST")
	r.HandleFunc("/quotes", quotesHandler.ListQuotes).Methods("GET")
	r.HandleFunc("/quotes/random", quotesHandler.GetRandomQuote).Methods("GET")
	r.HandleFunc("/quotes/{id:[0-9]+}", quotesHandler.DeleteQuote).Methods("DELETE")

	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
