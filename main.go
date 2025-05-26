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

func main() {
	r := mux.NewRouter()

	quotes = append(quotes, Quote{
		ID:    nextID,
		Quote: "Aboba abobcius abobenko",
	})
	nextID++

	http.ListenAndServe(":8080", nil)
	r.HandleFunc("/quotes", getAllQuotes).Methods("GET")
}
