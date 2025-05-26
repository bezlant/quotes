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

type QuoteService struct {
	quotes []Quote
	nextID int
}

func NewQuoteService() *QuoteService {
	return &QuoteService{
		quotes: make([]Quote, 0),
		nextID: 1,
	}
}

type QuoteHandler struct {
	service *QuoteService
}

func NewQuoteHandler(service *QuoteService) *QuoteHandler {
	return &QuoteHandler{service: service}
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func (qs *QuoteService) AddQuote(author, quote string) Quote {
	newQuote := Quote{
		ID:     qs.nextID,
		Author: author,
		Quote:  quote,
	}

	qs.quotes = append(qs.quotes, newQuote)
	qs.nextID++

	return newQuote
}

func (qh *QuoteHandler) CreateQuote(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Author string `json:"author"`
		Quote  string `json:"quote"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Author == "" || req.Quote == "" {
		http.Error(w, "Author an Quote fields required", http.StatusBadRequest)
		return
	}

	quote := qh.service.AddQuote(req.Author, req.Quote)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(quote)
}

func (qs *QuoteService) GetRandomQuote() *Quote {
	if len(qs.quotes) == 0 {
		return nil
	}

	q := qs.quotes[rand.Intn(len(qs.quotes))]

	return &q
}

func (gh *QuoteHandler) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	quote := gh.service.GetRandomQuote()

	if quote == nil {
		http.Error(w, "No quotes available", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(quote)
}

func (qs *QuoteService) GetAllQuotes() []Quote {
	quotes := make([]Quote, len(qs.quotes))
	copy(quotes, qs.quotes)

	return quotes
}

func (qs *QuoteService) GetQuotesByAuthor(author string) []Quote {
	var filtered []Quote

	for _, q := range qs.quotes {
		if q.Author == author {
			filtered = append(filtered, q)
		}
	}

	return filtered
}

func (qh *QuoteHandler) GetQuotes(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")

	var quotes []Quote
	if author != "" {
		quotes = qh.service.GetQuotesByAuthor(author)
	} else {
		quotes = qh.service.GetAllQuotes()
	}

	json.NewEncoder(w).Encode(quotes)
}

func (qs *QuoteService) DeleteQuote(id int) bool {
	for i, q := range qs.quotes {
		if id == q.ID {
			qs.quotes = slices.Delete(qs.quotes, i, i+1)
			qs.nextID--
			return true
		}
	}
	return false
}

func (qh *QuoteHandler) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if qh.service.DeleteQuote(id) {
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, "Quote not found", http.StatusNotFound)
	}
}

func main() {
	// TODO: add tests

	service := NewQuoteService()
	handler := NewQuoteHandler(service)

	r := mux.NewRouter()

	r.HandleFunc("/quotes", handler.GetQuotes).Methods("GET")
	r.HandleFunc("/quotes", handler.CreateQuote).Methods("POST")
	r.HandleFunc("/quotes/random", handler.GetRandomQuote).Methods("GET")
	r.HandleFunc("/quotes/{id}", handler.DeleteQuote).Methods("DELETE")

	r.Use(jsonMiddleware)

	http.ListenAndServe(":8080", r)
}
