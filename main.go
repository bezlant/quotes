package main

import (
	"encoding/json"
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

func getAllQuotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}

func createQuote(w http.ResponseWriter, r *http.Request) {
	var q Quote
	// TODO: error handling
	json.NewDecoder(r.Body).Decode(&q)

	q.ID = nextID
	nextID++
	quotes = append(quotes, q)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
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

	http.ListenAndServe(":8080", r)
}
