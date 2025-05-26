package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"slices"
	"strconv"

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

func getQuotes(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	if author != "" {
		getQuotesByAuthor(w, r)
		return
	}

	getAllQuotes(w)
}

func getAllQuotes(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(quotes)
}

func getQuotesByAuthor(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")

	var filtered []Quote

	for _, q := range quotes {
		if q.Author == author {
			filtered = append(filtered, q)
		}
	}

	json.NewEncoder(w).Encode(make([]Quote, 0))
}

func deleteQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
	}

	for i, q := range quotes {
		if id == q.ID {
			quotes = slices.Delete(quotes, i, i+1)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Quote not found", http.StatusNotFound)
}

func main() {
	// TODO: add tests
	r := mux.NewRouter()

	quotes = append(quotes, Quote{
		ID:     nextID,
		Quote:  "Aboba abobcius abobenko",
		Author: "Aboba",
	})
	nextID++

	r.HandleFunc("/quotes", getQuotes).Methods("GET")
	r.HandleFunc("/quotes", createQuote).Methods("POST")
	r.HandleFunc("/quotes/random", getRandomQuote).Methods("GET")
	r.HandleFunc("/quotes/{id}", deleteQuote).Methods("DELETE")

	r.Use(jsonMiddleware)

	http.ListenAndServe(":8080", r)
}
