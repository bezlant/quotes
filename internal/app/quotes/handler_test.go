package quotes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func setupRouter() *mux.Router {
	s := NewService()
	h := NewHandler(s)
	r := mux.NewRouter()
	h.RegisterRoutes(r)
	return r
}

func TestCreateQuote(t *testing.T) {
	r := setupRouter()
	body := `{"author":"Test","quote":"This is a test quote"}`

	req := httptest.NewRequest(http.MethodPost, "/quotes", strings.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected 201 Created, got %d", w.Code)
	}
	var q Quote
	err := json.NewDecoder(w.Body).Decode(&q)
	if err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}
	if q.Author != "Test" || q.Quote != "This is a test quote" {
		t.Errorf("Unexpected response: %+v", q)
	}
}

func TestGetQuotes(t *testing.T) {
	r := setupRouter()
	reqBody := `{"author":"Alice","quote":"Hello"}`
	r.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("POST", "/quotes", strings.NewReader(reqBody)))

	req := httptest.NewRequest("GET", "/quotes", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", w.Code)
	}
	var quotes []Quote
	err := json.NewDecoder(w.Body).Decode(&quotes)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}
	if len(quotes) != 1 {
		t.Fatalf("Expected 1 quote, got %d", len(quotes))
	}
}

func TestDeleteQuote(t *testing.T) {
	r := setupRouter()
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(map[string]string{
		"author": "X",
		"quote":  "Y",
	})

	// Create quote
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest("POST", "/quotes", &buf))
	var created Quote
	json.NewDecoder(w.Body).Decode(&created)

	// Delete it
	delReq := httptest.NewRequest("DELETE", "/quotes/"+strconv.Itoa(created.ID), nil)
	delResp := httptest.NewRecorder()
	r.ServeHTTP(delResp, delReq)

	if delResp.Code != http.StatusNoContent {
		t.Fatalf("Expected 204 No Content, got %d", delResp.Code)
	}
}
