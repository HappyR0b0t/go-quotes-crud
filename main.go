package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	handlers "example.com/go-scout-ai-crud/handlers"
	storage "example.com/go-scout-ai-crud/storage"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		user, password, host, port, dbname,
	)

	store, err := storage.NewPostgresStorage(connString)
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
