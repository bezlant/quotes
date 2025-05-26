package quotes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/quotes", h.GetQuotes).Methods("GET")
	r.HandleFunc("/quotes", h.CreateQuote).Methods("POST")
	r.HandleFunc("/quotes/random", h.GetRandomQuote).Methods("GET")
	r.HandleFunc("/quotes/{id}", h.DeleteQuote).Methods("DELETE")
}

func (h *Handler) CreateQuote(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Author string `json:"author"`
		Quote  string `json:"quote"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Author == "" || req.Quote == "" {
		http.Error(w, "Author and Quote required", http.StatusBadRequest)
		return
	}

	q := h.service.AddQuote(req.Author, req.Quote)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(q)
}

func (h *Handler) GetQuotes(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")

	var quotes []Quote

	if author != "" {
		quotes = h.service.GetQuotesByAuthor(author)
	} else {
		quotes = h.service.GetAllQuotes()
	}

	json.NewEncoder(w).Encode(quotes)
}

func (h *Handler) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	q := h.service.GetRandomQuote()

	if q == nil {
		http.Error(w, "No quotes found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(q)
}

func (h *Handler) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if h.service.DeleteQuote(id) {
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, "Quote not found", http.StatusNotFound)
	}
}
