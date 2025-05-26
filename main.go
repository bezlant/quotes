package main

import (
	"encoding/json"
	"net/http"
)

type Quote struct {
	ID    int    `json:"id"`
	Quote string `json:"quote"`
}

var quotes []Quote
var nextID = 1

func main() {
	quotes = append(quotes, Quote{
		ID:    nextID,
		Quote: "Aboba abobcius abobenko",
	})
	nextID++

	http.HandleFunc("/quotes", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(quotes)
	})

	http.ListenAndServe(":8080", nil)
}
