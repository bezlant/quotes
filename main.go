package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

type Quote struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

var quotes []Quote
var nextID = 1

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func getAllQuotes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(quotes)
}

func createQuote(w http.ResponseWriter, r *http.Request) {
	var q Quote
	// TODO: error handling
	json.NewDecoder(r.Body).Decode(&q)

	q.ID = nextID
	nextID++
	quotes = append(quotes, q)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(q)
}

func getRandomQuote(w http.ResponseWriter, r *http.Request) {
	// TODO: error handling (no quotes)

	q := quotes[rand.Intn(len(quotes))]

	json.NewEncoder(w).Encode(q)
}

func main() {
	r := mux.NewRouter()

	quotes = append(quotes, Quote{
		ID:    nextID,
		Quote: "Aboba abobcius abobenko",
	})
	nextID++

	r.HandleFunc("/quotes", getAllQuotes).Methods("GET")
	r.HandleFunc("/quotes", createQuote).Methods("POST")
	r.HandleFunc("/quotes/random", getRandomQuote).Methods("GET")

	r.Use(jsonMiddleware)

	http.ListenAndServe(":8080", r)
}
