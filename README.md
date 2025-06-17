# Go Quotes REST API

A simple in-memory REST API for managing quotes.  
Built with Go, using Gorilla Mux for routing.

## Features

- Add a new quote (`POST /quotes`)
- List all quotes (`GET /quotes`)
- List all quotes filtered by author (`GET /quotes?author=Name`)
- Delete a quote by ID (`DELETE /quotes/{id}`)
- Get a random quote (`GET /quotes/random`)
- In-memory data store with auto-incrementing IDs
- Basic unit tests using Go's standard library

## Tech Stack

- Language: Go
- Routing: [Gorilla Mux](https://github.com/gorilla/mux)
- Testing: `net/http/httptest`, `testing`, `encoding/json`
- No external database – data is stored in memory

## Running Tests
go test ./...

## Running the Server
go run .

## Example Requests with curl:

### Create a new quote:
curl -X POST http://localhost:8080/quotes \
  -H "Content-Type: application/json" \
  -d '{"author":"Confucius", "text":"Life is simple, but we insist on making it complicated."}'

### Get all quotes:
curl http://localhost:8080/quotes

### Filter quotes by author:
curl http://localhost:8080/quotes?author=Confucius

### Delete a quote:
curl -X DELETE http://localhost:8080/quotes/1

### Get a random quote:
curl http://localhost:8080/quotes/random
